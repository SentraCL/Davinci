package controllers

import (
	"fmt"
	"log"

	models "../models"
)



//ExportToXML retorna los usuarios de un proyecto
func (pc *ProjectController) ExportToXML(projectCode string, epicRQ EpicExportRequest) models.Epic {
	epic := models.Epic{} 
	return epic
}


//AddUser agrega un usuario al proyecto
func (pc *ProjectController) GetUserStoryByCodeVersion(projectCode string, uscode string, indexVersion int) models.UserStory {
	log.Println("Controller - GetUserStoryByCodeVersion ", projectCode)
	userStory := projectModel.GetUserStoryByCode(projectCode, uscode)[0]
	log.Println("Controller - GetUserStoryByCodeVersion ", userStory.ID)
	return userStory
}

//AddUser agrega un usuario al proyecto
func (pc *ProjectController) AddUser(projectCode string, user string, pass string, isDesign bool) bool {
	log.Println("Controller - AddUser ", projectCode)
	projectModel.AddUser(projectCode, user, pass,isDesign)
	return true
}

//DelUser elimina un usuario al proyecto
func (pc *ProjectController) DelUser(projectCode string, user string) bool {
	log.Println("Controller - DelUser ", projectCode)
	projectModel.DelUser(projectCode, user)
	return true
}

//UpdateUser modifica un usuario al proyecto
func (pc *ProjectController) UpdateUser(projectCode string, code string, user string, pass string) bool {
	log.Println("Controller - updateUser ", projectCode)
	projectModel.UpdateUser(projectCode, code, user, pass)
	return true
}

//LoginUser modifica un usuario al proyecto
func (pc *ProjectController) LoginUser(projectName string, user string, pass string) bool {
	log.Println("Controller - LoginUser ", projectName)
	return projectModel.LoginUser(projectName, user, pass)
}

//GetUsers retorna los usuarios de un proyecto
func (pc *ProjectController) GetUsers(projectCode string) []models.UserProject {
	log.Println("Controller - GetUsers ", projectCode)
	return projectModel.GetUsers(projectCode)
}

//GetUsers retorna los usuarios de un proyecto
func (pc *ProjectController) SaveUserStory(projectCode string, userStory models.UserStory) models.UserStory {
	userStory = projectModel.SaveUserStory(projectCode, userStory)
	return userStory
}

//GetAllUserStoriesByCode retorna los usuarios de un proyecto
func (pc *ProjectController) GetUserStoryByCode(projectCode string, code string) []models.UserStory {
	fmt.Println("ProjectController>> projectCode:", projectCode)
	fmt.Println("ProjectController>> codeValue:", code)
	userStories := projectModel.GetUserStoryByCode(projectCode, code)
	return userStories
}

//GetAllUserStoriesByCode retorna los usuarios de un proyecto
func (pc *ProjectController) GetAllUserStoriesbyPreconditions(projectCode string, queryPC []models.PreConditionQuery) []models.UserStory {
	userStories := projectModel.GetAllUserStoriesbyPreconditions(projectCode, queryPC)
	return userStories
}

//SaveEpic retorna los usuarios de un proyecto
func (pc *ProjectController) SaveEpic(projectCode string, epic models.Epic) models.Epic {
	projectModel.SaveEpic(projectCode, &epic)
	return epic
}

//GetAllEpics : Guarda proyectos.
func (dc *DesignController) GetAllEpics(projectCode string) []models.Epic {
	result := projectModel.GetAllEpics(projectCode)
	return result
}

//SaveDataType Guarda el tipo de dato y retorna el objecto con su ID generado
func (pc *ProjectController) SaveDataType(projectCode string, dataType models.DataType) models.DataType {
	projectModel.SaveDataType(projectCode, &dataType)
	return dataType
}

//SaveRepositoryData Guarda la informacion generada para los tipos de dato
func (pc *ProjectController) SaveRepositoryData(projectCode string, dataType models.TestData) string {
	return projectModel.SaveRepositoryData(projectCode, &dataType)
}

//GetAllDataTypes Retorna todos los tipos de inventos por proyecto
func (pc *ProjectController) GetAllDataTypes(projectCode string) []models.DataType {
	dataTypes := projectModel.GetAllDataTypes(projectCode)
	return dataTypes
}

//GetDataTypeByID Retorna un tipo de data segun ID
func (pc *ProjectController) GetDataTypeByID(projectCode string, id string) models.DataType {
	dataType := projectModel.GetDataTypeByID(projectCode, id)
	return dataType
}
