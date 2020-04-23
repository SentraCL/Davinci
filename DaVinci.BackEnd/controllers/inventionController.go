package controllers

import (
	model "../models"
	util "../util"
)

//InventionController : Inventor!
type InventionController struct {
	dcode util.DavinciCode
}

//GetAvatarByAlias : Upsert Project!!
func (pc *InventionController) GetInventionByCode(code string) model.Invention {
	invention := inventionModel.GetInventionByCode(code)
	return invention
}

func (ic *InventionController) getTypeName(artifactVO *model.ArtifactVO) {
	artifactVO.IsJSON = artifactVO.TypeName == "json" || !artifactVO.IsEssential
	artifactVO.IsText = artifactVO.TypeName == "text"
	artifactVO.IsNumeric = artifactVO.TypeName == "numeric"
	artifactVO.IsDate = artifactVO.TypeName == "date"
	artifactVO.IsBool = artifactVO.TypeName == "bool"
}

func (ic *InventionController) GetTypeRefByCode(codeTypeRef string) model.TypeRef {
	return typeRefModel.GetTypeRefByCode(codeTypeRef)
}


func (ic *InventionController) getTypeRef(artifactVO model.ArtifactVO) model.TypeRef {
	if !artifactVO.IsJSON {
		codeTypeRef := ic.dcode.Encript(artifactVO.TypeName)
		return typeRefModel.GetTypeRefByCode(codeTypeRef)
	} else {
		//artifactCode := ic.dcode.Encript(artifactVO.Name)
		//artifact := artifactModel.GetArtifactByCode(artifactCode)
		artifactTypeRef := model.TypeRef{}
		artifactTypeRef.IsList = artifactVO.IsList
		artifactTypeRef.Code = ic.dcode.Encript(artifactVO.Name)
		artifactTypeRef.Name = artifactVO.TypeName
		artifactTypeRef.IsEssential = false
		return artifactTypeRef
	}
}

func (ic *InventionController) translateRequestToBO(vo model.InventionVO) model.Invention {
	invBO := model.Invention{}
	invBO.Code = vo.Code
	invBO.Name = vo.Name
	invBO.KeyLabel = vo.KeyLabel
	invBO.KeyValue = vo.KeyValue
	invBO.Icon = vo.Icon

	for _, artifactVO := range vo.Artifacts {
		artifact := model.Artifact{}
		ic.getTypeName(&artifactVO)
		artifact.Code = ic.dcode.Encript(artifactVO.NickName)
		artifact.Name = artifactVO.Name

		artifact.NickName = artifactVO.NickName
		artifact.TypeRef = ic.getTypeRef(artifactVO)

		//log.Println("artifactVO : " + util.StringifyJSON(artifactVO))
		//log.Println("artifact : " + util.StringifyJSON(artifact))
		invBO.Artifacts = append(invBO.Artifacts, artifact)
	}
	return invBO
}

//SaveAll : Guarda todos los ivnentos
func (ic *InventionController) SaveAll(InventionVO []model.InventionVO) []model.InventionVO {
	inventionCheck := []model.InventionVO{}

	for _, invention := range InventionVO {
		invention.Code = ic.dcode.Encript(invention.Name)
		inventionBO := ic.translateRequestToBO(invention)
		//log.Println(util.StringifyJSON(inventionBO))
		inventionModel.Save(&inventionBO)
		inventionCheck = append(inventionCheck, invention)
	}
	return inventionCheck
}

//Save : Guarda un invento.
func (ic *InventionController) Save(invention model.InventionVO) model.InventionVO {
	invention.Code = ic.dcode.Encript(invention.Name)
	inventionBO := ic.translateRequestToBO(invention)
	//log.Println("GUARDANDO : " + util.StringifyJSON(inventionBO))
	inventionModel.Save(&inventionBO)
	return invention
}

func (ic *InventionController) translateBOToRequest(invention model.Invention) model.InventionVO {
	inVO := model.InventionVO{}
	inVO.Name = invention.Name
	inVO.Code = invention.Code

	inVO.KeyLabel = invention.KeyLabel
	inVO.KeyValue = invention.KeyValue

	inVO.Title = invention.Name
	inVO.Subtitle = invention.Name
	inVO.Slots = invention.Code
	inVO.Icon = invention.Icon

	for _, artifact := range invention.Artifacts {
		artifactVO := model.ArtifactVO{}

		//log.Println("artifact.Code : " + artifact.Code)
		//log.Println("artifact.Name : " + artifact.Name)
		//log.Println("artifact.NickName : " + artifact.NickName)
		//log.Println("artifact.TypeRef : " + util.StringifyJSON(artifact.TypeRef))

		artifactVO.Name = artifact.Name
		artifactVO.Code = artifact.TypeRef.Code
		artifactVO.TypeName = artifact.TypeRef.Name
		code := ic.dcode.Encript(artifact.TypeRef.Name)
		typeRef := typeRefModel.GetTypeRefByCode(code)

		//log.Println("typeRef.Name : " + typeRef.Name)
		//log.Println("typeRef.Code : " + typeRef.Code)

		artifactVO.TypeRef = artifact.Code
		artifactVO.IsEssential = typeRef.IsEssential
		artifactVO.IsList = artifact.TypeRef.IsList
		artifactVO.IsJSON = !typeRef.IsEssential
		ic.getTypeName(&artifactVO)
		artifactVO.NickName = artifact.NickName
		inVO.Artifacts = append(inVO.Artifacts, artifactVO)
	}
	return inVO
}

//GetAll : trae todos los ivnentos
func (ic *InventionController) GetAll() []model.InventionVO {
	inventions := inventionModel.GetAll()
	inventionVOs := []model.InventionVO{}
	for _, invention := range inventions {
		inventionVO := ic.translateBOToRequest(invention)
		inventionVOs = append(inventionVOs, inventionVO)
	}
	return inventionVOs
}

//GetDefaultArtifact : trae defaultArtifact
func (ic *InventionController) GetDefaultArtifact() []model.ArtifactVO {
	typeRefs := typeRefModel.GetAll()
	artifactVOs := []model.ArtifactVO{}
	for _, typeRef := range typeRefs {
		if typeRef.IsEssential {
			artifactVO := model.ArtifactVO{}
			artifactVO.Code = typeRef.Code
			artifactVO.TypeName = typeRef.Name
			artifactVO.IsEssential = true
			artifactVOs = append(artifactVOs, artifactVO)
		}
	}
	return artifactVOs
}

//Drop : Guarda todos los ivnentos
func (ic *InventionController) Drop(inventionVO model.InventionVO) model.ExecuteResponse {
	response := model.ExecuteResponse{}
	dcode := util.DavinciCode{}
	code := dcode.Encript(inventionVO.Name)

	inventions := inventionModel.GetAll()
	toDrop := true
	parents := []string{}

	for _, invention := range inventions {
		if invention.Code != code {
			for _, artifact := range invention.Artifacts {
				if artifact.Code == code {
					toDrop = false
					parents = append(parents, invention.Name)
				}
			}
		}
	}
	if toDrop {
		isDrop, err := inventionModel.DeleteByCode(code)
		if err != nil {
			response.StatusCode = "Error:DB"
			response.Details = "Message:" + err.Error()
		}
		if isDrop {
			response.StatusCode = "OK"
			response.Details = "Success"
		}
	} else {

		var parentsDatail struct {
			Parents []string `json:"parents"`
		}

		parentsDatail.Parents = parents
		response.StatusCode = "Error:Used"
		response.Details = parentsDatail
	}
	return response
}
