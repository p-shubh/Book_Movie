package controllers

import (
	"fmt"
	"log"
	"net/http"
	Db "restAPI/DB"
	models "restAPI/Models"

	"github.com/gin-gonic/gin"
)

func PostMoviesAvailability(c *gin.Context) {

	reqBody := models.Input{}

	if err := c.ShouldBind(&reqBody); err != nil {
		log.Println("*********FAILED TO BIND REQUEST BODY in PostMovies****** ")
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		c.Abort()
		return
	}

	fmt.Println("reqBody.Movie_name=", reqBody.Movie_name)

	id, err := Db.InsertMoviesInDb(reqBody.Movie_name)

	if err != nil {
		log.Println("*********FAILED in InsertMoviesInDb in PostMovies****** ")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
		c.Abort()
		return
	}

	result, err := Db.GetMoviesFromDb(id.String())
	if err != nil {
		log.Println("*********FAILED in GetMoviesFromDb in PostMovies****** ")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

func GetMoviesAvailability(c *gin.Context) {

	id, bools := c.GetQuery("id")

	if !bools {
		log.Println("*********FAILED TO BIND REQUEST BODY in GetMovies****** ")
		c.JSON(http.StatusBadRequest, gin.H{
			"err": bools,
		})
		c.Abort()
		return
	}

	result, err := Db.GetMoviesFromDb(id)
	if err != nil {
		log.Println("*********FAILED in GetMoviesFromDb in GetMovies****** ")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

func UpdatetMoviesAvailability(c *gin.Context) {

	reqBody := models.Input{}

	if err := c.ShouldBind(&reqBody); err != nil {
		log.Println("*********FAILED TO BIND REQUEST BODY in PutMovies****** ")
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		c.Abort()
		return
	}

	fmt.Println("reqBody.Movie_name=", reqBody.Movie_name, reqBody.MovieUUID)

	id, err := Db.UpdateMoviesFromDb(reqBody.Movie_name, reqBody.MovieUUID)

	if err != nil {
		log.Println("*********FAILED in UpdateMoviesFromDb in PutMovies****** ")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
		c.Abort()
		return
	}

	result, err := Db.GetMoviesFromDb(id)
	if err != nil {
		log.Println("*********FAILED in GetMoviesFromDb in PutMovies****** ")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

func DeleteMoviesAvailability(c *gin.Context) {

	reqBody := models.Input{}

	if err := c.ShouldBind(&reqBody); err != nil {
		log.Println("*********FAILED TO BIND REQUEST BODY in PutMovies****** ")
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		c.Abort()
		return
	}

	fmt.Println("reqBody.Movie_name=", reqBody.Movie_name, reqBody.MovieUUID)

	bools, err := Db.DeleteMoviesFromDb(reqBody.MovieUUID)

	if err != nil {
		log.Println("*********FAILED in UpdateMoviesFromDb in PutMovies****** ")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{
			"result": bools,
		})
	}
}
