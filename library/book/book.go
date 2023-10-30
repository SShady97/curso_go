package book

import "fmt"

type Book struct {
	//Atributos en minuscula son privados, en mayuscula publicos. Aplica para nombres de funciones tambien
	title  string
	author string
	pages  int
}

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

type TexBook struct {
	Book
	editorial string
	level     string
}

func NewTexBook(title, author string, pages int, editorial, level string) *TexBook {
	return &TexBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TexBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNivel: %s\n", b.title, b.author, b.pages, b.editorial, b.level)
}
