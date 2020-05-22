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

//Copy : Copiar Proyecto
func (pc *ProjectController) Copy(projectCopyRQ models.CopyProjectRequest) bool {

	project := projectModel.GetProjectByCode(projectCopyRQ.Code)
	project.Name = projectCopyRQ.Name
	project.Code = ""
	//Se concatena el ".CP", para determinar que es una copia, y concer esta referencia.
	project.Alias = strings.Replace(projectCopyRQ.Name, " ", "", -1)

	if projectCopyRQ.Data == 0 {
		project.Repository = nil
	}
	if projectCopyRQ.Users == 0 {
		project.Users = nil
	}
	if projectCopyRQ.Epics == 0 {
		project.Epics = models.Epics{}
	}
	if projectCopyRQ.UserStories == 0 {
		project.UserStories = models.UserStories{}
	}

	if projectModel.Create(&project) {
		return true
	}
	return false
}


/*

type ProjectInOutRequest struct{
	Project Project `json:"project"`
	Inventions []InventionVO `json:"inventions"`
}

*/

//ImportProject , lee un TO ProjectInOutRequest y lo registra en BD
func (pc *ProjectController) ImportProject(proIO *models.ProjectInOutRequest) {
	//TODO: Aca leer el TO y hacer el insert
	//Insertarlos.. OJO!!! Origin
	//TODO : Si los inventos ya existen
	//proIO.Inventions
	//Creai el project 
	//proIO.Project
	
	/*
	inventionController := InventionController{}
	inventVO := proIO.Inventions // Leer listado de inventos
	inventBO := inventionController.TranslateRequestToBO(inventVO.) //cambiar de VO a BO
	log.Println(proIO)
	log.Println(inventBO)
	*/

	inventionController := InventionController{}
	log.Println(proIO.Inventions)
	projectModel.Save(&proIO.Project)
	inventionController.SaveAll(proIO.Inventions) 
		
		for _, user := range proIO.Project.Users{
			pc.AddUser(proIO.Project.Code, user.UserName, user.Password)
		}

}

//Export : Copiar Proyecto
func (pc *ProjectController) Export(projectExportRQ models.ExportProjectRequest) models.ProjectInOutRequest {

	project := pc.GetProjectByCode(projectExportRQ.Code)

	if projectExportRQ.Avatar == 0 {
		project.Avatar64 = ""
		log.Print("avatar = nil")
	}
	if projectExportRQ.Data == 0 {
		project.Repository = nil
		log.Print("repository = nil")
	}
	if projectExportRQ.Users == 0 {
		project.Users = nil
		log.Print("users = nil")
	}
	if projectExportRQ.Epics == 0 {
		project.Epics = models.Epics{}
		log.Print("epics = nil")
	}
	if projectExportRQ.UserStories == 0 {
		project.UserStories = models.UserStories{}
		log.Print("stories = nil")
	}

	proyectExport := models.ProjectInOutRequest{}
	proyectExport.Project = project
	proyectExport.Inventions = pc.FiltersInventionsFromProject(project)

	return proyectExport
}

//FiltersInventionsFromProject , Traer inventos a partir del proyecto.
func (pc *ProjectController) FiltersInventionsFromProject(project models.Project) []models.InventionVO {
	inventions := []models.InventionVO{}
	inventionController := InventionController{}
	// 1) Del "project", capturo el listado de inventos desde (project.Repository) * foreach, en la iteracion (item.CODE)
	for _, data := range project.Repository {
		bo := inventionController.GetInventionByCode(data.InventionCode)
		vo := inventionController.TranslateBOToRequest(bo)
		inventions = append(inventions, vo)
		log.Print(util.StringifyJSON(vo))
	}

	// 2) Agregai inventos al "inventions", append

	return inventions
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

//GetProjectByCode : Retorna el proyecto segun su codigo
func (pc *ProjectController) GetProjectByCode(code string) models.Project {
	project := projectModel.GetProjectByCode(code)
	//log.Println(util.StringifyJSON(project))
	return project
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
		inventionVO := invCtrol.TranslateBOToRequest(invention)
		inventionVO.Data = repository.WareHouse
		inventions = append(inventions, inventionVO)
	}
	return inventions
}
