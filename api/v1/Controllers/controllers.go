package controllers

import (
	"log"
	Db "restAPI/DB"
	models "restAPI/Models"

	"github.com/gin-gonic/gin"
)

func PostMovies(c *gin.Context) {

	reqBody := models.Input{}

	if err := c.ShouldBind(&reqBody); err != nil {
		log.Println("*********FAILED TO BIND REQUEST BODY in InsertMovies****** ")

	}

	id := Db.InsertMoviesInDb(reqBody.Movie_name)

	result, err := Db.GetMoviesFromDb(id)
	if err != nil {
		log.Println("*********FAILED in GetMoviesFromDb in InsertMovies****** ")
	}

	c.JSON(200, gin.H{
		"result": result,
	})

}
