package controller

import (
	"log"
	"net/http"

	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/helper"
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/model/web"
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/service"
)

type UsersControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController{
	return &UsersControllerImpl{
		UserService: userService,
	}
}

func (controller *UsersControllerImpl) Create(w http.ResponseWriter,r *http.Request){
	userRequest := web.RegisterUserPayload{}
	helper.ReadFromRequestBody(r,&userRequest)
	log.Println(userRequest)

	userResponse := controller.UserService.Create(r.Context(),userRequest)

	webReponse := web.WebResponse{
		Code : 200,
		Status: "OK",
		Data: userResponse,
	}

	helper.WriteToResponseBody(w,webReponse)
}