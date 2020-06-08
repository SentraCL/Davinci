package routes

import (
	"path/filepath"

	handlers "./handlers"

	"github.com/gorilla/mux"
)

// Router listados de rutas
type Router struct {
	PublicHTML string
}

//SetPublicHTML : Seteo de path con  recursos estaticos.
func (r *Router) SetPublicHTML(publicPath string) {
	r.PublicHTML, _ = filepath.Abs(publicPath)
}

// Routers , retorna los routes asociados a la rutina de login
func (r *Router) Routers() *mux.Router {
	route := mux.NewRouter()
	handlers := &handlers.Handler{}
	handlers.SetPathPublicHTML(r.PublicHTML)
	//Login
	route.HandleFunc("/api/", handlers.Status).Methods("GET")
	route.HandleFunc("/api/login/", handlers.Login).Methods("POST")
	route.HandleFunc("/api/logout/", handlers.Logout).Methods("POST")
	route.HandleFunc("/api/token/", handlers.IsValidHash).Methods("POST")
	//Project
	route.HandleFunc("/api/project/save/", handlers.SaveProject).Methods("POST")
	route.HandleFunc("/api/project/drop/", handlers.DropProject).Methods("POST")
	route.HandleFunc("/api/project/copy/", handlers.CopyProject).Methods("POST")
	route.HandleFunc("/api/project/export/", handlers.ExportProject).Methods("POST")
	route.HandleFunc("/api/project/import/", handlers.ImportProject).Methods("POST")
	//TODO : Nika, acuardate de usar este metodo cuando el archivo a exportar no tiene conflictos con el ambiente.
	route.HandleFunc("/api/project/importFile/", handlers.ImportProjectFile).Methods("POST")

	route.HandleFunc("/api/project/getAll/", handlers.GetAllProject).Methods("POST")
	route.HandleFunc("/api/project/inventions/", handlers.GetAllProjectInventions).Methods("POST")
	route.HandleFunc("/api/project/avatar/{alias}.png", handlers.GetAvatarProject)

	//User Project
	route.HandleFunc("/api/project/{project}/user/save/", handlers.SaveUserIntoProjectByFormPost).Methods("POST")
	route.HandleFunc("/api/project/{project}/user/{userName}", handlers.DeleteUserIntoProjectByFormDelete).Methods("DELETE")
	route.HandleFunc("/api/project/{project}/user/update/", handlers.UpdateUserIntoProjectByFormPut).Methods("PUT")
	/* route.HandleFunc("/api/project/{project}/user/changePassword/", handlers.UpdatePasswordUserByFormPut).Methods("PUT") */
	route.HandleFunc("/api/project/{project}/users", handlers.GetUsersForProjectByFormGet).Methods("GET")

	//Project.Repository
	route.HandleFunc("/api/project/{project}/{invention}/save", handlers.SaveInventionIntoProject).Methods("POST")
	route.HandleFunc("/api/project/{project}/{invention}/delete", handlers.DropInventionIntoProject).Methods("POST")
	route.HandleFunc("/api/project/{project}/{invention}/get", handlers.GetInventionByProject)
	//Project.Design
	route.HandleFunc("/api/project/{project}/design/userstories", handlers.SaveUserStoriesType).Methods("POST")
	route.HandleFunc("/api/project/{project}/design/userstories", handlers.GetAllUserStoriesType)

	route.HandleFunc("/api/project/{project}/design/epic", handlers.SaveEpicType).Methods("POST")
	route.HandleFunc("/api/project/{project}/design/epic", handlers.GetAllEpicType)
	//Invention
	route.HandleFunc("/api/invention/all/", handlers.GetAllInventions).Methods("GET")
	route.HandleFunc("/api/invention/saveAll/", handlers.SaveInventions).Methods("POST")
	route.HandleFunc("/api/invention/save/", handlers.SaveInvention).Methods("POST")
	route.HandleFunc("/api/invention/drop/", handlers.DropInventions).Methods("POST")
	route.HandleFunc("/api/invention/artifact/default/", handlers.GetDefaultArtifact).Methods("GET")

	//Util
	route.HandleFunc("/api/util/encript", handlers.DavinciEncode).Methods("POST")

	return route
}
