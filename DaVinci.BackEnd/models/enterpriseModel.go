package models

import (
	"log"
	"gopkg.in/mgo.v2/bson"
)

// EnterpriseModel : Acciones sobre collection de project
type EnterpriseModel struct {
}

//GetEnterprise , Retorna una lista de empresas del usuario
func (em *EnterpriseModel) GetEnterprise( user string) []Enterprise{
	session, err := GetSession()
	defer session.Close()
	userDAO := session.DB(DataBaseName).C(UserColl)
	enterpriseDAO := session.DB(DataBaseName).C(EnterpriseColl)
	
	userResult := User{}
	enterpriseResult:=[]Enterprise{}
	log.Println("user: "+user)
	err = userDAO.Find(bson.M{"username": user}).One(&userResult)
	if err != nil {
		log.Println("Error al obtener el usuario .", err)
		return nil
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
//GetAllEnterprises , Retorna una lista de todas las empresas
func (em *EnterpriseModel) GetAllEnterprises() []Enterprise{
	session, err := GetSession()
	defer session.Close()
	enterpriseDAO := session.DB(DataBaseName).C(EnterpriseColl)
	enterpriseResult:=[]Enterprise{}
	err = enterpriseDAO.Find(nil).All(&enterpriseResult)
	if err != nil {
		log.Fatal("Error al obtener las empresas.", err)
	}
	return enterpriseResult
}

//GetEnterprise , Retorna una lista de las empresas del proyecto
func (em *EnterpriseModel) GetProjectEnterprise(project string) string{
	session, err := GetSession()
	defer session.Close()
	projectDAO := session.DB(DataBaseName).C(ProjectColl)

	projectResult := Project{}
	log.Println("project: "+project)
	err = projectDAO.Find(bson.M{"_id": project}).One(&projectResult)
	if err != nil {
		log.Println("Error al obtener la empresa .", err)
	}
	return projectResult.Enterprise
}

//Save , Guarda cualquier cambio de la empresa
func (em *EnterpriseModel) Save(enterprise *Enterprise) (bool) {
	session, err := GetSession()
	defer session.Close()
	enterpriseDAO := session.DB(DataBaseName).C(EnterpriseColl)
	if enterprise.EnterpriseId=="" {
		enterprise.Status=1
		enterprise.EnterpriseId=bson.NewObjectId().Hex()
	}
	log.Println("enterpriseId",enterprise.EnterpriseId)
	if err == nil {
			upsertdata := bson.M{
				"$set": bson.M{
					"description":   enterprise.Description,
					"name":          enterprise.Name,
					"direction":     enterprise.Direction,
					"rut":        	 enterprise.Rut,
					"avatar":      	 enterprise.Avatar,
					"status":		 enterprise.Status,
				}}

			enterpriseDAO.UpsertId(
				//Where
				enterprise.EnterpriseId,
				//Set
				upsertdata,
			)
			return true
	}
	return false
}

//Drop , Realiza un borrado logico de la empresa
func (em *EnterpriseModel) Drop(enterprise *Enterprise) bool {
	session, err := GetSession()
	defer session.Close()
	enterpriseDAO := session.DB(DataBaseName).C(EnterpriseColl)
	if err == nil {
		enterpriseDAO.Update(bson.M{"_id": enterprise.EnterpriseId}, bson.M{"$set": bson.M{"status": 0}})
		return true
	}
	return false
}

//Recovery , Recupera una empresa eliminada
func (em *EnterpriseModel) Recovery(enterprise *Enterprise) bool {
	session, err := GetSession()
	defer session.Close()
	enterpriseDAO := session.DB(DataBaseName).C(EnterpriseColl)
	if err == nil {
		enterpriseDAO.Update(bson.M{"_id": enterprise.EnterpriseId}, bson.M{"$set": bson.M{"status": 1}})
		return true
	}
	return false
}