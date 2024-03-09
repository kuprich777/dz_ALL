import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/transactions", createTransaction).Methods("POST")
	r.HandleFunc("/transactions", getTransactions).Methods("GET")
	r.HandleFunc("/transactions/{id}", getTransactionByID).Methods("GET")
	r.HandleFunc("/transactions/{id}", updateTransaction).Methods("PUT")
	r.HandleFunc("/transactions/{id}", deleteTransaction).Methods("DELETE")
	return r
}