package router

import (
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/controller"
	"github.com/gorilla/mux"
)


func NewRouter(router *mux.Router, userController controller.UserController){
	router.HandleFunc("/users/register",userController.Create).Methods("POST")
}