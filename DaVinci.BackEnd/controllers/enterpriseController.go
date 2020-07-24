
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

//GetProjectEnterprise retorna una lista de empresas para un proyecto
func (uc *EnterpriseController) GetProjectEnterprise(project string)  string {
	log.Println("Controller - GetProjectEnterprise ", project)
	return enterpriseModel.GetProjectEnterprise(project)
}