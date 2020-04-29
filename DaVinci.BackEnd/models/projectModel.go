package models

import (
	"fmt"
	"log"
	"strconv"
	"time"

	util "../util"
	"gopkg.in/mgo.v2/bson"
)

// ProjectModel : Acciones sobre collection de project
type ProjectModel struct {
}

//GetAll , Retorna todos los proyectos
func (pm *ProjectModel) GetAll() []Project {
	session, err := GetSession()
	defer session.Close()
	projectDAO := session.DB(DataBaseName).C(ProjectColl)
	var results []Project
	err = projectDAO.Find(nil).All(&results)
	if err != nil {
		log.Fatal("Error al obtener los proyectos .", err)
	}
	return results
}

//Drop , Guarda cualquier cambio en el proyecto
func (pm *ProjectModel) Drop(project *Project) bool {
	session, err := GetSession()
	defer session.Close()
	projectDAO := session.DB(DataBaseName).C(ProjectColl)
	//log.Println("response :", util.StringifyJSON(projectResult))
	if err == nil {
		project.Date = time.Now()
		dcode := util.DavinciCode{}
		project.Code = dcode.Encript(project.Alias)
		log.Println("Project Code >>", project.Code)
		projectDAO.Remove(bson.M{"_id": project.Code})
		return true
	}
	return false
}

//Create , Crea un proyecto
func (pm *ProjectModel) Create(project *Project) bool {
	session, err := GetSession()
	defer session.Close()
	projectDAO := session.DB(DataBaseName).C(ProjectColl)
	//log.Println("response :", util.StringifyJSON(projectResult))
	if err == nil {
		project.Date = time.Now()
		dcode := util.DavinciCode{}
		project.Code = dcode.Encript(project.Alias)
		log.Println("Project Code >>", project.Code)
		//upsertdata := bson.M{"$set": project}

		upsertdata := bson.M{
			"$set": project}

		projectDAO.UpsertId(
			//Where
			project.Code,
			//Set
			upsertdata,
		)

		return true
	}
	return false
}

//Save , Guarda cualquier cambio en el proyecto
func (pm *ProjectModel) Save(project *Project) bool {
	session, err := GetSession()
	defer session.Close()
	projectDAO := session.DB(DataBaseName).C(ProjectColl)
	//log.Println("response :", util.StringifyJSON(projectResult))
	if err == nil {
		project.Date = time.Now()
		dcode := util.DavinciCode{}
		project.Code = dcode.Encript(project.Alias)
		log.Println("Project Code >>", project.Code)
		//upsertdata := bson.M{"$set": project}

		upsertdata := bson.M{
			"$set": bson.M{
				"author":        project.Author,
				"name":          project.Name,
				"alias":         project.Alias,
				"resume":        project.Resume,
				"administrator": project.Administrator,
				"avatar64":      project.Avatar64,
			}}

		projectDAO.UpsertId(
			//Where
			project.Code,
			//Set
			upsertdata,
		)

		return true
	}
	return false
}

//GetProjectByCode : Retorna un Projecto por su ID
func (pm *ProjectModel) GetProjectByCode(projectCode string) Project {
	session, err := GetSession()
	defer session.Close()
	projectDAO := session.DB(DataBaseName).C(ProjectColl)
	project := Project{}
	if err == nil {
		//log.Println("CODE >>", projectCode)
		projectDAO.FindId(projectCode).One(&project)
		//log.Println(util.StringifyJSON(project))
		return project
	} else {
		log.Fatal("Error GetProjectByCode:", err)
		return project
	}
}

//SaveEpicType , Almacena tipos de Epicos
func (pm *ProjectModel) SaveEpicType(projectCode string, epicType EpicType) {
	session, err := GetSession()
	defer session.Close()
	projectClr := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("MODEL>> projectCode:", projectCode)

	err = projectClr.Find(bson.M{"_id": projectCode}).One(&projectResult)
	epicTypes := projectResult.Epics.Types
	find := false
	for index, epicTyped := range epicTypes {
		if epicTyped.Code == epicType.Code {
			find = true
			fmt.Println("UPDATE [OK]")
			projectClr.Update(
				//Where
				bson.M{"_id": projectCode},
				//Set
				bson.M{"$set": bson.M{"epics.types." + strconv.Itoa(index): epicType}})
		}
	}
	if !find {
		epicTypes = append(epicTypes, epicType)

		if err == nil {
			fmt.Println("INSERT [OK]")
			projectClr.Update(
				//Where
				bson.M{"_id": projectCode},
				//Set
				bson.M{"$set": bson.M{"epics": bson.M{"types": epicTypes}}},
			)
		} else {
			fmt.Println("SAVE [NOK]")
		}
	}
}

//GetAllEpicType , Obtiene tipos de Epicos
func (pm *ProjectModel) GetAllEpicType(projectCode string) []EpicType {
	session, err := GetSession()
	defer session.Close()
	projectClr := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("MODEL>> projectCode:", projectCode)

	err = projectClr.Find(bson.M{"_id": projectCode}).One(&projectResult)
	if err == nil {
		return projectResult.Epics.Types
	} else {
		return []EpicType{}
	}

}

//AddInvention , Guarda las sesiones asociadas al usuario
func (pm *ProjectModel) AddInvention(pid ProjectInventionData) {
	session, err := GetSession()
	defer session.Close()
	projectClr := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	err = projectClr.Find(bson.M{"_id": pid.ProjectCode}).One(&projectResult)
	allInvention := []DataInvention{}
	repository := []Repository{}

	if err == nil {
		find := false
		for _, repo := range projectResult.Repository {
			log.Println("repo.InventionCode", repo.InventionCode)
			log.Println("pid.InventionCode", pid.InventionCode)

			if repo.InventionCode == pid.InventionCode {
				find = true
				//Validar que no este repetido
				for _, data := range pid.Container {
					data.Update = time.Now()
					doCreate := true
					i := 0
					for _, preData := range repo.WareHouse {
						if preData.KeyValue == data.KeyValue {
							doCreate = false
							break
						}
						i++
					}
					if !doCreate {
						repo.WareHouse = append(repo.WareHouse[:i], repo.WareHouse[i+1:]...)
						doCreate = true
					}

					if doCreate {
						repo.WareHouse = append(repo.WareHouse, data)
					}
				}
				allInvention = repo.WareHouse
			}
			repository = append(repository, repo)
		}
		if !find {
			newRepository := Repository{}
			newRepository.LastTime = time.Now()
			dcode := util.DavinciCode{}
			newRepository.Invention = dcode.Decript(pid.InventionCode)
			newRepository.InventionCode = pid.InventionCode
			newRepository.LastTime = time.Now()
			newRepository.KeyRef = pid.KeyRef
			newRepository.WareHouse = pid.Container
			repository = append(repository, newRepository)
		}

		projectClr.Update(
			//Where
			bson.M{"_id": pid.ProjectCode},
			//Set
			bson.M{"$set": bson.M{"repository": repository}},
		)
		pid.Container = allInvention

	}

}

//GetInventions : Retorna el listado de inventos por proyecto
func (pm *ProjectModel) GetInventions(pid ProjectInventionData) []DataInvention {

	session, err := GetSession()
	defer session.Close()
	projectClr := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	err = projectClr.Find(bson.M{"_id": pid.ProjectCode}).One(&projectResult)
	dataInventions := []DataInvention{}

	exist := false

	if err == nil {
		for _, repo := range projectResult.Repository {
			log.Println("repo.InventionCode", repo.InventionCode)
			log.Println("pid.InventionCode", pid.InventionCode)

			if repo.InventionCode == pid.InventionCode {
				for _, data := range repo.WareHouse {
					dataInventions = append(dataInventions, data)
				}
				exist = true
				break
			}
		}
	}

	if !exist {

	}

	return dataInventions
}

//DelInvention , Elimina inveto hacia el repositorio de projectos.
func (pm *ProjectModel) DelInvention(pid ProjectInventionData) {
	session, err := GetSession()
	defer session.Close()
	projectClr := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	err = projectClr.Find(bson.M{"_id": pid.ProjectCode}).One(&projectResult)
	allInvention := []DataInvention{}
	repository := []Repository{}

	if err == nil {
		for _, repo := range projectResult.Repository {
			log.Println("repo.InventionCode", repo.InventionCode)
			log.Println("pid.InventionCode", pid.InventionCode)

			if repo.InventionCode == pid.InventionCode {
				//Validar que no este repetido
				for _, data := range pid.Container {
					data.Update = time.Now()
					doCreate := true
					i := 0
					for _, preData := range repo.WareHouse {
						if preData.KeyValue == data.KeyValue {
							doCreate = false
							break
						}
						i++
					}
					if !doCreate {
						repo.WareHouse = append(repo.WareHouse[:i], repo.WareHouse[i+1:]...)
					}
				}
				allInvention = repo.WareHouse
			}
			repository = append(repository, repo)
		}

		projectClr.Update(
			//Where
			bson.M{"_id": pid.ProjectCode},
			//Set
			bson.M{"$set": bson.M{"repository": repository, "lasttime": time.Now()}},
		)
		pid.Container = allInvention

	}

}

//GetAllInventions : Retorna el listado de todos los inventos por proyecto
func (pm *ProjectModel) GetAllInventions(projectCode string) []Repository {
	log.Println("Trayendo inventos desde ", projectCode)
	session, err := GetSession()
	defer session.Close()
	projectClr := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	err = projectClr.Find(bson.M{"_id": projectCode}).One(&projectResult)
	dataInventions := []Repository{}

	if err == nil {
		for _, repo := range projectResult.Repository {
			log.Println("repo.InventionCode", repo.InventionCode)
			dataInventions = append(dataInventions, repo)
		}
	}
	return dataInventions
}
