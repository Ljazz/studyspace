package main

import "fmt"

/*
我们创建了一个author struct，它包含字段名、lastName和bio。我们还添加了一个方法fullName()，将作者作为接收者类型，这将返回作者的全名。
*/
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

/*
post struct有字段标题、内容。它还有一个嵌入式匿名字段作者。这个字段表示post struct是由author组成的。现在post struct可以访问作者结构的所有字段和方法。我们还在post struct中添加了details()方法，它打印出作者的标题、内容、全名和bio。
*/
type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.author.fullName())
	fmt.Println("Bio: ", p.bio)
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()
}

