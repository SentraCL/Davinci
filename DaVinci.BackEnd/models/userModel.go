package models

import (
	util "../util"
	"gopkg.in/mgo.v2/bson"
)

// UserModel : Acciones sobre collection de user
type UserModel struct {
}

//AddUser , Agrega un usuario a la collection user
func (em *UserModel) AddUser( user string, pass string, isDesign bool) bool{
	session, err := GetSession()
	defer session.Close()
	userCollector := session.DB(DataBaseName).C(UserColl)
	userResult := User{}
	err = userCollector.Find(bson.M{"username": user}).One(&userResult)
	salida:=false
	if err != nil && err.Error() == "not found" {	
		dcode := util.DavinciCode{}
		var user = bson.M{"_id": dcode.Encript(user), "user": user, "password": pass, "isDesign":isDesign}
		userCollector.Insert(bson.M(user))
		salida=true
	}
	return salida
}
