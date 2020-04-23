package models

import "log"

//ArtifactModel : Artefactos que sirven para crear inventos
type ArtifactModel struct {
}

//GetArtifactByCode : Retorna un Collection de Artifact segun su codigo
func (ar *ArtifactModel) GetArtifactByCode(code string) Artifact {
	session, err := GetSession()
	defer session.Close()
	artifactDAO := session.DB(DataBaseName).C(ArtifactColl)
	artifact := Artifact{}
	if err == nil {
		artifactDAO.FindId(code).One(&artifact)
		return artifact
	} else {
		log.Fatal("Error GetArtifactByCode:", err)
		return artifact
	}

}
