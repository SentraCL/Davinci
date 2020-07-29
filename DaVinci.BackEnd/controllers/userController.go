
package controllers

import (
	"log"
)
type UserController struct {
}

//AddUser agrega un usuario al proyecto
func (uc *UserController) AddUser(user string, pass string, isDesign bool,enterprise string) bool {
	log.Println("Controller - AddUser ", user)
	return userModel.AddUser( user, pass,isDesign,enterprise)
}


//findRoleUserByName retorna el rol del usuario por su nombre
func (uc *UserController) FindRoleUserByName(user string) string {
	log.Println("Controller - findRoleUserByName ", user)
	return userModel.FindRoleUserByName(user)
}