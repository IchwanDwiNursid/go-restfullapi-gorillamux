package app

import (
	"database/sql"
	"log"
)

func NewDB() (*sql.DB,error){
	// TODO : Change Configuration Mysql
	db , err := sql.Open("mysql","root:iwan@tcp(localhost:3306)/go_shop_api")

	if err != nil {
		log.Fatal(err)
	}

	return db,nil
}