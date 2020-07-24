package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

//GetEnterpriseByToken, retorna la lista de empresas en las que el usuario tiene acceso
func (h *Handler)  GetEnterpriseByToken(responseW http.ResponseWriter, request *http.Request){
	user := h.getUser(request)
	var res = enterpriseCtrl.GetEnterprise(user)
	h.ResponseJSON(responseW, res)
}

func (h *Handler)  GetProjectEnterprise(responseW http.ResponseWriter, request *http.Request){	
	params := mux.Vars(request)
	project := params["project"]
	var res = enterpriseCtrl.GetProjectEnterprise(project)
	h.ResponseJSON(responseW, res)
}