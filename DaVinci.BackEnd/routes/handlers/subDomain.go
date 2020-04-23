package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	models "../../models"
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
