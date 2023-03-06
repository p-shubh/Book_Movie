package Db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// var DB *

const (
	host     = "satao.db.elephantsql.com"
	port     = 5432
	user     = "jicuanqn"
	password = "wDUTP9P8Tem7AXm2ojr5q_CfPvfozXJJ"
	dbname   = "jicuanqn"
)

func DBconnnection() *sql.DB{

	/* CREATING THE CONNECTING STRING */
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// OPENING DB
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Println("**************DB CONNECTION FAILED*******************")
		log.Println(err.Error())
		// return err
	} else {
		log.Println("**************DB CONNECTED SUCCESFULLY*******************")

	}
	// defer db.Close()
	errPing := db.Ping()

	if errPing != nil {
		log.Println("**************DB PING STATUS*******************")
		log.Println(errPing.Error())
		// return errPing
	}

	_, err = db.Exec(CREATEDATABASE)

	if err != nil {
		log.Println("**************CREATING DATABASE FAILED*******************")
		log.Println(err.Error())
		// return err
	} else {
		log.Println("**************DATABASE CREATED SUCCESSFULLY*******************")

	}

	_, err = db.Exec(CREATETABLE)

	if err != nil {
		log.Println("**************CREATING TABLE FAILED*******************")
		log.Println(err.Error())
		// return err
	} else {
		log.Println("**************TABLE CREATED SUCCESSFULLY*******************")

	}
	// return nil

	return db
}
