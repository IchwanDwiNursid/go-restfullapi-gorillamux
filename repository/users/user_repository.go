package users

import (
	"context"

	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/model/domain"
)


type UserRepository interface {
	Create(ctx context.Context,category domain.User) (domain.User)
}