package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	translations := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите 'add' для добавления нового перевода, 'search' для поиска или 'exit' для выхода:")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "add":
			fmt.Println("Введите слово и его перевод через пробел:")
			scanner.Scan()
			parts := strings.Split(scanner.Text(), " ")
			if len(parts) != 2 {
				fmt.Println("Необходимо ввести ровно два слова. Попробуйте еще раз.")
				continue
			}
			word, translation := parts[0], parts[1]
			translations[word] = translation
			fmt.Printf("Перевод для '%s' добавлен.\n", word)
		case "search":
			fmt.Println("Введите слово для поиска его перевода:")
			scanner.Scan()
			word := scanner.Text()
			translation, found := translations[word]
			if !found {
				fmt.Printf("Перевод для слова '%s' не найден.\n", word)
			} else {
				fmt.Printf("Перевод для '%s' - '%s'.\n", word, translation)
			}
		case "exit":
			return
		default:
			fmt.Println("Неизвестная команда. Попробуйте еще раз.")
		}
	}
}
