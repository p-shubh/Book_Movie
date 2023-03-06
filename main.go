package main

import (
	Db "restAPI/DB"

	"github.com/joho/godotenv"
)

func main() {

	// load env
	godotenv.Load(".env")

	Db.DBconnnection()

}
