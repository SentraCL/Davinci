
package controllers

import (
	"log"
)
type UserController struct {
}

//AddUser agrega un usuario al proyecto
func (uc *UserController) AddUser(user string, pass string, isDesign bool) bool {
	log.Println("Controller - AddUser ", user)
	return userModel.AddUser( user, pass,isDesign)
}