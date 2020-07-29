package handlers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	models "../../models"
)

//GetEnterpriseByToken, retorna la lista de empresas en las que el usuario tiene acceso
func (h *Handler)  GetEnterpriseByToken(responseW http.ResponseWriter, request *http.Request){
	user := h.getUser(request)
	var res = enterpriseCtrl.GetEnterprise(user)
	h.ResponseJSON(responseW, res)
}

//GetAllEnterprises , Retorna una lista de todas las empresas
func (h *Handler)  GetAllEnterprises(responseW http.ResponseWriter, request *http.Request){
	var res = enterpriseCtrl.GetAllEnterprises()
	h.ResponseJSON(responseW, res)
}

func (h *Handler)  GetProjectEnterprise(responseW http.ResponseWriter, request *http.Request){	
	params := mux.Vars(request)
	project := params["project"]
	var res = enterpriseCtrl.GetProjectEnterprise(project)
	h.ResponseJSON(responseW, res)
}

//SaveEnterprise : Guarda la empresa.
func (h *Handler) SaveEnterprise(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	enterpriseRQ := models.EnterpriseRequest{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&enterpriseRQ)
	fmt.Println(enterpriseRQ)
	enterpriseCtrl.Save(enterpriseRQ)

}

//DropEnterprise : Elimina una empresa.
func (h *Handler) DropEnterprise(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	enterpriseRQ := models.EnterpriseRequest{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&enterpriseRQ)
	enterpriseCtrl.Drop(enterpriseRQ)
}

//Recovery : Recupera una empresa eliminada
func (h *Handler) RecoveryEnterprise(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	enterpriseRQ := models.EnterpriseRequest{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&enterpriseRQ)
	enterpriseCtrl.Recovery(enterpriseRQ)
}