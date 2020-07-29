
package controllers

import (
	"log"
	models "../models"
)
type EnterpriseController struct {
}

//GetEnterprise retorna una lista de empresas para un usuario
func (uc *EnterpriseController) GetEnterprise(user string)  []models.Enterprise {
	log.Println("Controller - GetEnterprise ", user)
	return enterpriseModel.GetEnterprise( user)
}

//GetAllEnterprises , Retorna una lista de todas las empresas
func (uc *EnterpriseController) GetAllEnterprises()  []models.Enterprise {
	return enterpriseModel.GetAllEnterprises()
}

//GetProjectEnterprise retorna una lista de empresas para un proyecto
func (uc *EnterpriseController) GetProjectEnterprise(project string)  string {
	log.Println("Controller - GetProjectEnterprise ", project)
	return enterpriseModel.GetProjectEnterprise(project)
}

//Save : Upsert Enterprise!!
func (uc *EnterpriseController) Save(enterpriseRQ models.EnterpriseRequest) (bool) {
	enterprise := uc.translateRequestToBO(enterpriseRQ)
	status:=enterpriseModel.Save(&enterprise)
	if status {
		return true
	}
	return false
}


//Drop : realiza un borrado logico de la empresa
func (uc *EnterpriseController) Drop(enterpriseRQ models.EnterpriseRequest) bool {
	enterprise := uc.translateRequestToBO(enterpriseRQ)
	if enterpriseModel.Drop(&enterprise) {
		return true
	}
	return false
}

//Recovery : Recupera una empresa eliminada
func (uc *EnterpriseController) Recovery(enterpriseRQ models.EnterpriseRequest) bool {
	enterprise := uc.translateRequestToBO(enterpriseRQ)
	if enterpriseModel.Recovery(&enterprise) {
		return true
	}
	return false
}


func (uc *EnterpriseController) translateRequestToBO(enterpriseRQ models.EnterpriseRequest) models.Enterprise {
	enterprise := models.Enterprise{}
	enterprise.EnterpriseId=enterpriseRQ.EnterpriseId
	enterprise.Name = enterpriseRQ.Name
	enterprise.Direction = enterpriseRQ.Direction
	enterprise.Description = enterpriseRQ.Description
	enterprise.Rut = enterpriseRQ.Rut
	enterprise.Avatar = enterpriseRQ.Avatar64
	enterprise.Status = enterpriseRQ.Status
	return enterprise
}

func (uc *EnterpriseController) translateBOToRequest(enterprise models.Enterprise) models.EnterpriseRequest {
	enterpriseRQ := models.EnterpriseRequest{}
	enterpriseRQ.EnterpriseId=enterprise.EnterpriseId
	enterpriseRQ.Name = enterprise.Name
	enterpriseRQ.Direction = enterprise.Direction
	enterpriseRQ.Description = enterprise.Description
	enterpriseRQ.Rut = enterprise.Rut
	enterpriseRQ.Avatar64 = enterprise.Avatar
	enterpriseRQ.Status = enterprise.Status
	return enterpriseRQ
}