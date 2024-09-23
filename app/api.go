package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/controller"
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/repository/users"
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/router"
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer{
	return &ApiServer{
		addr: addr,
		db : db,
	}
}

func (s *ApiServer) Run() error {
	route:= mux.NewRouter()
	subrouter := route.PathPrefix("/api/v1").Subrouter()
	validate := validator.New()


	// user-router
	userRepostirory := users.NewUserRepository(s.db)
	userService := service.NewHandler(userRepostirory,validate)
    userController := controller.NewUserController(userService)
	router.NewRouter(subrouter,userController)


	log.Println("Server Lister On Port" + s.addr)
	return http.ListenAndServe(s.addr,route)
}