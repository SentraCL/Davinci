package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "../../models"
	util "../../util"
	"github.com/gorilla/mux"
)

//SaveEpicType : Guarda proyectos.
func (h *Handler) SaveEpicType(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	epicType := models.EpicTypeRequest{}
	fmt.Println(util.StringifyJSON(epicType))
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&epicType)
	code := dcCtrl.SaveEpicType(epicType)
	h.ResponseJSON(responseW, code)
}

func (h *Handler) GetAllEpicType(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	params := mux.Vars(request)
	projectCode := params["project"]
	epicTypes := dcCtrl.GetAllEpicType(projectCode)
	h.ResponseJSON(responseW, epicTypes)
}

//SaveProject : Guarda proyectos.
func (h *Handler) SaveUserStoriesType(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}

	params := mux.Vars(request)
	projectCode := params["project"]

	us := models.UserStoriesType{}
	fmt.Println(util.StringifyJSON(us))
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&us)
	code := dcCtrl.SaveUserStoriesType(us, projectCode)
	h.ResponseJSON(responseW, code)
}

func (h *Handler) GetAllUserStoriesType(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	params := mux.Vars(request)
	projectCode := params["project"]
	userStoriesTypes := dcCtrl.GetAllUserStories(projectCode)
	h.ResponseJSON(responseW, userStoriesTypes)
}
