package models

import (
	"log"
	"gopkg.in/mgo.v2/bson"
)

// EnterpriseModel : Acciones sobre collection de project
type EnterpriseModel struct {
}

//GetEnterprise , Retorna una lista de empresas
func (em *EnterpriseModel) GetEnterprise( user string) []Enterprise{
	session, err := GetSession()
	defer session.Close()
	userDAO := session.DB(DataBaseName).C(UserColl)
	enterpriseDAO := session.DB(DataBaseName).C(EnterpriseColl)
	
	userResult := User{}
	enterpriseResult:=[]Enterprise{}
	
	err = userDAO.Find(bson.M{"username": user}).One(&userResult)
	if err != nil {
		log.Fatal("Error al obtener el usuario .", err)
	}

	orQuery := []bson.M{}
	for _, data := range userResult.Enterprises {
		temp:=bson.M{"_id":data}
		orQuery=append(orQuery,temp)	
	}

	err = enterpriseDAO.Find(bson.M{"$or":orQuery}).All(&enterpriseResult)
	if err != nil {
		log.Fatal("Error al obtener las empresas.", err)
	}
		
	return enterpriseResult
}
