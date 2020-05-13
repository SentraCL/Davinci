package handlers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"strings"

	models "../../models"
	util "../../util"

	"github.com/gorilla/mux"
)

//CopyProject : Copia proyectos.
func (h *Handler) CopyProject(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	copyProjectRQ := models.CopyProjectRequest{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&copyProjectRQ)
	projectCtrl.Copy(copyProjectRQ)
}

//ExportProject : Exportar proyecto
func (h *Handler) ExportProject(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	exportProjectRQ := models.ExportProjectRequest{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&exportProjectRQ)
	
	
	project := projectCtrl.Export(exportProjectRQ)
	request.Header.Set("Content-Type", "text/plain")

	projectString := util.StringifyJSON(project);

	dcode := util.DavinciCode{}
	exportEncript := dcode.Decript(projectString)

	
	fmt.Fprintln(responseW, exportEncript)
	
}

//SaveProject : Guarda proyectos.
func (h *Handler) SaveProject(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	projectRQ := models.ProjectRequest{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&projectRQ)
	projectCtrl.Save(projectRQ)
}

//DropProject : Elimina un proyecto.
func (h *Handler) DropProject(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	projectRQ := models.ProjectRequest{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&projectRQ)
	projectCtrl.Drop(projectRQ)
}

//GetAllProject : Obtiene el listado de todos los proyectos.
func (h *Handler) GetAllProject(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	projects := projectCtrl.GetAll()
	request.Header.Set("Content-Type", "application/json")
	h.ResponseJSON(responseW, projects)
}



//GetAvatarProject : Obtiene el listado de todos los proyectos.
func (h *Handler) GetAvatarProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	alias := params["alias"]
	log.Println(alias)
	//data:image/png;base64,iVBO
	h.ResponseError = "EOF"
	image64 := projectCtrl.GetAvatarByAlias(alias)
	if len(image64) > 0 {
		//log.Println(image64)
		format := strings.Split(image64, ";")[0]
		format = strings.Split(format, ":")[1]

		log.Println(format)

		content64 := strings.Split(image64, ",")[1]
		w.Header().Set("Content-Type", format)
		p, err := base64.StdEncoding.DecodeString(content64)
		if err != nil {
			http.Error(w, "internal error", 500)
			return
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(p))) //len(dec)
		io.WriteString(w, string(p))
	} else {
		r.Header.Set("Content-Type", "application/json")
		h.ResponseJSON(w, h.ResponseError)
	}
}

//SaveInventionIntoProjectByFormPost , Guarda invento en repositorio de projecto
func (h *Handler) SaveInventionIntoProjectByFormPost(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	params := mux.Vars(request)
	projectCode := params["project"]
	inventionCode := params["invention"]
	_ = request.ParseForm()
	invention := inventionCtrl.GetInventionByCode(inventionCode)

	type MapArtifactType map[string][]string

	data := make(MapArtifactType)
	for a := range invention.Artifacts {
		artifact := invention.Artifacts[a]
		data[artifact.Name] = append(data[artifact.Name], request.Form.Get(artifact.Name))
	}

	pid := models.ProjectInventionData{}
	pid.ProjectCode = projectCode
	pid.KeyRef = invention.KeyValue
	pid.InventionCode = inventionCode

	originData := util.StringifyJSON(data)

	log.Println("Original DATA : ", originData)

	result := make(MapArtifactType)

	//var result = map[string]string{}
	json.Unmarshal([]byte(originData), &result)
	keyData := ""
	for key, value := range result {
		if key == invention.KeyValue {
			keyData = value[0]
			break
		}
	}
	dataInvention := models.DataInvention{}
	dataInvention.KeyValue = keyData
	dataInvention.ContentJSON = originData
	pid.Container = append(pid.Container, dataInvention)

	projectCtrl.AddInvention(pid)
	h.ResponseJSON(responseW, pid)
}

//SaveInventionIntoProject , Guarda invento en repositorio de projecto
func (h *Handler) SaveInventionIntoProject(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	params := mux.Vars(request)
	projectCode := params["project"]
	inventionCode := params["invention"]
	invention := inventionCtrl.GetInventionByCode(inventionCode)

	type MapArtifactType map[string][]string
	postData := make(MapArtifactType)

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&postData)

	log.Println("REQUEST postData >>", util.StringifyJSON(postData))

	pid := models.ProjectInventionData{}
	pid.ProjectCode = projectCode
	pid.KeyRef = invention.KeyValue
	pid.InventionCode = inventionCode
	originData := util.StringifyJSON(postData)

	result := make(MapArtifactType)

	//var result = map[string]string{}
	json.Unmarshal([]byte(originData), &result)
	keyData := ""
	for key, value := range result {
		if key == invention.KeyValue {
			keyData = value[0]
			break
		}
	}
	dataInvention := models.DataInvention{}
	dataInvention.KeyValue = keyData
	dataInvention.ContentJSON = originData
	pid.Container = append(pid.Container, dataInvention)

	projectCtrl.AddInvention(pid)
	h.ResponseJSON(responseW, pid)
}

//DropInventionIntoProject : Elimina un invento de un proyecto
func (h *Handler) DropInventionIntoProject(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	params := mux.Vars(request)
	projectCode := params["project"]
	inventionCode := params["invention"]

	log.Println(projectCode)
	log.Println(inventionCode)

	_ = request.ParseForm()
	invention := inventionCtrl.GetInventionByCode(inventionCode)

	data := map[string]string{}
	for a := range invention.Artifacts {
		artifact := invention.Artifacts[a]
		data[artifact.Name] = request.Form.Get(artifact.Name)
	}

	pid := models.ProjectInventionData{}
	pid.ProjectCode = projectCode
	pid.KeyRef = invention.KeyValue
	pid.InventionCode = inventionCode

	originData := util.StringifyJSON(data)
	var result = map[string]string{}
	json.Unmarshal([]byte(originData), &result)
	keyData := ""
	for key, value := range result {
		if key == invention.KeyValue {
			keyData = value
		}
	}
	dataInvention := models.DataInvention{}
	dataInvention.KeyValue = keyData
	dataInvention.ContentJSON = originData
	pid.Container = append(pid.Container, dataInvention)

	log.Println(util.StringifyJSON(pid))

	projectCtrl.DelInvention(pid)
	h.ResponseJSON(responseW, pid)
}

//GetInventionByProject , Obtiene el listado de todos los inventos segun contexto de proyecto.
func (h *Handler) GetInventionByProject(responseW http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	projectCode := params["project"]
	inventionCode := params["invention"]
	result := projectCtrl.GetInventions(projectCode, inventionCode)
	h.ResponseJSON(responseW, result)
}

//GetAllProjectInventions , Obtiene el listado de todos los inventos segun contexto de proyecto.
func (h *Handler) GetAllProjectInventions(responseW http.ResponseWriter, request *http.Request) {
	if !h.isDavinciOnline(request) {
		return
	}
	postData := h.getPostValues(request)
	inventions := projectCtrl.GetAllInventions(postData["projectCode"])
	h.ResponseJSON(responseW, inventions)
}

//SaveUserIntoProjectByFormPost , Guarda un usuario en un proyecto
func (h *Handler) SaveUserIntoProjectByFormPost(responseW http.ResponseWriter, request *http.Request) {

	postData := h.getPostValues(request)
	log.Println("REQUEST postData >>", postData["name"])
	params := mux.Vars(request)
	projectCode := params["project"]
	log.Print("Project Code  : " + projectCode)
	var res = projectCtrl.AddUser(projectCode, postData["name"], postData["pass"])
	log.Println("Resultado ", res)
}

//DeleteUserIntoProjectByFormDelete , Elimina un usuario del proyecto
func (h *Handler) DeleteUserIntoProjectByFormDelete(responseW http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	projectCode := params["project"]
	userName := params["userName"]
	log.Print("Project : " + projectCode)
	log.Print("UserName : " + userName)
	var res = projectCtrl.DelUser(projectCode, userName)
	log.Println("Resultado ", res)
}

//UpdateUserIntoProjectByFormPut , Actualiza un usuario de un proyecto
func (h *Handler) UpdateUserIntoProjectByFormPut(responseW http.ResponseWriter, request *http.Request) {
	postData := h.getPostValues(request)
	log.Println("REQUEST postData >>", postData["name"])
	params := mux.Vars(request)
	projectCode := params["project"]
	log.Print("Project Code  : " + projectCode)
	var res = projectCtrl.UpdateUser(projectCode, postData["code"], postData["name"], postData["pass"])
	log.Println("Resultado ", res)
}

//UpdatePasswordUserByFormPut , Actualiza un usuario de un proyecto
func (h *Handler) UpdatePasswordUserByFormPut(responseW http.ResponseWriter, request *http.Request) {

	var user = h.GetUserAliasProject(request, DavinciCookieName)
	log.Println("Usuario ", user)
	/*postData := h.getPostValues(request)
	log.Println("REQUEST postData >>", postData["name"])
	params := mux.Vars(request)
	projectCode := params["project"]
	log.Print("Project Code  : " + projectCode)
	var res = projectCtrl.UpdateUser(projectCode, postData["code"], postData["name"], postData["pass"])
	log.Println("Resultado ", res)*/
}

//GetUsersForProjectByFormGet , Actualiza un usuario de un proyecto
func (h *Handler) GetUsersForProjectByFormGet(responseW http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	projectCode := params["project"]
	log.Print("Project Code  : " + projectCode)
	var res = projectCtrl.GetUsers(projectCode)
	log.Println("Resultado ", res)
	h.ResponseJSON(responseW, res)

}
