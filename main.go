package main

import (
	Db "restAPI/DB"
	routers "restAPI/api/Routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load env
	godotenv.Load(".env")

	Db.DBconnnection()

	routers.RoutesHandler(gin.Default())
}
