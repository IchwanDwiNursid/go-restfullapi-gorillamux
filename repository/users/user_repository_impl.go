package users

import (
	"context"
	"database/sql"
	"log"

	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/model/domain"
)

type UserRepositoryImpl struct {
	db *sql.DB
} 

func NewUserRepository(db *sql.DB) UserRepository{
	return &UserRepositoryImpl{db : db}
}

func(repository *UserRepositoryImpl) Create(ctx context.Context,user domain.User) domain.User {
	SQL := "INSERT INTO users (first_name,last_name,email,password) VALUES(?,?,?,?)"

	result , err := repository.db.ExecContext(ctx,SQL,user.FirstName,user.LastName,user.Email,user.Password)

	if err != nil{
		panic(err)
	}

	id , err := result.LastInsertId()
	if err != nil {
		panic(err)
	} 

	user.ID = int(id)
	log.Println(user)
	return user
}