package models

import (
	"log"

	util "../util"
)

//TypeRefModel : Artefactos que sirven para crear inventos
type TypeRefModel struct {
}

//GetTypeRefByCode : Retorna un Collection de TypeRef segun su codigo
func (tr *TypeRefModel) GetTypeRefByCode(code string) TypeRef {
	session, err := GetSession()
	defer session.Close()
	typeRefDAO := session.DB(DataBaseName).C(TypeRefColl)
	typeRef := TypeRef{}
	if err == nil {
		typeRefDAO.FindId(code).One(&typeRef)
		return typeRef
	} else {
		log.Fatal("Error GettypeRefByCode:", err)
		return typeRef
	}

}

//CreateBasicTypeRef , Artefacto por defecto
func (tr *TypeRefModel) CreateBasicTypeRef() bool {
	session, err := GetSession()
	defer session.Close()

	typeRefColl := session.DB(DataBaseName).C(TypeRefColl)

	typeText := tr.createTypeRef("text", true)
	typeNumber := tr.createTypeRef("numeric", true)
	typeDate := tr.createTypeRef("date", true)
	typeBool := tr.createTypeRef("bool", true)
	typeMultiText := tr.createTypeRef("multitext", true)

	for _, typeBasic := range [5]TypeRef{typeText, typeNumber, typeDate, typeBool, typeMultiText} {
		err = typeRefColl.Insert(typeBasic)
		if err != nil {
			log.Fatal(err)
		}
	}
	return true
}

func (tr *TypeRefModel) createTypeRef(name string, isEssential bool) TypeRef {
	typeRef := TypeRef{}
	typeRef.Name = name
	typeRef.IsEssential = isEssential
	typeRef.IsList = false
	dcode := util.DavinciCode{}
	typeRef.Code = dcode.Encript(name)
	return typeRef
}

//GetAll , Retorna todos los proyectos
func (tr *TypeRefModel) GetAll() []TypeRef {
	session, err := GetSession()
	defer session.Close()
	typeRefDAO := session.DB(DataBaseName).C(TypeRefColl)
	var results []TypeRef
	err = typeRefDAO.Find(nil).All(&results)
	if err != nil {
		log.Fatal("Error al obtener los proyectos .", err)
	}
	return results
}
