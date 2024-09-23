package controller

import "net/http"

type UserController interface {
	Create(w http.ResponseWriter,r *http.Request)
} 