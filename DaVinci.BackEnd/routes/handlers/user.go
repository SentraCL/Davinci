package handlers

import (
	/*
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	models "../../models"
	util "../../util"

	"github.com/gorilla/mux"
	*/
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