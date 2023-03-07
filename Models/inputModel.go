package models

import (
	"time"

	"github.com/google/uuid"
)

type Input struct {
	MovieUUID uuid.UUID `json:"movieUUID"`
	MovieId           string    `json:"movieId"`
	Tickets_available int       `json:"ticketsAvailable"`
	Price             float64   `json:"Price"`
	Dates             time.Time `json:"Dates"`
	Id                uuid.UUID `json:"id"`
	Movie_name        string    `json:"movieName"`
	Name              string    `json:"name"`
	Tickets_unit      int       `json:"ticketsUnit"`
	Total_price       float64   `json:"ticketsPrice"`
	Date              time.Time `json:"date"`
}
