package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	models "../../models"
	controllers "../../controllers"
	util "../../util"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"net/http"
)

//UserStoryRefRQ , requerst para obtener los datos de referen cias de los usuarios.
type UserStoryRefRQ struct {
	Code         string
	IndexVersion int
}

//GetUserStoryByRef , busca historias de usuarios por codigo y version
func (h *Handler) GetUserStoryByRef(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)

		usRefRQ := UserStoryRefRQ{}
		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&usRefRQ)
		log.Println("BUSCO : " + util.StringifyJSON(usRefRQ))
		userStory := projectCtrl.GetUserStoryByCodeVersion(projectCode, usRefRQ.Code, usRefRQ.IndexVersion)
		h.ResponseJSON(responseW, userStory)
	}
}

func (h *Handler) isOnline(request *http.Request) bool {
	params := mux.Vars(request)
	project := params["project"]

	dcode := util.DavinciCode{}
	cookieProjectName := dcode.Encript(project)

	subkey := []byte(cookieProjectName)
	subStore := sessions.NewCookieStore(subkey)
	projectCookie := project + cookieProjectName
	session, _ := subStore.Get(request, projectCookie)

	auth, isOnline := session.Values["authenticated"].(bool)

	fmt.Println("isOnline : auth 		=>", auth)
	fmt.Println("isOnline : isOnline	=>", isOnline)

	return auth && isOnline
}

//IsOnline : Verifica si hay una sesion por proyecto
func (h *Handler) IsOnline(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		fmt.Fprintln(responseW, "OK")
		return
	}
	// Print secret message
	fmt.Fprintln(responseW, "NOK")
}

//SubDomain : Pagina de Proyecto
func (h *Handler) SubDomain(responseW http.ResponseWriter, request *http.Request) {
	index := h.PathPublicHTML + string(os.PathSeparator) + "project.html"
	var projectHTML, _ = util.LoadFile(index)
	fmt.Fprintf(responseW, projectHTML)
}

//DoLogin , SubLogin
func (h *Handler) DoLogin(responseW http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	login := models.Login{}
	decoder.Decode(&login)
	login.Online = false
	params := mux.Vars(request)
	project := params["project"]
	log.Println(project)
	//Cookie que identifica el login por subdominio/proyecto
	dcode := util.DavinciCode{}
	cookieProjectName := dcode.Encript(project)
	subCookieHash := h.UpSetCookie(login.User, responseW, cookieProjectName)
	login.MakeDavinciCODE(subCookieHash)

	log.Println(request.Body)
	if projectCtrl.LoginUser(project, login.User, login.Pass) {
		dcode := util.DavinciCode{}
		cookieProjectName := dcode.Encript(project)
		subkey := []byte(cookieProjectName)
		subStore := sessions.NewCookieStore(subkey)
		projectCookie := project + cookieProjectName
		session, _ := subStore.Get(request, projectCookie)
		login.Online = true
		session.Values["authenticated"] = login.Online
		session.Save(request, responseW)
	}

	h.ResponseJSON(responseW, login)
}

//DoLogOut , deslogeo sub dominio
func (h *Handler) DoLogOut(responseW http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	project := params["project"]

	login := models.Login{}
	cookieHash := h.UpSetCookie(login.User, responseW, project)
	login.MakeDavinciCODE(cookieHash)
	login.User = h.GetUserAlias(request, project)
	//Registra deslogeo
	dcode := util.DavinciCode{}
	cookieProjectName := dcode.Encript(project)
	subkey := []byte(cookieProjectName)
	h.ClearCookie(responseW, cookieProjectName)
	subStore := sessions.NewCookieStore(subkey)
	projectCookie := project + cookieProjectName
	session, _ := subStore.Get(request, projectCookie)
	session.Values["authenticated"] = false
	session.Save(request, responseW)
}

//GetEpicTypesProject , EpicTypes by Project ??
func (h *Handler) GetEpicTypesProject(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)

		epicTypes := dcCtrl.GetAllEpicType(projectCode)
		h.ResponseJSON(responseW, epicTypes)
	}
}

//GetEpicsProject , EpicTypes by Project ??
func (h *Handler) GetEpicsProject(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)

		epics := dcCtrl.GetAllEpics(projectCode)
		h.ResponseJSON(responseW, epics)
	}
}

//GetCode , EpicTypes by Project ??
func (h *Handler) GetCode(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)
		h.ResponseJSON(responseW, projectCode)
	}
}

//GetInventionDef , EpicTypes by Project ??
func (h *Handler) GetInventionDef(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		inventionCode := params["invention"]
		invention := inventionCtrl.GetInventionByCode(inventionCode)
		h.ResponseJSON(responseW, invention)
	}
}

//GetTypeRef , EpicTypes by Project ??
func (h *Handler) GetTypeRef(responseW http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	typeRef := params["ref"]
	artifact := inventionCtrl.GetTypeRefByCode(typeRef)
	h.ResponseJSON(responseW, artifact)
}

//AddStep , Permite agregar un paso
func (h *Handler) AddStep(responseW http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	project := params["project"]
	log.Println("Project ", project)
	postData := h.getPostValues(request)
	log.Println("REQUEST paso >>", postData["paso"])
	log.Println("REQUEST resultado >>", postData["resultado"])
	/* decoder := json.NewDecoder(request.Body)
	login := models.Login{}
	decoder.Decode(&login)
	login.Online = false
	params := mux.Vars(request)
	project := params["project"]
	log.Println(project)
	//Cookie que identifica el login por subdominio/proyecto
	dcode := util.DavinciCode{}
	cookieProjectName := dcode.Encript(project)
	subCookieHash := h.UpSetCookie(login.User, responseW, cookieProjectName)
	login.MakeDavinciCODE(subCookieHash)

	log.Println(request.Body)
	if projectCtrl.LoginUser(project, login.User, login.Pass) {
		dcode := util.DavinciCode{}
		cookieProjectName := dcode.Encript(project)
		subkey := []byte(cookieProjectName)
		subStore := sessions.NewCookieStore(subkey)
		projectCookie := project + cookieProjectName
		session, _ := subStore.Get(request, projectCookie)
		login.Online = true
		session.Values["authenticated"] = login.Online
		session.Save(request, responseW)
	}

	h.ResponseJSON(responseW, login) */
}

//GetEpicTypesProject , EpicTypes by Project ??
func (h *Handler) GetAllUserStoriesTypeProject(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)

		userStoriesTypes := dcCtrl.GetAllUserStories(projectCode)
		h.ResponseJSON(responseW, userStoriesTypes)
	}
}

//SaveUserStoriesProject , EpicTypes by Project ??
func (h *Handler) SaveUserStoriesProject(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)

		userStory := models.UserStory{}
		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&userStory)
		log.Println("GUARDAR : " + util.StringifyJSON(userStory))
		userStory = projectCtrl.SaveUserStory(projectCode, userStory)
		h.ResponseJSON(responseW, userStory)

	}
}

//GetUserStoryByCode , EpicTypes by Project ??
func (h *Handler) GetUserStoryByCode(responseW http.ResponseWriter, request *http.Request) {
	fmt.Println("GetUserStoryByCode!!")
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		codeValue := params["code"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)
		fmt.Println("ProjectController>> projectCode:", projectCode)
		fmt.Println("ProjectController>> codeValue:", codeValue)
		userStoriesTypes := projectCtrl.GetUserStoryByCode(projectCode, codeValue)
		h.ResponseJSON(responseW, userStoriesTypes)
	}
}

//GetAllUserStoriesbyCode , EpicTypes by Project ??
func (h *Handler) GetAllUserStoriesbyPreconditions(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)

		queryPC := []models.PreConditionQuery{}
		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&queryPC)

		userStoriesTypes := projectCtrl.GetAllUserStoriesbyPreconditions(projectCode, queryPC)
		h.ResponseJSON(responseW, userStoriesTypes)
	}
}


//ExportEpicProject , EpicTypes by Project ??
func (h *Handler) ExportEpicProject(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		typeFile := params["type"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)

		epicRQ := controllers.EpicExportRequest{}
		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&epicRQ)

		log.Println("GUARDAR : ", util.StringifyJSON(epicRQ))

		if typeFile == "XML"{
			projectCtrl.ExportToXML(projectCode, epicRQ)
		}

		//Como retornar un XML

	}
}


//SaveEpicProject , EpicTypes by Project ??
func (h *Handler) SaveEpicProject(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)

		epic := models.Epic{}
		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&epic)
		log.Println("GUARDAR : ", util.StringifyJSON(epic))
		epic = projectCtrl.SaveEpic(projectCode, epic)
		h.ResponseJSON(responseW, epic)

	}
}

//SaveDataType , Save a Type Of Data Project
func (h *Handler) SaveDataType(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)
		dataType := models.DataType{}
		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&dataType)
		dataType = projectCtrl.SaveDataType(projectCode, dataType)
		h.ResponseJSON(responseW, dataType)
	}
}

//GetAllDataTypes , retorna en formato JSON Tipos de datos segun codigo de proyecto
func (h *Handler) GetAllDataTypes(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)
		dataTypes := projectCtrl.GetAllDataTypes(projectCode)
		h.ResponseJSON(responseW, dataTypes)
	}
}

//GetDataTypeByID , Retorna JSON con tipo de dato segun su ID
func (h *Handler) GetDataTypeByID(responseW http.ResponseWriter, request *http.Request) {
	if h.isOnline(request) {
		params := mux.Vars(request)
		projectName := params["project"]
		id := params["id"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)
		dataType := projectCtrl.GetDataTypeByID(projectCode, id)
		h.ResponseJSON(responseW, dataType)
	}
}

func (h *Handler) SaveRepositoryData(responseW http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	projectName := params["project"]
	dataType := models.TestData{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&dataType)
	salida := projectCtrl.SaveRepositoryData(projectName, dataType)
	h.ResponseJSON(responseW,salida)	
}

//GetIconDataTypeByID , Imagen en base64
func (h *Handler) GetIconDataTypeByID(w http.ResponseWriter, r *http.Request) {
	if h.isOnline(r) {
		params := mux.Vars(r)
		projectName := params["project"]
		id := params["id"]
		dcode := util.DavinciCode{}
		projectCode := dcode.Encript(projectName)
		dataType := projectCtrl.GetDataTypeByID(projectCode, id)
		h.ResponseError = "EOF"
		image64 := dataType.Icon
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
	
}
