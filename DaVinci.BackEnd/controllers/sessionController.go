package controllers

import (
	"strings"

	models "../models"
	util "../util"
)

// SessionController posee los atributos y los metodos para el manejo de login
type SessionController struct {
}

// DoLogin , Valida usuario y posterior a esto registra su inicio de sesion
func (sc *SessionController) DoLogin(login *models.Login) bool {
	//log.Println("SessionController.go >> func (sc *SessionController) DoLogin(login *models.Login) bool {")
	isValid := false
	if sesionModel.IsUserValid(login) {
		isValid = true
		login.Online = isValid
		sesionModel.SaveSession(login)
	}
	return isValid
}

//DoLogout : Cierre de Sesion
func (sc *SessionController) DoLogout(login *models.Login) {
	sesionModel.SaveSession(login)
}

//CheckHashOnline : Revisa si el Hash esta online
func (sc *SessionController) CheckHashOnline(user string, loginHash string) bool {
	//log.Println("SessionController.go >> func (sc *SessionController) CheckHash(user string, loginHash string) bool {")
	isValid := false
	//Rutina de desencriptacion
	dcode := util.DavinciCode{}
	decodeHash := dcode.Decript(loginHash)
	if strings.Index(decodeHash, ":") > 0 {
		userFromHash := strings.Split(decodeHash, ":")[1]
		//log.Println("user>>", user)
		//log.Println("userFromHash>>", userFromHash)
		if user == userFromHash {
			if sesionModel.IsHashOnline(user, loginHash) {
				isValid = true
			}
		}
	}
	return isValid
}

