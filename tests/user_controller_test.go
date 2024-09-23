package tests

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/controller"
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/repository/users"
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/router"
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql","root:iwan@tcp(localhost:3306)/go_shop_api")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setUpRouter(db *sql.DB) http.Handler {
	route:= mux.NewRouter()
	subrouter := route.PathPrefix("/api/v1").Subrouter()
	validate := validator.New()

	userRepostirory := users.NewUserRepository(db)
	userService := service.NewHandler(userRepostirory,validate)
    userController := controller.NewUserController(userService)
	router.NewRouter(subrouter,userController)

	return route
}

func truncateUsers(db *sql.DB){
	db.Exec("TRUNCATE users")
}

func TestCreateUserSuccess(t *testing.T){
	db := setupTestDB()
	truncateUsers(db)

	router := setUpRouter(db)

	requestBody := strings.NewReader(`{ "first_name" : "test",
    "last_name" : "test",
    "email" : "test@mail.com",
    "password" : "test123"}`)

	request := httptest.NewRequest(http.MethodPost,"http://localhost:8080/api/v1/users/register",requestBody)
	request.Header.Add("Content-Type","application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body,&responseBody)

	assert.Equal(t, 200 , int(responseBody["code"].(float64)))
	assert.Equal(t,"OK",responseBody["status"])
	assert.Equal(t,"test",responseBody["data"].(map[string]any)["first_name"])
}



