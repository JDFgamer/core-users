package entities

import "time"

type User struct {
	ID                 uint               `json:"id"`
	Name               string             `json:"name"`
	LastName           string             `json:"last_name"`
	IdentifireDocument IdentifireDocument `json:"identifire_document"`
	Birthdate          time.Time          `json:"birthdate"`
	Gender             string             `json:"gender"`
}

type IdentifireDocument struct {
	Type   string `json:"type"`
	Number int    `json:"number"`
}
