// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/tools/internal/imports"
	"golang.org/x/tools/internal/lsp/debug/tag"
	"golang.org/x/tools/internal/lsp/mod"
	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
	"golang.org/x/tools/internal/telemetry/event"
	errors "golang.org/x/xerrors"
)

func (s *Server) codeAction(ctx context.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	snapshot, fh, ok, err := s.beginFileRequest(params.TextDocument.URI, source.UnknownKind)
	if !ok {
		return nil, err
	}
	uri := fh.Identity().URI

	// Determine the supported actions for this file kind.
	supportedCodeActions, ok := snapshot.View().Options().SupportedCodeActions[fh.Identity().Kind]
	if !ok {
		return nil, fmt.Errorf("no supported code actions for %v file kind", fh.Identity().Kind)
	}

	// The Only field of the context specifies which code actions the client wants.
	// If Only is empty, assume that the client wants all of the possible code actions.
	var wanted map[protocol.CodeActionKind]bool
	if len(params.Context.Only) == 0 {
		wanted = supportedCodeActions
	} else {
		wanted = make(map[protocol.CodeActionKind]bool)
		for _, only := range params.Context.Only {
			wanted[only] = supportedCodeActions[only]
		}
	}
	if len(wanted) == 0 {
		return nil, errors.Errorf("no supported code action to execute for %s, wanted %v", uri, params.Context.Only)
	}

	var codeActions []protocol.CodeAction
	switch fh.Identity().Kind {
	case source.Mod:
		if diagnostics := params.Context.Diagnostics; len(diagnostics) > 0 {
			codeActions = append(codeActions, mod.SuggestedFixes(ctx, snapshot, fh, diagnostics)...)
		}
		if !wanted[protocol.SourceOrganizeImports] {
			codeActions = append(codeActions, protocol.CodeAction{
				Title: "Tidy",
				Kind:  protocol.SourceOrganizeImports,
				Command: &protocol.Command{
					Title:     "Tidy",
					Command:   "tidy",
					Arguments: []interface{}{fh.Identity().URI},
				},
			})
		}
	case source.Go:
		diagnostics := params.Context.Diagnostics

		var importEdits []protocol.TextEdit
		var importEditsPerFix []*source.ImportFix
		var analysisQuickFixes []protocol.CodeAction
		var highConfidenceEdits []protocol.TextDocumentEdit

		// Retrieve any necessary import edits or fixes.
		if wanted[protocol.QuickFix] && len(diagnostics) > 0 || wanted[protocol.SourceOrganizeImports] {
			importEdits, importEditsPerFix, err = source.AllImportsFixes(ctx, snapshot, fh)
			if err != nil {
				return nil, err
			}
		}
		// Retrieve any necessary analysis fixes or edits.
		if (wanted[protocol.QuickFix] || wanted[protocol.SourceFixAll]) && len(diagnostics) > 0 {
			analysisQuickFixes, highConfidenceEdits, err = analysisFixes(ctx, snapshot, fh, diagnostics)
			if err != nil {
				event.Error(ctx, "analysis fixes failed", err, tag.URI.Of(uri))
			}
		}

		if wanted[protocol.QuickFix] && len(diagnostics) > 0 {
			// First, add the quick fixes reported by go/analysis.
			codeActions = append(codeActions, analysisQuickFixes...)

			// If we also have diagnostics for missing imports, we can associate them with quick fixes.
			if findImportErrors(diagnostics) {
				// Separate this into a set of codeActions per diagnostic, where
				// each action is the addition, removal, or renaming of one import.
				for _, importFix := range importEditsPerFix {
					// Get the diagnostics this fix would affect.
					if fixDiagnostics := importDiagnostics(importFix.Fix, diagnostics); len(fixDiagnostics) > 0 {
						codeActions = append(codeActions, protocol.CodeAction{
							Title: importFixTitle(importFix.Fix),
							Kind:  protocol.QuickFix,
							Edit: protocol.WorkspaceEdit{
								DocumentChanges: documentChanges(fh, importFix.Edits),
							},
							Diagnostics: fixDiagnostics,
						})
					}
				}
			}

			// Get any actions that might be attributed to missing modules in the go.mod file.
			actions, err := mod.SuggestedGoFixes(ctx, snapshot, diagnostics)
			if err != nil {
				event.Error(ctx, "quick fixes failed", err, tag.URI.Of(uri))
			}
			if len(actions) > 0 {
				codeActions = append(codeActions, actions...)
			}
		}
		if wanted[protocol.SourceOrganizeImports] && len(importEdits) > 0 {
			codeActions = append(codeActions, protocol.CodeAction{
				Title: "Organize Imports",
				Kind:  protocol.SourceOrganizeImports,
				Edit: protocol.WorkspaceEdit{
					DocumentChanges: documentChanges(fh, importEdits),
				},
			})
		}
		if wanted[protocol.SourceFixAll] && len(highConfidenceEdits) > 0 {
			codeActions = append(codeActions, protocol.CodeAction{
				Title: "Simplifications",
				Kind:  protocol.SourceFixAll,
				Edit: protocol.WorkspaceEdit{
					DocumentChanges: highConfidenceEdits,
				},
			})
		}
	default:
		// Unsupported file kind for a code action.
		return nil, nil
	}
	return codeActions, nil
}

func (s *Server) getSupportedCodeActions() []protocol.CodeActionKind {
	allCodeActionKinds := make(map[protocol.CodeActionKind]struct{})
	for _, kinds := range s.session.Options().SupportedCodeActions {
		for kind := range kinds {
			allCodeActionKinds[kind] = struct{}{}
		}
	}
	var result []protocol.CodeActionKind
	for kind := range allCodeActionKinds {
		result = append(result, kind)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

// findImports determines if a given diagnostic represents an error that could
// be fixed by organizing imports.
// TODO(rstambler): We need a better way to check this than string matching.
func findImportErrors(diagnostics []protocol.Diagnostic) bool {
	for _, diagnostic := range diagnostics {
		// "undeclared name: X" may be an unresolved import.
		if strings.HasPrefix(diagnostic.Message, "undeclared name: ") {
			return true
		}
		// "could not import: X" may be an invalid import.
		if strings.HasPrefix(diagnostic.Message, "could not import: ") {
			return true
		}
		// "X imported but not used" is an unused import.
		// "X imported but not used as Y" is an unused import.
		if strings.Contains(diagnostic.Message, " imported but not used") {
			return true
		}
	}
	return false
}

func importFixTitle(fix *imports.ImportFix) string {
	var str string
	switch fix.FixType {
	case imports.AddImport:
		str = fmt.Sprintf("Add import: %s %q", fix.StmtInfo.Name, fix.StmtInfo.ImportPath)
	case imports.DeleteImport:
		str = fmt.Sprintf("Delete import: %s %q", fix.StmtInfo.Name, fix.StmtInfo.ImportPath)
	case imports.SetImportName:
		str = fmt.Sprintf("Rename import: %s %q", fix.StmtInfo.Name, fix.StmtInfo.ImportPath)
	}
	return str
}

func importDiagnostics(fix *imports.ImportFix, diagnostics []protocol.Diagnostic) (results []protocol.Diagnostic) {
	for _, diagnostic := range diagnostics {
		switch {
		// "undeclared name: X" may be an unresolved import.
		case strings.HasPrefix(diagnostic.Message, "undeclared name: "):
			ident := strings.TrimPrefix(diagnostic.Message, "undeclared name: ")
			if ident == fix.IdentName {
				results = append(results, diagnostic)
			}
		// "could not import: X" may be an invalid import.
		case strings.HasPrefix(diagnostic.Message, "could not import: "):
			ident := strings.TrimPrefix(diagnostic.Message, "could not import: ")
			if ident == fix.IdentName {
				results = append(results, diagnostic)
			}
		// "X imported but not used" is an unused import.
		// "X imported but not used as Y" is an unused import.
		case strings.Contains(diagnostic.Message, " imported but not used"):
			idx := strings.Index(diagnostic.Message, " imported but not used")
			importPath := diagnostic.Message[:idx]
			if importPath == fmt.Sprintf("%q", fix.StmtInfo.ImportPath) {
				results = append(results, diagnostic)
			}
		}
	}
	return results
}

func analysisFixes(ctx context.Context, snapshot source.Snapshot, fh source.FileHandle, diagnostics []protocol.Diagnostic) ([]protocol.CodeAction, []protocol.TextDocumentEdit, error) {
	if len(diagnostics) == 0 {
		return nil, nil, nil
	}

	var codeActions []protocol.CodeAction
	var sourceFixAllEdits []protocol.TextDocumentEdit

	phs, err := snapshot.PackageHandles(ctx, fh)
	if err != nil {
		return nil, nil, err
	}
	// We get the package that source.Diagnostics would've used. This is hack.
	// TODO(golang/go#32443): The correct solution will be to cache diagnostics per-file per-snapshot.
	ph, err := source.WidestPackageHandle(phs)
	if err != nil {
		return nil, nil, err
	}
	for _, diag := range diagnostics {
		// This code assumes that the analyzer name is the Source of the diagnostic.
		// If this ever changes, this will need to be addressed.
		srcErr, analyzer, err := snapshot.FindAnalysisError(ctx, ph.ID(), diag.Source, diag.Message, diag.Range)
		if err != nil {
			continue
		}
		for _, fix := range srcErr.SuggestedFixes {
			action := protocol.CodeAction{
				Title:       fix.Title,
				Kind:        protocol.QuickFix,
				Diagnostics: []protocol.Diagnostic{diag},
				Edit:        protocol.WorkspaceEdit{},
			}
			for uri, edits := range fix.Edits {
				fh, err := snapshot.GetFile(uri)
				if err != nil {
					event.Error(ctx, "no file", err, tag.URI.Of(uri))
					continue
				}
				docChanges := documentChanges(fh, edits)
				if analyzer.HighConfidence {
					sourceFixAllEdits = append(sourceFixAllEdits, docChanges...)
				}
				action.Edit.DocumentChanges = append(action.Edit.DocumentChanges, docChanges...)
			}
			codeActions = append(codeActions, action)
		}
	}
	return codeActions, sourceFixAllEdits, nil
}

func documentChanges(fh source.FileHandle, edits []protocol.TextEdit) []protocol.TextDocumentEdit {
	return []protocol.TextDocumentEdit{
		{
			TextDocument: protocol.VersionedTextDocumentIdentifier{
				Version: fh.Identity().Version,
				TextDocumentIdentifier: protocol.TextDocumentIdentifier{
					URI: protocol.URIFromSpanURI(fh.Identity().URI),
				},
			},
			Edits: edits,
		},
	}
}