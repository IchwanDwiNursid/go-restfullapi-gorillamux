package main

import (
	"database/sql"
	"log"

	"github.com/IchwanDwiNursid/go-restfullapi-gorillamux/app"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db,err:= app.NewDB()

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)


	server := app.NewApiServer(":8080",db)
	if err := server.Run() ; err != nil {
		log.Fatal(err)
	}

}


func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	log.Println("DB : Successfully connected!")
}