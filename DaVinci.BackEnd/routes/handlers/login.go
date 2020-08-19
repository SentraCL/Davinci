package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"

	models "../../models"
	util "../../util"
)

/*
//IndexPageHandler : Pagina INDEX
func (h *Handler) IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	var index, _ = h.PublicHTML("index.html")
	fmt.Fprintf(response, index)
}
*/

const DavinciCookieName = "GIOCONDAEATCOOKIES"

//Login : registra el login del usuario
func (h *Handler) Login(responseW http.ResponseWriter, request *http.Request) {
	//log.Println("login.go >> func (h *Handler) Login(responseW http.ResponseWriter, request *http.Request) {")
	decoder := json.NewDecoder(request.Body)
	login := models.Login{}
	decoder.Decode(&login)
	login.Online = false
	fmt.Println(login.User);
	if !util.IsEmpty(login.User) && !util.IsEmpty(login.Pass) {
		login.RemoteAddr = request.RemoteAddr
		cookieHash := h.UpSetCookie(login.User, responseW, DavinciCookieName)
		login.MakeDavinciCODE(cookieHash)
		isCorrect := sesionCtrl.DoLogin(&login)
		if isCorrect {
			session, _ := store.Get(request, "cookie-name")
			session.Values["authenticated"] = login.Online
			//jsonResponse := util.StringifyJSON(login)
			//log.Println("Inicio Sesion :" + jsonResponse)
			session.Save(request, responseW)
		}
	}
	h.ResponseJSON(responseW, login)
}

//IsValidHash : Valida Hash de Sesion
func (h *Handler) IsValidHash(responseW http.ResponseWriter, request *http.Request) {
	//log.Println("login.go >>  func (h *Handler) IsValidHash(responseW http.ResponseWriter, request *http.Request) {")
	decoder := json.NewDecoder(request.Body)
	var hash struct {
		LoginHash string
	}
	decoder.Decode(&hash)
	user := h.GetUserAlias(request, DavinciCookieName)
	validHash := sesionCtrl.CheckHashOnline(user, hash.LoginHash)
	session, _ := store.Get(request, "cookie-name")
	session.Values["authenticated"] = validHash
	session.Save(request, responseW)
	h.ResponseJSON(responseW, validHash)
}

//Logout : Cierra sesion el usuario
func (h *Handler) Logout(responseW http.ResponseWriter, request *http.Request) {
	login := models.Login{}
	cookieHash := h.UpSetCookie(login.User, responseW, DavinciCookieName)
	login.MakeDavinciCODE(cookieHash)
	login.User = h.GetUserAlias(request, DavinciCookieName)
	sesionCtrl.DoLogout(&login)
	h.ClearCookie(responseW, DavinciCookieName)
	session, _ := store.Get(request, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(request, responseW)

}

func (h *Handler) isDavinciOnline(request *http.Request) bool {
	//TODO : Es solo cortesia , hacer mas seguro.
	session, _ := store.Get(request, "cookie-name")
	fmt.Println("session",session)
	auth, isOnline := session.Values["authenticated"].(bool)
	fmt.Println("isOnline",isOnline)
	fmt.Println("auth",auth)
	//user := h.GetUserAlias(request, DavinciCookieName)
	return auth && isOnline
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

//Status : retorna el usuario logeado sin entregar la clave
func (h *Handler) getUser( request *http.Request) string {
	userOnline := h.GetUserAlias(request, DavinciCookieName)
	return userOnline
}

//Status : Ver si se ve la url
func (h *Handler) Status(responseW http.ResponseWriter, request *http.Request) {
	userOnline := h.GetUserAlias(request, DavinciCookieName)
	fmt.Fprintln(responseW, userOnline)
}