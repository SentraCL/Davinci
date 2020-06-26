package handlers

import (
	"net/http"
)

//GetEnterpriseByToken, retorna la lista de empresas en las que el usuario tiene acceso
func (h *Handler)  GetEnterpriseByToken(responseW http.ResponseWriter, request *http.Request){
	user := h.getUser(request)
	var res = enterpriseCtrl.GetEnterprise(user)
	h.ResponseJSON(responseW, res)
}