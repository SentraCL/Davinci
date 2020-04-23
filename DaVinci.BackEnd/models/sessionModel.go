package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SessionModel : Acciones sobre collection de User
type SessionModel struct {
}

//IsUserValid , valida crenciales
func (sm *SessionModel) IsUserValid(login *Login) bool {
	//log.Println("SessionModel >> func (sm *SessionModel) IsUserValid(login *Login) bool {")
	//Obtener Conexion.
	session, err := GetSession()
	//Se ejecuta una vez salido de la funcion
	defer session.Close()
	userDAO := session.DB(DataBaseName).C(UserColl)
	userResult := User{}

	err = userDAO.Find(bson.M{"username": login.User, "password": login.Pass}).One(&userResult)
	//log.Println("response :", util.StringifyJSON(userResult))
	isValid := false

	if err == nil {
		isValid = (login.User == userResult.UserName && login.Pass == userResult.Password)
	} else {
		isValid = false
	}
	////log.Println("Es un Usuario Valido : ", isValid)
	return isValid
}

//IsHashOnline : chequea si el hash enviado desde front-end se encuentra online
func (sm *SessionModel) IsHashOnline(user string, DavinciCode string) bool {
	//log.Println("SessionModel >> func (sm *SessionModel) IsHashOnline(user string, DavinciCode string) bool {")
	//Obtener Conexion.
	session, err := GetSession()
	//Se ejecuta una vez salido de la funcion
	defer session.Close()
	userDAO := session.DB(DataBaseName).C(UserColl)
	userResult := User{}
	//log.Println("username :", user)
	err = userDAO.Find(bson.M{"username": user}).One(&userResult)
	//log.Println("response :", util.StringifyJSON(userResult))
	//log.Println("HistoryLogin :", len(userResult.HistoryLogin))
	if err != nil {
		//log.Println("ERROR!!", err)
	}
	if err == nil {
		for _, session := range userResult.HistoryLogin {
			//log.Println("session.DavinciCode == DavinciCode :", session.DavinciCode == DavinciCode)
			//log.Println("session.Online :", session.Online)
			if session.DavinciCode == DavinciCode && session.Online {
				return true
			}
		}
	}
	return false
}

//SaveSession , Guarda las sesiones asociadas al usuario
func (sm *SessionModel) SaveSession(login *Login) {
	//log.Println("SessionModel >> func (sm *SessionModel) SaveSession(login *Login) {")

	//Obtener Conexion.
	session, err := GetSession()
	//Se ejecuta una vez salido de la funcion
	defer session.Close()
	userDAO := session.DB(DataBaseName).C(UserColl)
	userResult := User{}
	err = userDAO.Find(bson.M{"username": login.User}).One(&userResult)
	//log.Println("response :", util.StringifyJSON(userResult))
	if err == nil {
		sesion := Session{}
		sesion.DavinciCode = login.DavinciCode
		sesion.Date = time.Now()
		//Capturar fecha actual y registrar remote ip
		sesion.Online = login.Online
		sesion.RemoteAddr = login.RemoteAddr
		//ultima semana
		newHistoryLogin := []Session{}
		for _, sesionHist := range userResult.HistoryLogin {
			sesionHist.Online = false
			if sesionHist.DavinciCode != login.DavinciCode {
				newHistoryLogin = append(newHistoryLogin, sesionHist)
			}
		}
		newHistoryLogin = append(newHistoryLogin, sesion)
		userDAO.Update(
			//Where
			bson.M{"username": login.User, "password": login.Pass},
			//Set
			bson.M{"$set": bson.M{"historylogin": newHistoryLogin, "lasttime": sesion.Date}},
		)
	}
}
