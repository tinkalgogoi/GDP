package main

import (
	"fmt"
)

type A struct {
	year int
}

func (a A) Greet() { fmt.Println("Hello GolangUK", a.year) }

type B struct {
	A
}

func (b B) Greet() { fmt.Println("Welcome to GolangUK", b.year) }

func main() {
	var a A
	a.year = 2016
	var b B

	//b := B{A{year:2222}}

	//b.year = 2017

	a.Greet() // Hello GolangUK 2016
	b.Greet() // Welcome to GolangUK 2016
}

// So embedding is a powerful tool which allows Go’s types to be open for extension.
