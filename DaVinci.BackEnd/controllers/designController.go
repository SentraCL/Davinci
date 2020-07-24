package controllers

import (
	"fmt"

	models "../models"
	util "../util"
)

// DesignController ,Orquestador de Modelo del proyecto!!
type DesignController struct {
}

//SaveProject : Guarda proyectos.
func (dc *DesignController) SaveEpicType(epicTypeReq models.EpicTypeRequest,status bool) string {
	projectCode := epicTypeReq.ProjectCode
	epicType := dc.RQtoBO(epicTypeReq)
	fmt.Println("Controller BO : ", util.StringifyJSON(epicType))
	projectModel.SaveEpicType(projectCode, epicType,status)
	return epicType.Code	
}

func (dc *DesignController) VerifyBasic(projectCode string) {
	//save pred epics
	epicTypeReq:= models.EpicTypeRequest{}
	epicTypeReq.ProjectCode=projectCode
	epicTypeReq.Name="Predeterminado"
	epicTypeReq.Reference=""
	epicTypeReq.Definition="Plantilla de epico predeterminado"
	epicTypeReq.Attributes=append(epicTypeReq.Attributes,models.AttributeEpic{"Como","1i2c1c1f",""})
	epicTypeReq.Attributes=append(epicTypeReq.Attributes,models.AttributeEpic{"Quiero","20661j652063225h60",""})
	epicTypeReq.Attributes=append(epicTypeReq.Attributes,models.AttributeEpic{"Para","20661j652063225h60",""})

	dc.SaveEpicType(epicTypeReq,true)
}

func (dc *DesignController) RQtoBO(epicTypeReq models.EpicTypeRequest) models.EpicType {
	epicType := models.EpicType{}
	dcode := util.DavinciCode{}
	code := dcode.Encript(epicTypeReq.Name)
	epicType.Code = code
	epicType.Name = epicTypeReq.Name
	epicType.Reference = epicTypeReq.Reference
	epicType.Definition = epicTypeReq.Definition
	epicType.Attributes = epicTypeReq.Attributes
	epicType.Inventions = epicTypeReq.Inventions
	return epicType
}

//GetAllEpicType : Guarda proyectos.
func (dc *DesignController) GetAllEpicType(projectCode string) []models.EpicType {
	result := projectModel.GetAllEpicType(projectCode)
	return result
}

//SaveProject : Guarda proyectos.
func (dc *DesignController) SaveUserStoriesType(usType models.UserStoriesType, projectCode string) string {
	dcode := util.DavinciCode{}
	usCode := dcode.Encript(usType.Title)
	usType.Code = usCode
	fmt.Println("Controller BO : ", util.StringifyJSON(usType))
	projectModel.SaveUserStoriesType(projectCode, usType)
	return usType.Code
}

//GetAllUserStories : Guarda proyectos.
func (dc *DesignController) GetAllUserStories(projectCode string) []models.UserStoriesType {
	result := projectModel.GetAllUserStories(projectCode)
	return result
}
