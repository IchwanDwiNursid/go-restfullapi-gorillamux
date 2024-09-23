package service

import (
	"context"

	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/model/domain"
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/model/web"
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/repository/users"
	"github.com/go-playground/validator/v10"
)


type Handler struct {
	Repository users.UserRepository
	Validate *validator.Validate
}

func NewHandler(repository users.UserRepository,validate *validator.Validate) UserService{
	return &Handler{
		Repository: repository,
		Validate: validate,
	}
}

func (h *Handler) Create(ctx context.Context,request web.RegisterUserPayload) web.UserResponse {
	err := h.Validate.Struct(request)

	if err != nil {
		panic(err)
	}

	user := domain.User{
		FirstName: request.FirstName,
		LastName: request.LastName,
		Email: request.Email,
		Password: request.Password,
	}

	user = h.Repository.Create(ctx,user)

	return (web.UserResponse{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
	})
}

