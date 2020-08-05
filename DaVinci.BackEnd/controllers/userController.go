
package controllers

import (
	"log"
	models "../models"
)
type UserController struct {
}

//AddUser agrega un usuario al proyecto
func (uc *UserController) AddUser(user string, pass string, isDesign bool,enterprise string) bool {
	log.Println("Controller - AddUser ", user)
	return userModel.AddUser( user, pass,isDesign,enterprise)
}

//SaveUserForm agrega un usuario a la collection
func (uc *UserController) SaveUserForm(user string, pass string,enterprise []string,rol string) bool {
	log.Println("Controller - SaveUserForm ", user)
	return userModel.SaveUserForm( user, pass,enterprise,rol)
}

//SaveUserForm agrega un usuario a la collection
func (uc *UserController) EditUser(user string, pass string,enterprise []string,rol string) bool {
	log.Println("Controller - EditUser ", user)
	return userModel.EditUser( user, pass,enterprise,rol)
}

//DeleteUser elimina un usuario de la collection user
func (uc *UserController) DeleteUser(name string) bool {
	log.Println("Controller - DeleteUser ", name)
	return userModel.DeleteUser( name)
}


//findRoleUserByName retorna el rol del usuario por su nombre
func (uc *UserController) FindRoleUserByName(user string) string {
	log.Println("Controller - findRoleUserByName ", user)
	return userModel.FindRoleUserByName(user)
}

//GetAllDomainUsers retorna todos los usuarios de dominio
func (uc *UserController) GetAllDomainUsers() []models.UserVO {
	log.Println("Controller - GetAllDomainUsers ", )
	return userModel.GetAllDomainUsers()
}

