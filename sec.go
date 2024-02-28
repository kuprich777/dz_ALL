package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Student структура для хранения информации о студенте
type Student struct {
	Name   string
	Age    int
	Grades []float64
}

// addStudent добавляет студента в слайс и возвращает обновленный слайс
func addStudent(students []Student, student Student) []Student {
	return append(students, student)
}

// printStudents выводит список всех студентов и информацию о них
func printStudents(students []Student) {
	fmt.Println("Список студентов:")
	for _, s := range students {
		fmt.Printf("Имя: %s, Возраст: %d, Оценки: %v\n", s.Name, s.Age, s.Grades)
	}
}

func main() {
	var students []Student
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите информацию о студенте (имя, возраст, оценки через запятую) или 'end' для завершения:")
		scanner.Scan()
		input := scanner.Text()

		// Проверка на команду завершения ввода
		if input == "end" {
			break
		}

		// Разбиваем ввод на части
		parts := strings.Split(input, ",")
		if len(parts) < 3 {
			fmt.Println("Некорректный ввод, попробуйте еще раз.")
			continue
		}

		name := strings.TrimSpace(parts[0])
		age, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			fmt.Println("Ошибка при вводе возраста:", err)
			continue
		}

		// Преобразование строк с оценками в числа
		gradesStrings := strings.Split(strings.TrimSpace(parts[2]), " ")
		var grades []float64
		for _, gradeStr := range gradesStrings {
			grade, err := strconv.ParseFloat(gradeStr, 64)
			if err != nil {
				fmt.Println("Ошибка при вводе оценок:", err)
				continue
			}
			grades = append(grades, grade)
		}

		// Создание студента и добавление в слайс
		student := Student{
			Name:   name,
			Age:    age,
			Grades: grades,
		}
		students = addStudent(students, student)
	}

	// Выводим информацию о всех студентах
	printStudents(students)
}
