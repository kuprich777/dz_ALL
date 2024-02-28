package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите строку: ")
	reader := bufio.NewReader(os.Stdin) // Создаем reader для чтения ввода с консоли
	input, _ := reader.ReadString('\n')  // Читаем строку, введенную пользователем

	// Выводим исходную строку
	fmt.Println("Исходная строка: ", input)

	// Создаем слайс rune для хранения символов строки
	var runes []rune
	for _, r := range input {
		runes = append(runes, r) // Добавляем каждый символ (rune) в слайс
	}

	// Преобразуем слайс rune обратно в строку для корректного вывода символов
	output := string(runes)
	fmt.Println("Строка символов: ", output)

	// Выводим полученный слайс символов в виде числовых значений Unicode
	fmt.Println("Слайс символов (коды Unicode): ", runes)
}
