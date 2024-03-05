package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

type Readable interface {
	Read()
}

func (b Book) Read() {
	fmt.Printf("Читатель начал читать книгу: '%s' автора %s\n", b.Title, b.Author)
}
