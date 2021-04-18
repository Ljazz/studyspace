package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("babcbef", "b")
	want := []string{"a", "c", "ef"}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("want:%v but got:%v", want, ret)
	}
}

func TestSplit2(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := []testCase{
		{"babcbef", "b", []string{"a", "c", "ef"}},
		{"a:b:c:d", ":", []string{"a", "b", "c", "d"}},
		{"abcdbcef", "bc", []string{"a", "d", "ef"}},
		{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("want: %#v\t\tgot: %#v\n", tc.want, got)
		}
	}
}

// 子测试
func TestSplit3(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case_1": {"babcbef", "b", []string{"a", "c", "ef"}},
		"case_2": {"a:b:c:d", ":", []string{"a", "b", "c", "d"}},
		"case_3": {"abcdbcef", "bc", []string{"a", "d", "ef"}},
		"case_4": {"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want: %#v\t\tgot: %#v\n", tc.want, got)
			}
		})
	}
}

// BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
