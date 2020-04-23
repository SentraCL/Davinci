package controllers

import (
	"encoding/json"
	"log"

	//	"log"
	"strings"

	models "../models"
	util "../util"
)

// ProjectController ,Orquestador de Modelo del proyecto!!
type ProjectController struct {
}

type MapArtifactType map[string][]string

func (pc *ProjectController) translateRequestToBO(projectRQ models.ProjectRequest) models.Project {
	project := models.Project{}
	project.Author = projectRQ.Admin
	project.Name = projectRQ.Name
	project.Alias = strings.Replace(projectRQ.Name, " ", "", -1)
	project.Resume = projectRQ.Resume
	project.Company = projectRQ.Company
	project.Avatar64 = projectRQ.Avatar64
	contactAdmin := models.Contact{}
	contactAdmin.FullName = projectRQ.Name
	contactAdmin.Email = projectRQ.Email
	project.Administrator = contactAdmin
	return project
}

func (pc *ProjectController) translateBOToRequest(project models.Project) models.ProjectRequest {
	projectRQ := models.ProjectRequest{}
	projectRQ.Code = project.Code
	projectRQ.Admin = project.Author
	projectRQ.Alias = strings.Replace(project.Name, " ", "", -1)
	projectRQ.Name = project.Name
	projectRQ.Resume = project.Resume
	projectRQ.Company = project.Company
	projectRQ.Avatar64 = project.Avatar64
	projectRQ.Users = project.Users

	for _, repo := range project.Repository {
		repoRQ := models.RepositoryRef{}
		repoRQ.InventionCode = repo.InventionCode
		repoRQ.Invention = repo.Invention
		repoRQ.KeyRef = repo.KeyRef
		repoRQ.LastTime = repo.LastTime
		projectRQ.Repository = append(projectRQ.Repository, repoRQ)
	}

	projectRQ.Email = project.Administrator.Email
	return projectRQ
}

//GetAll : Upsert Project!!
func (pc *ProjectController) GetAll() []models.ProjectRequest {
	//models.Project
	projectsRqs := []models.ProjectRequest{}
	projects := projectModel.GetAll()
	for _, project := range projects {
		projectRQ := pc.translateBOToRequest(project)
		projectsRqs = append(projectsRqs, projectRQ)
	}
	return projectsRqs
}

//Save : Upsert Project!!
func (pc *ProjectController) Save(projectRQ models.ProjectRequest) bool {
	project := pc.translateRequestToBO(projectRQ)
	if projectModel.Save(&project) {
		return true
	}
	return false
}

//Drop : delete all references and Project!!
func (pc *ProjectController) Drop(projectRQ models.ProjectRequest) bool {
	project := pc.translateRequestToBO(projectRQ)
	if projectModel.Drop(&project) {
		return true
	}
	return false
}

//GetAvatarByAlias : Upsert Project!!
func (pc *ProjectController) GetAvatarByAlias(alias string) string {
	dcode := util.DavinciCode{}
	projectCode := dcode.Encript(alias)
	project := projectModel.GetProjectByCode(projectCode)
	//log.Println(util.StringifyJSON(project))
	return project.Avatar64
}

//AddInvention :Agrega datos de inventos por Codigo Proyecto
//projectCode,inventionCode,util.StringifyJSON(data)
func (pc *ProjectController) AddInvention(pid models.ProjectInventionData) models.ProjectInventionData {
	projectModel.AddInvention(pid)
	return pid
}

//DelInvention :elimina invento por Codigo Proyecto
//projectCode,inventionCode,util.StringifyJSON(data)
func (pc *ProjectController) DelInvention(pid models.ProjectInventionData) models.ProjectInventionData {
	projectModel.DelInvention(pid)
	return pid
}

//GetInventions , retorna inventos por proyecto
func (pc *ProjectController) GetInventions(projectCode string, inventionCode string) map[string]MapArtifactType {
	pid := models.ProjectInventionData{}
	pid.ProjectCode = projectCode
	pid.InventionCode = inventionCode
	inventions := projectModel.GetInventions(pid)

	//log.Println(">>> " + util.StringifyJSON(inventions))

	var response = map[string]MapArtifactType{}
	for _, data := range inventions {
		result := make(MapArtifactType)
		json.Unmarshal([]byte(data.ContentJSON), &result)
		result[data.KeyValue+"_date"] = []string{data.Update.String()}
		response[data.KeyValue] = result
	}
	return response
}

//GetAllInventions , retorna inventos por proyecto
func (pc *ProjectController) GetAllInventions(projectCode string) []models.InventionVO {
	log.Println("Controller - Trayendo inventos desde ", projectCode)
	repos := projectModel.GetAllInventions(projectCode)
	invCtrol := InventionController{}
	inventions := []models.InventionVO{}
	for _, repository := range repos {
		invention := inventionModel.GetInventionByCode(repository.InventionCode)
		inventionVO := invCtrol.translateBOToRequest(invention)
		inventionVO.Data = repository.WareHouse
		inventions = append(inventions, inventionVO)
	}
	return inventions
}

