package Db

import (
	"log"
	models "restAPI/Models"

	"github.com/google/uuid"
)

var db = DBconnnection()

func InsertMoviesInDb(movieName string) uuid.UUID {

	var id uuid.UUID

	/* Id         uuid.UUID `json:"id"`
	Movie_name string    `json:"movieName"` */

	sqlStatement := `INSERT INTO MOVIES(movie_name) VALUES($1) RETURNING id ;`

	if err := db.QueryRow(sqlStatement, movieName).Scan(&id); err != nil {
		log.Println("*********FAILED TO EXECUTE sqlStatement in InsertMoviesInDb****** ")
	}

	return id
}

func GetMoviesFromDb(id uuid.UUID) (models.Movies, error) {

	/* Id         uuid.UUID `json:"id"`
	Movie_name string    `json:"movieName"` */

	movies := models.Movies{}

	sqlStatement := `SELECT FROM MOVIES WHERE id = $1 ;`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Println("*********FAILED TO EXECUTE sqlStatement in GetMoviesFromDb****** ", err.Error())
		return movies, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&movies.Id, &movies.Movie_name); err != nil {
			log.Println("*********FAILED TO SCAN ROWS in GetMoviesFromDb****** ", err.Error())
			return movies, err
		}
	}
	return movies, nil
}
