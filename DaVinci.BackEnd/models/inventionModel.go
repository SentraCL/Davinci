package models

import (
	"log"
	"time"

	util "../util"
	"gopkg.in/mgo.v2/bson"
)

// InventionModel : Acciones sobre collection de invention
type InventionModel struct {
}

//GetAll , Retorna todos los proyectos
func (pm *InventionModel) GetAll() []Invention {
	session, err := GetSession()
	defer session.Close()
	inventionDAO := session.DB(DataBaseName).C(InventionColl)
	var results []Invention
	err = inventionDAO.Find(nil).All(&results)
	if err != nil {
		log.Fatal("Error al obtener los proyectos .", err)
	}
	return results
}

//DeleteByCode , Borra segun codigo.
func (pm *InventionModel) DeleteByCode(code string) (bool, error) {
	session, err := GetSession()
	defer session.Close()
	inventionDAO := session.DB(DataBaseName).C(InventionColl)
	//log.Println("response :", util.StringifyJSON(inventionResult))
	if err == nil {
		inventionDAO.Remove(bson.M{"_id": code})
		return true, nil
	}
	return false, err
}

//Save , Guarda cualquier cambio en el proyecto
func (pm *InventionModel) Save(invention *Invention) bool {
	session, err := GetSession()
	defer session.Close()
	inventionDAO := session.DB(DataBaseName).C(InventionColl)
	//log.Println("response :", util.StringifyJSON(inventionResult))
	if err == nil {
		invention.Date = time.Now()
		dcode := util.DavinciCode{}
		invention.Code = dcode.Encript(invention.Name)
		log.Println("Invention Code >>", invention.Code)
		upsertdata := bson.M{"$set": invention}
		inventionDAO.UpsertId(
			//Where
			invention.Code,
			//Set
			upsertdata,
		)

		return true
	}
	return false
}

//GetInventionByCode : Retorna un Inventiono por su ID
func (pm *InventionModel) GetInventionByCode(inventionCode string) Invention {
	session, err := GetSession()
	defer session.Close()
	inventionDAO := session.DB(DataBaseName).C(InventionColl)
	invention := Invention{}
	if err == nil {
		//log.Println("CODE >>", inventionCode)
		inventionDAO.FindId(inventionCode).One(&invention)
		//log.Println(util.StringifyJSON(invention))
		return invention
	} else {
		log.Fatal("Error GetInventionByCode:", err)
		return invention
	}
}
