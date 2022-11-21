package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var dbConnection *sql.DB

const ( //estos datos también pueden ser llevados a una variable de entorno
	host     = "localhost" //este va a ser el mismo siempre que sea en local
	port     = 8000        //el puerto
	user     = "alfred"    //mi user
	password = "4lfr3d"    //mi password
	dbname   = "db_1"      //mi db que cree
)

func InitDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	dbConnection, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer dbConnection.Close()

	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Println(dbConnection)
	fmt.Println("Successfully connected to db!")
}
