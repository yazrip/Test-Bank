package controller

import (
	"bank-test/dto"
	"bank-test/entity"
	"bank-test/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func NewTransactionController(transactionService *service.TransactionService) TransactionController {
	return TransactionController{TransactionService: *transactionService}
}

type TransactionController struct {
	TransactionService service.TransactionService
}

func (controller *TransactionController) Route(router, auth *mux.Router) {
	auth.HandleFunc("/transactions", controller.GetAll).Methods("GET")
	auth.HandleFunc("/transaction/{id}", controller.Create).Methods("POST")
}

func (controller *TransactionController) Create(w http.ResponseWriter, r *http.Request) {
	var transaction entity.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := context.Get(r, "user").(*dto.UserDto)
	transaction.SetAccountId(user.Id)

	err := controller.TransactionService.Create(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Transaction Successfully"}`))
}

func (controller *TransactionController) GetAll(w http.ResponseWriter, r *http.Request) {
	transactions := controller.TransactionService.GetAll()

	message, err := json.Marshal(&transactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}
