package main

import (
	"database/sql"
	"log"

	"github.com/gaba-bouliva/buyit/cmd/api"
	"github.com/gaba-bouliva/buyit/config"
	"github.com/gaba-bouliva/buyit/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: 								config.Envs.DBUser,
		Passwd: 							config.Envs.DBPassword,
		Addr: 								config.Envs.DBAddress,
		DBName: 							config.Envs.DBName,
		Net: 									"tcp",
		AllowNativePasswords: true,
		ParseTime: 						true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8081", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}


func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: successfully connected!")
}