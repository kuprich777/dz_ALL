package main

import (
	"sync"
	"time"
)

// Transaction описывает структуру транзакции
type Transaction struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"` // USD, EUR и т.д.
	Type        string    `json:"type"`     // доход, расход, перевод
	Category    string    `json:"category"` // зарплата, еда, жилье и т.д.
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// Тут храним транзакции
var (
	transactions = make(map[string]Transaction)
	mutex        sync.RWMutex
)
