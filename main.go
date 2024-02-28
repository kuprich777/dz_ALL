package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Budget struct {
	Incomes  map[string]float64
	Expenses map[string]float64
}

func NewBudget() *Budget {
	return &Budget{
		Incomes:  make(map[string]float64),
		Expenses: make(map[string]float64),
	}
}

func (b *Budget) AddIncome(category string, amount float64) {
	b.Incomes[category] += amount
}

func (b *Budget) AddExpense(category string, amount float64) {
	b.Expenses[category] += amount
}

func (b *Budget) TotalIncome() float64 {
	var total float64
	for _, amount := range b.Incomes {
		total += amount
	}
	return total
}

func (b *Budget) TotalExpenses() float64 {
	var total float64
	for _, amount := range b.Expenses {
		total += amount
	}
	return total
}

func (b *Budget) NetIncome() float64 {
	return b.TotalIncome() - b.TotalExpenses()
}

func (b *Budget) AnalyzeBudget() {
	for category, expense := range b.Expenses {
		percentage := (expense / b.TotalIncome()) * 100
		if percentage > 20 {
			fmt.Printf("Внимание: Траты в категории \"%s\" составляют %.2f%% от общего дохода\n", category, percentage)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	budget := NewBudget()

	fmt.Println("Введите месячные доходы по категориям (формат: категория=сумма). Для завершения введите 'конец'.")
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "конец" {
			break
		}

		parts := strings.Split(input, "=")
		if len(parts) != 2 {
			fmt.Println("Неправильный формат ввода. Пожалуйста, используйте формат 'категория=сумма'.")
			continue
		}

		amount, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Println("Ошибка при конвертации суммы в число:", err)
			continue
		}

		budget.AddIncome(parts[0], amount)
	}

	fmt.Println("Введите месячные расходы по категориям (формат: категория=сумма). Для завершения введите 'конец'.")
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "конец" {
			break
		}

		parts := strings.Split(input, "=")
		if len(parts) != 2 {
			fmt.Println("Неправильный формат ввода. Пожалуйста, используйте формат 'категория=сумма'.")
			continue
		}

		amount, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Println("Ошибка при конвертации суммы в число:", err)
			continue
		}

		budget.AddExpense(parts[0], amount)
	}

	fmt.Printf("Общий доход: %.2f\n", budget.TotalIncome())
	fmt.Printf("Общие расходы: %.2f\n", budget.TotalExpenses())
	fmt.Printf("Чистый доход: %.2f\n", budget.NetIncome())

	budget.AnalyzeBudget()
}
