package main

import (
	"bufio"
	"fmt"
	"os"
)

// Функция для добавления студента на курс
func addStudent(courseStudents map[string]int, course string) {
	// Увеличиваем счётчик студентов для заданного курса
	courseStudents[course]++
}

// Функция для вывода общего количества студентов по курсам
func printStudentCount(courseStudents map[string]int) {
	fmt.Println("Количество студентов по курсам:")
	for course, count := range courseStudents {
		fmt.Printf("Курс %s: %d студент(ов)\n", course, count)
	}
}

func main() {
	courseStudents := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите курс, чтобы добавить студента, или 'print' для вывода количества студентов по курсам, 'exit' для выхода:")
		scanner.Scan()
		input := scanner.Text()

		if input == "print" {
			printStudentCount(courseStudents)
			continue
		}

		if input == "exit" {
			break
		}

		// Добавление студента на указанный курс
		addStudent(courseStudents, input)
		fmt.Printf("Студент добавлен на курс '%s'.\n", input)
	}
}
