package models

import (
	"fmt"
	"log"
	"strconv"
	"time"

	util "../util"
	"gopkg.in/mgo.v2/bson"
)

//AddUser , Guarda los nuevos usuarios de un proyecto
func (pm *ProjectModel) AddUser(projectCode string, user string, pass string, isDesign bool) {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	if err == nil {
		dcode := util.DavinciCode{}
		var user = bson.M{"_id": dcode.Encript(user), "user": user, "password": pass, "isDesign":isDesign}
		projectCollector.Update(
			//Where
			bson.M{"_id": projectCode},
			//Add
			bson.M{"$addToSet": bson.M{"users": user}},
		)
	}
}

//DelUser , Elimina un usuario de un proyecto
func (pm *ProjectModel) DelUser(projectCode string, user string) {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	if err == nil {
		log.Println("Encontrado")
		log.Println(projectResult.Users)
		dataUsers := []UserProject{}

		for _, repo := range projectResult.Users {
			log.Println("UserName", repo.UserName)
			if repo.Code != user {
				dataUsers = append(dataUsers, repo)
			}
		}
		log.Println("final : ", dataUsers)
		projectCollector.Update(
			//Where
			bson.M{"_id": projectCode},
			//Add
			bson.M{"$set": bson.M{"users": dataUsers}},
		)
	}
}

//UpdateUser , Modifica un usuarios de un proyecto
func (pm *ProjectModel) UpdateUser(projectCode string, code string, user string, pass string) {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	if err == nil {
		dataUsers := []UserProject{}

		for _, repo := range projectResult.Users {
			log.Println("UserName", repo.UserName)
			if repo.Code == code {
				repo.UserName = user
				repo.Password = pass
			}
			dataUsers = append(dataUsers, repo)
		}
		projectCollector.Update(
			//Where
			bson.M{"_id": projectCode},
			//Add
			bson.M{"$set": bson.M{"users": dataUsers}},
		)
	}
}

//LoginUser , Modifica un usuarios de un proyecto
func (pm *ProjectModel) LoginUser(projectName string, user string, pass string) (bool,bool) {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	err = projectCollector.Find(bson.M{"alias": projectName}).One(&projectResult)
	var result = false
	var isDesign = false
	if err == nil {
		for _, repo := range projectResult.Users {
			if repo.UserName == user {
				if repo.IsDesign==false{
					if repo.Password==pass {
						result = true
					}
				}else{
					isDesign=true
				}
			}
		}
	}
	return result,isDesign
}

//GetUsers retorna usuarios de un proyecto
func (pm *ProjectModel) GetUsers(projectCode string) []UserProject {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	dataUsers := []UserProject{}
	if err == nil {
		for _, repo := range projectResult.Users {
			log.Println("UserName", repo.UserName)
			user := UserProject{}
			user.UserName = repo.UserName
			user.Code = repo.Code
			dataUsers = append(dataUsers, user)
		}
	}
	return dataUsers
}

//SaveUserStoriesType , Almacena Tipos de Usuario
func (pm *ProjectModel) SaveUserStoriesType(projectCode string, usType UserStoriesType) {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("MODEL>> projectCode:", projectCode)

	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	usTypes := projectResult.UserStories.Types
	find := false
	for index, usTyped := range usTypes {
		if usTyped.Code == usType.Code {
			find = true
			//fmt.Println("UPDATE [OK]")
			projectCollector.Update(
				//Where
				bson.M{"_id": projectCode},
				//Set
				bson.M{"$set": bson.M{"userStories.types." + strconv.Itoa(index): usType}})
		}
	}
	if !find {
		usTypes = append(usTypes, usType)

		if err == nil {
			//fmt.Println("INSERT [OK]")
			projectCollector.Update(
				//Where
				bson.M{"_id": projectCode},
				//Set
				bson.M{"$set": bson.M{"userStories.types": usTypes}},
			)
		} else {
			//fmt.Println("SAVE [NOK]")
		}
	}
}

//GetAllUserStories , Almacena tipos de historias de usuario.
func (pm *ProjectModel) GetAllUserStories(projectCode string) []UserStoriesType {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("MODEL>> projectCode:", projectCode)

	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	if err == nil {
		return projectResult.UserStories.Types
	} else {
		return []UserStoriesType{}
	}

}

func (pm *ProjectModel) getSequenceByRef(projectCode string, codeValue string) string {
	session, _ := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("MODEL>> projectCode:", projectCode)

	projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	usRepository := projectResult.UserStories.Repository
	//TODO : Es muy charcha iterar para sacar el resultado.
	count := 1
	for _, usDB := range usRepository {
		if usDB.Code.Value == codeValue {
			count++
		}
	}
	sequence := fmt.Sprintf("%04d", count)
	return sequence
}

//SaveUserStory , Almacena Tipos de Usuario
func (pm *ProjectModel) SaveUserStory(projectCode string, us UserStory) UserStory {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("MODEL>> projectCode:", projectCode)

	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	usRepository := projectResult.UserStories.Repository

	if us.ID != "" {
		for index, usDB := range usRepository {

			us.Versions[len(us.Versions)-1].Date = time.Now()
			if usDB.ID == us.ID {

				if us.LastVersionIndex+1 == len(us.Versions) {
					//Actualizo fecha, solo se puede modificar la ultima version de una Historia de Usuario
					//fmt.Println("UPDATE [OK] >>", us.ID)
				} else {
					//Aumento la Version, considerando que el arreglo trae una nueva version.
					us.LastVersionIndex = us.LastVersionIndex + 1
				}
				projectCollector.Update(
					//Where
					bson.M{"_id": projectCode},
					//Set
					bson.M{"$set": bson.M{"userStories.repository." + strconv.Itoa(index): us}})
			}
		}
	} else {
		sequence := pm.getSequenceByRef(projectCode, us.Code.Value) //fmt.Sprintf("%04d", len(projectResult.UserStories.Repository)+1)
		us.ID = us.Code.Value + sequence
		us.Versions[0].Date = time.Now()
		usRepository = append(usRepository, us)
		if err == nil {
			//fmt.Println("INSERT [OK] =>", us.ID)
			projectCollector.Update(
				//Where
				bson.M{"_id": projectCode},
				//Set
				//bson.M{"$set": bson.M{"userStories": bson.M{"repository": usRepository}}},
				//bson.M{"$set": bson.M{"userStories.repository." + strconv.Itoa(index): us}})
				bson.M{"$set": bson.M{"userStories.repository": usRepository}})

		} else {
			//fmt.Println("SAVE [NOK]")
		}
	}
	return us
}

//GetUserStoryByCode , trae todos los casos similares
func (pm *ProjectModel) GetUserStoryByCode(projectCode string, codeValue string) []UserStory {
	session, _ := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("MODEL>> projectCode:", projectCode)
	//fmt.Println("MODEL>> codeValue:", codeValue)

	projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	usRepository := projectResult.UserStories.Repository
	userStories := []UserStory{}

	for _, usDB := range usRepository {
		if usDB.ID == codeValue {
			userStories = append(userStories, usDB)
		}
	}
	return userStories
}

/*
//GetUserStoryByCodeVersion , Trae User Story segun version y codigo
func (pm *ProjectModel) GetUserStoryByCodeVersion(projectCode string, uscode string, indexVersion int) {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	err = projectCollector.Find(bson.M{"_id": projectCode}).Select(bson.M{"userStories.repository": bson.M{"$elemMatch": bson.M{"id": uscode}}}).One(&projectResult)

	if err == nil {
		//fmt.Println("MODEL>> projectCollector:", util.StringifyJSON(projectResult))
	} else {
		//fmt.Println(err)
	}


		err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
		if err == nil {

		}

}
*/

//GetAllUserStoriesbyPreconditions , trae todos los casos similares
func (pm *ProjectModel) GetAllUserStoriesbyPreconditions(projectCode string, queryPC []PreConditionQuery) []UserStory {
	session, _ := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	usRepository := projectResult.UserStories.Repository
	userStories := []UserStory{}
	ids := []string{}
	count := 0
	for _, usDB := range usRepository {
		if pm.index(ids, usDB.ID) == -1 {
			count = 0
			for _, queryPreCondition := range queryPC {
				for _, precondition := range usDB.Versions[usDB.LastVersionIndex].PreConditions {
					if precondition.ID == queryPreCondition.ID && precondition.Value == queryPreCondition.Value {
						count++
					}
				}
			}
			if count == len(queryPC) {
				userStories = append(userStories, usDB)
				ids = append(ids, usDB.ID)
			}

		}
	}
	return userStories
}

func (pm *ProjectModel) index(slice []string, item string) int {
	for i, _ := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

//SaveEpic , Guarda cualquier cambio en el proyecto
func (pm *ProjectModel) SaveEpic(projectCode string, epic *Epic) bool {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("ID=", epic.ID)

	//fmt.Println("MODEL>> projectCode:", projectCode)

	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)

	if err == nil {
		//fmt.Println("MODEL>> Epics:", util.StringifyJSON(projectResult.Epics))
		epicRepository := projectResult.Epics.Repository
		fmt.Println("MODEL>> epic.ID	:", epic.ID)
		fmt.Println("MODEL>> epic.Code	:", epic.Code)
		epicId:=epic.ID.Hex()
		if  epicId == "" {
			position:=len(projectResult.Epics.Repository)
			epic.ID = bson.NewObjectId()
			epic.Date = time.Now()
			if err == nil {
				fmt.Println(epic.ID, " NEW [OK] =>", epic.Code)
				projectCollector.Update(
					//Where
					bson.M{"_id": projectCode},
					//Set
					bson.M{"$set": bson.M{"epics.repository."+strconv.Itoa(position): epic}})
				epic.Save = true
			} else {
				fmt.Println("SAVE [NOK]", err)
				return false
			}
		} else {
			fmt.Println(epic.ID, " Old [OK] =>", epic.Code)
			for index, epicDB := range epicRepository {
				epic.Date = time.Now()
				if epicDB.ID == epic.ID {
					//fmt.Println(epic.ID, " UPDATE [OK] =>", epic.Code)
					projectCollector.Update(
						//Where
						bson.M{"_id": projectCode},
						//Set
						bson.M{"$set": bson.M{"epics.repository." + strconv.Itoa(index): epic}})
					epic.Save = true
				}
			}
		}
		return true
	}
	fmt.Println("SAVE [NOK]", err)
	return false
}

//GetAllEpics , Retorna todos los Epicos
func (pm *ProjectModel) GetAllEpics(projectCode string) []Epic {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("MODEL>> projectCode:", projectCode)

	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	if err == nil {
		return projectResult.Epics.Repository
	} else {
		return []Epic{}
	}

}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

//SaveTestData, Guarda toda la informacion no duplicada
func (pm *ProjectModel) SaveRepositoryData(projectCode string, dataType *TestData) string {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}
	dcode := util.DavinciCode{}
	id:=dcode.Encript(projectCode)
	err = projectCollector.Find(bson.M{"_id": id}).One(&projectResult)

	if err == nil {
		dataTypes := projectResult.DataStored.Repository
		position:=0
		encontrado:=false
			for index, dataTypeDB := range dataTypes {
				position=position+1
				dataType.Date = time.Now()
				fmt.Println(index)
				if dataTypeDB.ID == dataType.ID {
					fmt.Println(dataTypeDB)
					fmt.Println(dataType)
					newRepo := dataTypeDB.Atributes	
					insertados:=0
					for _,atrib:=range dataType.Atributes{
						fmt.Println(atrib)
						if(!stringInSlice(atrib,newRepo)){
							newRepo=append(newRepo,atrib)
							insertados++
						}
					}
					dataType.Atributes=newRepo
					encontrado=true
					if(insertados>0){
						projectCollector.Update(
							//Where
							bson.M{"_id": id},
							//Set
							bson.M{"$set": bson.M{"dataStored.repository." + strconv.Itoa(index): dataType}})
							return "Se han agregado "+strconv.Itoa(insertados)+" datos para el tipo de dato existente."
						}else{
						return "No se pueden agregar datos duplicados."
					}
				}
			}				
				if(position==0||encontrado==false){
					dataType.Date = time.Now()
					projectCollector.Update(
						//Where
						bson.M{"_id": id},
						//Set
						bson.M{"$set": bson.M{"dataStored.repository."+strconv.Itoa(position): dataType}})
						return "Se ha ingresado informacion para un nuevo tipo de dato."
				}
	}

	//fmt.Println("SAVE [NOK]", err)
	return "No se pudo agregar la informacion ingresada."
}

//SaveDataType , Guarda cualquier cambio en el proyecto
func (pm *ProjectModel) SaveDataType(projectCode string, dataType *DataType) bool {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("ID=", dataType.ID)

	//fmt.Println("MODEL>> projectCode:", projectCode)
	
	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)

	if err == nil {
		dataTypes := projectResult.DataStored.DataTypes
		if dataType.ID == "" {
			if dataType.ID == "" {
				dcode := util.DavinciCode{}
				dataType.ID = dcode.Encript(dataType.Name)
			}
			dataType.Date = time.Now()
			if err == nil {
				newRepo := []DataType{}
				newRepo = append(dataTypes, *dataType)
				//fmt.Println(dataType.ID, " NEW [OK] =>", dataType.ID)
				projectCollector.Update(
					//Where
					bson.M{"_id": projectCode},
					//Set
					bson.M{"$set": bson.M{"dataStored.dataTypes": newRepo}})
			} else {
				//fmt.Println("SAVE [NOK]")
				return false
			}
		} else {
			for index, dataTypeDB := range dataTypes {
				dataType.Date = time.Now()
				if dataTypeDB.ID == dataType.ID {
					//fmt.Println(dataType.ID, " UPDATE [OK] =>", dataType.ID)
					projectCollector.Update(
						//Where
						bson.M{"_id": projectCode},
						//Set
						bson.M{"$set": bson.M{"dataStored.dataTypes." + strconv.Itoa(index): dataType}})
				}
			}
		}
		return true
	}
	//fmt.Println("SAVE [NOK]", err)
	return false
}

//GetAllDataTypes , retorna todos los tipos de datos segun el codigo de un proyecto
func (pm *ProjectModel) GetAllDataTypes(projectCode string) []DataType {
	session, err := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("MODEL>> projectCode:", projectCode)

	err = projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	if err == nil {
		return projectResult.DataStored.DataTypes
	} else {
		return []DataType{}
	}

}

//GetDataTypeByID , retorna un tipo de datos segun ID
func (pm *ProjectModel) GetDataTypeByID(projectCode string, idValue string) DataType {
	session, _ := GetSession()
	defer session.Close()
	projectCollector := session.DB(DataBaseName).C(ProjectColl)
	projectResult := Project{}

	//fmt.Println("MODEL>> projectCode:", projectCode)
	//fmt.Println("MODEL>> ID:", idValue)

	projectCollector.Find(bson.M{"_id": projectCode}).One(&projectResult)
	usRepository := projectResult.DataStored.DataTypes
	dataType := DataType{}
	//TODO : Ver forma de evitar iterar para encontrar, Seguramente mongodb tiene algo mejor
	for _, dataTypeDB := range usRepository {
		if dataTypeDB.ID == idValue {
			dataType = dataTypeDB
		}
	}
	return dataType
}
