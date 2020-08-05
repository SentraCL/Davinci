package models

import (
	util "../util"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// UserModel : Acciones sobre collection de user
type UserModel struct {
}

//AddUser , Agrega un usuario a la collection user
func (um *UserModel) AddUser( user string, pass string, isDesign bool,enterprise string) bool{
	session, err := GetSession()
	defer session.Close()
	userCollector := session.DB(DataBaseName).C(UserColl)
	userResult := User{}
	err = userCollector.Find(bson.M{"username": user}).One(&userResult)
	salida:=false
	log.Println("enterprise",enterprise)
	enterprises:=[1]string{enterprise}
	historylogin:= []string{}

	if err != nil && err.Error() == "not found" {	
		dcode := util.DavinciCode{}
		var user = bson.M{"_id": dcode.Encript(user), "username": user, "password": pass, "isDesign":isDesign,"enterprise":enterprises,"historylogin":historylogin}
		userCollector.Insert(bson.M(user))
		salida=true
	}
	return salida
}

//SaveUserForm , Agrega un usuario a la collection user
func (um *UserModel) SaveUserForm( user string, pass string, enterprises []string,rol string) bool{
	session, err := GetSession()
	defer session.Close()
	userCollector := session.DB(DataBaseName).C(UserColl)
	userResult := User{}
	err = userCollector.Find(bson.M{"username": user}).One(&userResult)
	salida:=false
	log.Println("enterprises",enterprises)
	historylogin:= []string{}

	if err != nil && err.Error() == "not found" {	
		dcode := util.DavinciCode{}
		var user = bson.M{"_id": dcode.Encript(user), "username": user, "password": pass, "isDesign":true,"enterprise":enterprises,"historylogin":historylogin,"role":rol}
		userCollector.Insert(bson.M(user))
		salida=true
	}
	return salida
}

//EditUser , Edita un usuario de la collection user
func (um *UserModel) EditUser( user string, pass string, enterprises []string,rol string) bool{
	session, err := GetSession()
	defer session.Close()
	userCollector := session.DB(DataBaseName).C(UserColl)
	userResult := User{}
	err = userCollector.Find(bson.M{"username": user}).One(&userResult)
	salida:=false
	log.Println("enterprises",enterprises)

	log.Println("userResult:",userResult)
	log.Println("Error:",err)
	if err == nil {	
		dcode := util.DavinciCode{}
		if(pass==""){
			userCollector.Update(bson.M{"_id": dcode.Encript(user)}, bson.M{"$set": bson.M{"username": user,"enterprise":enterprises,"role":rol}})
		}else{
			userCollector.Update(bson.M{"_id": dcode.Encript(user)}, bson.M{"$set": bson.M{"username": user, "password": pass,"enterprise":enterprises,"role":rol}})
		}
		salida=true
	}
	return salida
}

//DeleteUser , Elimina un usuario de la collection user
func (um *UserModel) DeleteUser( user string) bool{
	session, err := GetSession()
	defer session.Close()
	userCollector := session.DB(DataBaseName).C(UserColl)
	dcode := util.DavinciCode{}
	err = userCollector.Remove(bson.M{"_id": dcode.Encript(user)})
	salida:=true
	if err != nil {	
		salida=false
		log.Println("Error:",err)
	}
	return salida
}


//findRoleUserByName, retorna el rol del usuario por su nombre
func (um *UserModel) FindRoleUserByName(user string) string{
	session, err := GetSession()
	defer session.Close()
	userCollector := session.DB(DataBaseName).C(UserColl)
	userResult := User{}
	rol:="any"
	err = userCollector.Find(bson.M{"username": user}).One(&userResult)
	log.Println("err",err)
	log.Println("userResult",userResult	)
	if err == nil {	
		rol=userResult.Role
	}
	return rol
}

//GetAllDomainUsers , Retorna una lista de todos los usuarios de dominio
func (um *UserModel) GetAllDomainUsers() []UserVO{
	session, err := GetSession()
	defer session.Close()
	userDAO := session.DB(DataBaseName).C(UserColl)
	//userResult:=[]User{}
	userResult:=[]UserVO{}
	err = userDAO.Find(nil).All(&userResult)
	if err != nil {
		log.Fatal("Error al obtener los usuarios.", err)
	}
	return userResult
}
