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
	var res = userCtrl.AddUser(postData["name"], postData["pass"],isDesign)
	log.Println("Resultado ", res)
}