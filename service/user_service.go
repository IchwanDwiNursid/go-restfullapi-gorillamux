package service

import (
	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/model/web"
	"golang.org/x/net/context"
)


type UserService interface {
	Create(ctx context.Context , request web.RegisterUserPayload) web.UserResponse
}