
package controllers

import (
	"log"
	models "../models"
)
type EnterpriseController struct {
}

//GetEnterprise retorna una lista de empresas
func (uc *EnterpriseController) GetEnterprise(user string)  []models.Enterprise {
	log.Println("Controller - GetEnterprise ", user)
	return enterpriseModel.GetEnterprise( user)
}