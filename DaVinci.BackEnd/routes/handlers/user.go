package handlers

import (
	"strings"
	"encoding/json"
	"net/http"
	"log"
)

//SaveUser , Guarda un usuario 
func (h *Handler) SaveUser(responseW http.ResponseWriter, request *http.Request) {

	postData := h.getPostValues(request)
	log.Println("REQUEST postData >>", postData["name"])
	isDesign := false
	if(postData["isDesign"]=="1"){
		isDesign=true
	}
	var res = userCtrl.AddUser(postData["name"], postData["pass"],isDesign,postData["enterprise"])
	log.Println("Resultado ", res)
}

//SaveUserForm , Guarda un usuario desde el panel de administracion
func (h *Handler) SaveUserForm(responseW http.ResponseWriter, request *http.Request) {

	postData := h.getPostValues(request)
	log.Println("REQUEST postData >>", postData["name"])
	log.Println("REQUEST postData >>", postData)
	enterprises:=strings.Split(postData["Enterprises"],";")
	log.Println("enterprises >>",enterprises)
	var res = userCtrl.SaveUserForm(postData["UserName"], postData["Password"],enterprises,postData["Role"])
	log.Println("Resultado ", res)
	h.ResponseJSON(responseW, res)
}

//DeleteUser , Borra un usuario
func (h *Handler) DeleteUser(responseW http.ResponseWriter, request *http.Request) {

	postData := h.getPostValues(request)
	log.Println("REQUEST postData >>", postData["UserName"])
	var res = userCtrl.DeleteUser(postData["UserName"])
	log.Println("Resultado ", res)
	h.ResponseJSON(responseW, res)
}

//EditUser , Edita un usuario
func (h *Handler) EditUser(responseW http.ResponseWriter, request *http.Request) {

	postData := h.getPostValues(request)
	log.Println("REQUEST postData >>", postData["UserName"])
	enterprises:=strings.Split(postData["Enterprises"],";")
	log.Println("enterprises >>",enterprises)
	var res = userCtrl.EditUser(postData["UserName"], postData["Password"],enterprises,postData["Role"])
	log.Println("Resultado ", res)
	h.ResponseJSON(responseW, res)
}

//SaveUser , Guarda un usuario 
func (h *Handler) GetRoleByHash(responseW http.ResponseWriter, request *http.Request) {
	log.Println("GetRoleByHash ")
	decoder := json.NewDecoder(request.Body)
	var hash struct {
		LoginHash string
	}
	decoder.Decode(&hash)
	
	user := h.GetUserAlias(request, DavinciCookieName)
	role:=userCtrl.FindRoleUserByName(user)

	log.Println("Resultado ", user)
	log.Println("role ", role)

	h.ResponseJSON(responseW, role)
}

//SaveUser , Guarda un usuario 
func (h *Handler) GetAllDomainUsers(responseW http.ResponseWriter, request *http.Request) {
	log.Println("getAllDomainUsers ")
	userList:=userCtrl.GetAllDomainUsers()
	log.Println("Resultado ", userList)
	h.ResponseJSON(responseW, userList)
}