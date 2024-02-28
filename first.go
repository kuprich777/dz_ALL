package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)  // Создаем scanner для чтения ввода с консоли
	scanner.Scan()                          // Читаем строку, введенную пользователем
	input := scanner.Text()                 // Получаем введенные данные в виде строки

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении ввода:", err)
		return
	}

	// Создаем слайс rune для хранения символов строки
	runes := []rune(input)

	// Выводим полученный слайс символов в удобочитаемом формате
	fmt.Print("Слайс символов: [")
	for i, r := range runes {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("'%c'", r)
	}
	fmt.Println("]")
}
