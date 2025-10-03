package main

import (
	"fmt"
)

type Book struct {
	id     int
	title  string
	author string
	taken  bool
}

type Library struct {
	books []*Book
}

func (lib *Library) find(title string) []*Book {
	var result []*Book
	for _, b := range lib.books {
		if b.title == title {
			result = append(result, b)
		}
	}
	return result
}

func main() {
	b1 := Book{id: 1, title: "Книга1", author: "Автор1"}
	b2 := Book{id: 2, title: "Книга2", author: "Автор2"}

	lib := Library{
		books: []*Book{&b1, &b2},
	}

	found := lib.find("Книга1")
	if len(found) > 0 {
		fmt.Println("Найдена:", found[0].title)
	} else {
		fmt.Println("Не найдено")
	}
}
