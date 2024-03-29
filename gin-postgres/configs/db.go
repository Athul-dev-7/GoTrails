package configs

import (
	"log"
	"os"

	"gin-postgres/controllers"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "db_username",
		Password: "db_password",
		Addr:     "localhost:5432",
		Database: "db_dbname",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to DB")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db

}
