package Db

import (
	"log"
	models "restAPI/Models"

	"github.com/google/uuid"
)

var db = DBconnnection()

func InsertMoviesInDb(movieName string) (uuid.UUID, error) {

	var id uuid.UUID

	sqlStatement := `INSERT INTO MOVIES(movie_name) VALUES($1) RETURNING id ;`

	if err := db.QueryRow(sqlStatement, movieName).Scan(&id); err != nil {
		log.Println("*********FAILED TO EXECUTE sqlStatement in InsertMoviesInDb****** ", err)
		return uuid.UUID{}, err
	}
	return id, nil
}

func GetMoviesFromDb(id string) (models.Movies, error) {

	movies := models.Movies{}

	sqlStatement := `SELECT id , movie_name FROM MOVIES WHERE id = $1 ;`

	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		log.Println("GetMoviesFromDb*********FAILED TO EXECUTE sqlStatement in GetMoviesFromDb****** ", err.Error())
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

func UpdateMoviesFromDb(movieName string, id uuid.UUID) (string, error) {

	sqlStatement := `UPDATE MOVIES SET movie_name =$1 WHERE id = $2 ;`

	if _, err := db.Exec(sqlStatement, movieName, id.String()); err != nil {
		log.Println("*********FAILED TO EXECUTE sqlStatement in UpdateMoviesFromDb****** ", err)
		return "", err

	}

	return id.String(), nil
}

func DeleteMoviesFromDb(id uuid.UUID) (bool, error) {

	sqlStatement := `DELETE FROM MOVIES WHERE id = $1 ;`

	if _, err := db.Exec(sqlStatement, id); err != nil {
		log.Println("*********FAILED TO EXECUTE sqlStatement in UpdateMoviesFromDb****** ", err)
		return false, err
	}
	return true, nil
}
