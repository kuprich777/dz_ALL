package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// createTransaction обрабатывает POST запросы
func createTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	transaction.ID = strconv.Itoa(rand.Intn(1000000))
	transactions = append(transactions, transaction)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func getTransactionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range transactions {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func updateTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range transactions {
		if item.ID == params["id"] {
			var updatedTransaction Transaction
			if err := json.NewDecoder(r.Body).Decode(&updatedTransaction); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updatedTransaction.ID = item.ID
			transactions[index] = updatedTransaction // Обновляем транзакцию в массиве
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTransaction)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range transactions {
		if item.ID == params["id"] {
			transactions = append(transactions[:index], transactions[index+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions) // Отправляем обновленный список транзакций
}
