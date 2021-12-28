package controller

import (
	"bank-test/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewMerchantController(MerchantService *service.MerchantService) MerchantController {
	return MerchantController{MerchantService: *MerchantService}
}

type MerchantController struct {
	MerchantService service.MerchantService
}

func (controller *MerchantController) Route(router, auth *mux.Router) {
	auth.HandleFunc("/merchants", controller.GetAll).Methods("GET")
}

func (controller *MerchantController) GetAll(w http.ResponseWriter, r *http.Request) {
	merchants := controller.MerchantService.GetAll()

	message, err := json.Marshal(&merchants)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}
