package routes

import (
	handlers "./handlers"

	"github.com/gorilla/mux"
)

// SubRouters , retorna los routes asociados a la rutina de login
func (r *Router) SubRouters() *mux.Router {
	route := mux.NewRouter()
	handlers := &handlers.Handler{}
	handlers.SetPathPublicHTML(r.PublicHTML)

	//SubDomais
	route.HandleFunc("/davinci/{alias}", handlers.SubDomain).Methods("GET")
	//SubDomains
	route.HandleFunc("/davinci/{project}/", handlers.SubDomain).Methods("GET")
	route.HandleFunc("/davinci/{project}/logIn", handlers.DoLogin).Methods("POST")
	route.HandleFunc("/davinci/{project}/logOut", handlers.DoLogOut).Methods("GET")
	route.HandleFunc("/davinci/{project}/isOnline", handlers.IsOnline).Methods("GET")
	route.HandleFunc("/davinci/{project}/user/changePassword/", handlers.UpdatePasswordUserByFormPut).Methods("PUT")
	route.HandleFunc("/davinci/{project}/addStep", handlers.AddStep).Methods("POST")
	route.HandleFunc("/davinci/token/{projectCode}", handlers.IsValidSubHash).Methods("POST")

	//Context
	route.HandleFunc("/davinci/{project}/epicTypes", handlers.GetEpicTypesProject).Methods("GET")
	route.HandleFunc("/davinci/{project}/userStories", handlers.GetAllUserStoriesTypeProject).Methods("GET")
	route.HandleFunc("/davinci/{project}/userStories", handlers.SaveUserStoriesProject).Methods("POST")
	route.HandleFunc("/davinci/{project}/userStories/ref/", handlers.GetUserStoryByRef).Methods("POST")
	route.HandleFunc("/davinci/{project}/userStories/{code}", handlers.GetUserStoryByCode).Methods("GET")
	route.HandleFunc("/davinci/{project}/userStories/similar", handlers.GetAllUserStoriesbyPreconditions).Methods("POST")

	//Data	
	route.HandleFunc("/davinci/{project}/datatype/save/", handlers.SaveDataType).Methods("POST")
	route.HandleFunc("/davinci/{project}/datatype/getAll/", handlers.GetAllDataTypes).Methods("GET")
	route.HandleFunc("/davinci/{project}/datatype/get/{id}/", handlers.GetDataTypeByID).Methods("GET")
	route.HandleFunc("/davinci/{project}/datatype/{id}/icon.png", handlers.GetIconDataTypeByID).Methods("GET")
	
	route.HandleFunc("/davinci/{project}/repository/save/", handlers.SaveRepositoryData).Methods("POST")
	
	//Project
	route.HandleFunc("/davinci/{project}/code", handlers.GetCode).Methods("GET")
	route.HandleFunc("/davinci/{project}/epics", handlers.GetEpicsProject).Methods("GET")
	route.HandleFunc("/davinci/typeref/{ref}/", handlers.GetTypeRef).Methods("GET")
	//route.HandleFunc("/davinci/invention/{project}/{ref}/", handlers.GetInventionByProject).Methods("GET")
	route.HandleFunc("/davinci/invention/{project}/all/", handlers.GetAllProjectInventions).Methods("GET")

	route.HandleFunc("/davinci/{project}/epic/export/{type}", handlers.ExportEpicProject).Methods("POST")
	route.HandleFunc("/davinci/{project}/epic", handlers.SaveEpicProject).Methods("POST")


	return route
}
