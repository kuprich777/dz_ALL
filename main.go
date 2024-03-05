package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l *Library) RemoveBook(title string) {
	for i, book := range l.Books {
		if book.Title == title {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			return
		}
	}
	fmt.Println("Книга не найдена.")
}

func (l *Library) DisplayBooks() {
	if len(l.Books) == 0 {
		fmt.Println("Библиотека пуста.")
		return
	}
	for i, book := range l.Books {
		fmt.Printf("%d: '%s' автора %s, %d г.\n", i+1, book.Title, book.Author, book.Year)
	}
}

func ReadBook(r Readable) {
	r.Read()
}

func main() {
	library := &Library{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nВыберите действие: \n1. Добавить книгу\n2. Удалить книгу\n3. Показать книги\n4. Читать книгу\n5. Выйти")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Введите название книги:")
			scanner.Scan()
			title := scanner.Text()

			fmt.Println("Введите автора книги:")
			scanner.Scan()
			author := scanner.Text()

			fmt.Println("Введите год издания:")
			scanner.Scan()
			year, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Ошибка ввода года издания, попробуйте снова.")
				continue
			}

			library.AddBook(Book{Title: title, Author: author, Year: year})
			fmt.Println("Книга добавлена.")

		case "2":
			fmt.Println("Введите название книги для удаления:")
			scanner.Scan()
			title := scanner.Text()

			library.RemoveBook(title)
			fmt.Println("Книга удалена, если она была в библиотеке.")

		case "3":
			library.DisplayBooks()

		case "4":
			if len(library.Books) == 0 {
				fmt.Println("В библиотеке нет книг.")
				continue
			}
			fmt.Println("Выберите номер книги для чтения:")
			library.DisplayBooks()
			scanner.Scan()
			num, err := strconv.Atoi(scanner.Text())
			if err != nil || num < 1 || num > len(library.Books) {
				fmt.Println("Неверный ввод, попробуйте снова.")
				continue
			}

			ReadBook(library.Books[num-1])

		case "5":
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Неверный ввод, попробуйте снова.")
		}
	}
}
