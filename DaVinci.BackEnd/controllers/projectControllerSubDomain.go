package controllers

import (
	"fmt"
	"log"

	models "../models"
)

//AddUser agrega un usuario al proyecto
func (pc *ProjectController) GetUserStoryByCodeVersion(projectCode string, uscode string, indexVersion int) models.UserStory {
	log.Println("Controller - GetUserStoryByCodeVersion ", projectCode)
	userStory := projectModel.GetUserStoryByCode(projectCode, uscode)[0]
	log.Println("Controller - GetUserStoryByCodeVersion ", userStory.ID)
	return userStory
}

//AddUser agrega un usuario al proyecto
func (pc *ProjectController) AddUser(projectCode string, user string, pass string) bool {
	log.Println("Controller - AddUser ", projectCode)
	projectModel.AddUser(projectCode, user, pass)
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
