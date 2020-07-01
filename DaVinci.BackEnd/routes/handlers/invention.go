package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"	
	models "../../models"
)

//GetAllInventions : Retorna todos los inventos!!
func (h *Handler) GetAllInventions(responseW http.ResponseWriter, request *http.Request) {
	inventionVOs := inventionCtrl.GetAll()
	//fmt.Println(util.StringifyJSON(inventionVOs))
	h.ResponseJSON(responseW, inventionVOs)
}

//GetAllInventions : Retorna todos los inventos!!
func (h *Handler) GetAllInventionsFiltered(responseW http.ResponseWriter, request *http.Request) {
	postData := h.getPostValues(request)
	fmt.Println(postData)

	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(request)
	fmt.Println(request.Body)
	fmt.Println(request.GetBody)

	inventionVOs := inventionCtrl.GetAll()
	//fmt.Println(util.StringifyJSON(inventionVOs))
	h.ResponseJSON(responseW, inventionVOs)
}

//SaveInventions : Guarda todos los inventos
func (h *Handler) SaveInventions(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	inventionRQs := []models.InventionVO{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&inventionRQs)
	inventionRQs = inventionCtrl.SaveAll(inventionRQs)
	h.ResponseJSON(responseW, inventionRQs)
}

//SaveInvention : Guarda un invento
func (h *Handler) SaveInvention(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	inventionRQ := models.InventionVO{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&inventionRQ)
	//log.Println("GUARDAR : " + util.StringifyJSON(inventionRQ))
	inventionRQ = inventionCtrl.Save(inventionRQ)
	h.ResponseJSON(responseW, inventionRQ)
}

//GetDefaultArtifact : Entrega Los Artefactos por defecto
func (h *Handler) GetDefaultArtifact(responseW http.ResponseWriter, request *http.Request) {

	artifactVOs := inventionCtrl.GetDefaultArtifact()
	h.ResponseJSON(responseW, artifactVOs)
}

//DropInventions : Guarda todos los inventos
func (h *Handler) DropInventions(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	inventionRQ := models.InventionVO{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&inventionRQ)
	response := inventionCtrl.Drop(inventionRQ)
	h.ResponseJSON(responseW, response)
}
