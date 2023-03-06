package models

import (
	"time"

	"github.com/google/uuid"
)

type MOVIES_AVAILABILITY struct {
	MovieId           string    `json:"movieId"`
	Tickets_available int       `json:"ticketsAvailable"`
	Price             float64   `json:"Price"`
	Dates             time.Time `json:"Dates"`
}

type Movies struct {
	Id         uuid.UUID `json:"id"`
	Movie_name string    `json:"movieName"`
}

type BOOKING struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Tickets_unit int       `json:"ticketsUnit"`
	Total_price  float64   `json:"ticketsPrice"`
	Date         time.Time `json:"date"`
}
