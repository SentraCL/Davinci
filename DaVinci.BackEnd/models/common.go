package models

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/mgo.v2"

	util "../util"
)

//CopyProjectRequest , Request que representa la copia de un proyecto
type CopyProjectRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Users       int    `json:"users"`
	Epics       int    `json:"epics"`
	UserStories int    `json:"userStories"`
	Data        int    `json:"data"`
}

//ExportProjectRequest , Request que representa exportacion de un proyecto
type ExportProjectRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Users       int    `json:"users"`
	Epics       int    `json:"epics"`
	UserStories int    `json:"userStories"`
	Data        int    `json:"data"`
	Avatar		int    `json:"avatar"`
}

//ExecuteResponse , Retonar mensajes hacia front-end
type ExecuteResponse struct {
	StatusCode string      `json:"status"`
	Details    interface{} `json:"datails"`
}

type ResponseInvention struct {
	ResumeInvention map[string]string
}

//ProjectInventionData , Data Tranfer Object de Guardado de Inventos por Proyecto
type ProjectInventionData struct {
	ProjectCode   string
	InventionCode string
	KeyRef        string
	Container     []DataInvention
}

//Repository :  Repositorio de inventos con contenido.
type RepositoryRef struct {
	InventionCode string    `json:"_id"`
	Invention     string    `json:"name"`
	KeyRef        string    `json:"key"`
	LastTime      time.Time `json:"lastTime"`
}

//ProjectRequest : Representacion de un Projecto desde Front-END
type ProjectRequest struct {
	Code       string          `json:"code"`
	Name       string          `json:"name"`
	Alias      string          `json:"alias"`
	Enterprise    string          `json:"enterprise"`
	Admin      string          `json:"admin"`
	Email      string          `json:"email"`
	Resume     string          `json:"resume"`
	Avatar64   string          `json:"avatar"`
	Repository []RepositoryRef `json:"repository"`
	Users      []UserProject   `bson:"users"`
}

//Login , Credenciales de usuario
type Login struct {
	User        string
	Pass        string
	DavinciCode string
	RemoteAddr  string
	Online      bool
}

//MakeDavinciCODE : Crear Hash de idenfiticacion de sesion
func (l *Login) MakeDavinciCODE(cookieHash string) string {
	//Reemplazar posteriormento l.User por algo mas complejo
	dcode := util.DavinciCode{}
	encript := dcode.Encript(cookieHash + ":" + l.User)
	l.DavinciCode = encript
	return encript
}

//DataBaseName :Nombre de la Base de Datos
const DataBaseName = "DAVINCI"

const UserColl = "User"
const ProjectColl = "Project"
const EnterpriseColl = "Enterprise"
const ArtifactColl = "Artifact"
const InventionColl = "Invention"
const TypeRefColl = "TypeRef"

// RunMongod is a simple wrapper around terminal commands
func RunMongod(path string, args []string, debug bool) (out string, err error) {

	cmd := exec.Command(path, args...)

	var b []byte
	b, err = cmd.CombinedOutput()
	out = string(b)

	if debug {
		fmt.Println(strings.Join(cmd.Args[:], " "))

		if err != nil {
			fmt.Println("RunMongod ERROR")
			fmt.Println(out)
		}
	}

	return
}

func createAdmin() bool {
	session, err := GetSession()
	defer session.Close()

	userDB := session.DB(DataBaseName).C("User")

	admin := User{}
	//user, pass, email
	admin.UserName = "admin"
	admin.Password = "admin"

	err = userDB.Insert(admin)
	if err != nil {
		return false
	} else {
		return true
	}
}

//WakeUPDataBase : Levanta base de datos
func WakeUPDataBase(datapath string) bool {
	dbPath, _ := filepath.Abs(datapath)
	doCreateAdmin := false
	//TODO : Preguntar mejor si la base de datos Existe en vez del directorio
	if !util.ExistsPath(dbPath) {
		os.MkdirAll(dbPath, os.ModePerm)
		doCreateAdmin = true
	}
	args := []string{"-dbpath", dbPath}
	go RunMongod("mongod", args, true)

	if doCreateAdmin {
		status := createAdmin()
		fmt.Println("Create Admin ", status)
		typeRef := TypeRefModel{}
		typeRef.CreateBasicTypeRef()
		fmt.Println("Create Default Types [OK]")
	}

	return true
}

//GetSession , Obtener session
func GetSession() (*mgo.Session, error) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session, err
}
