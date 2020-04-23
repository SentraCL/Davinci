package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	controllers "../../controllers"
	util "../../util"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var sesionCtrl = controllers.SessionController{}
var projectCtrl = controllers.ProjectController{}
var inventionCtrl = controllers.InventionController{}
var dcCtrl = controllers.DesignController{}

var (
	key   = []byte("DAVINCI2019")
	store = sessions.NewCookieStore(key)
)

//ProjectSession , Repositorio de datos para subsesiones por Proyecto
type ProjectSession struct {
	OnLine        bool
	UserName      string
	Project       string
	RemoteAddress string
}

//Handler : Definicion de POST/GET del sitio
type Handler struct {
	PathPublicHTML    string
	ResponseError     string
	SessionsInProject []ProjectSession
}

//SetPathPublicHTML : Seteo de la ruta de recursos estaticos
func (h *Handler) SetPathPublicHTML(path string) {
	h.PathPublicHTML = path
}

//ResponseJSON : Responde en formato JSON cualquier Objeto
func (h *Handler) ResponseJSON(w http.ResponseWriter, object interface{}) {

	fmt.Fprintln(w, util.StringifyJSON(object))
}

//PublicHTML :Carga archivos publicos
func (h *Handler) PublicHTML(fileName string) (string, error) {
	publicPath := h.PathPublicHTML + string(os.PathSeparator) + fileName
	bytes, err := ioutil.ReadFile(publicPath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//Status : Ver si se ve la url
func (h *Handler) Status(responseW http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(responseW, "OK")
}

//ClearCookie : Limpia las cookies
func (h *Handler) ClearCookie(response http.ResponseWriter, nameCookie string) {
	cookie := &http.Cookie{
		Name:   nameCookie,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//UpSetCookie : Setea nombre de jugador en cookie y retorna la cookie encriptada
func (h *Handler) UpSetCookie(userName string, response http.ResponseWriter, nameCookie string) string {
	cookieCode := ""

	value := map[string]string{
		"UserName": userName,
	}
	if encoded, err := cookieHandler.Encode(nameCookie, value); err == nil {
		cookie := &http.Cookie{
			Name:  nameCookie,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
		cookieCode = encoded
	}
	return cookieCode
}

//GetUserAlias : Obtiene el usuario logeado
func (h *Handler) GetUserAlias(request *http.Request, nameCookie string) (userName string) {
	if cookie, err := request.Cookie(nameCookie); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode(nameCookie, cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["UserName"]
		}
	}
	return userName
}

//GetUserAliasProject : Obtiene el usuario logeado en el proyecto
func (h *Handler) GetUserAliasProject(request *http.Request, nameCookie string) (userName string) {
	params := mux.Vars(request)
	project := params["project"]
	log.Println(project)
	//Cookie que identifica el login por subdominio/proyecto
	dcode := util.DavinciCode{}
	cookieProjectName := dcode.Encript(project)
	log.Println("cookieProjectName ", cookieProjectName)

	if cookie, err := request.Cookie(project); err == nil {
		cookieValue := make(map[string]string)
		log.Println("cookie", cookieValue)
		if err = cookieHandler.Decode(nameCookie, cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["UserName"]
		}
	}
	return userName
}

//getPostValues : Retorna simple envio de post values [string]string
func (h *Handler) getPostValues(request *http.Request) map[string]string {
	type MapArtifactType map[string]string
	postData := make(MapArtifactType)
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&postData)
	return postData
}

func (h *Handler) ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
