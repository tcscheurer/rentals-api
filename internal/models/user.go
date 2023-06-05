package models

type User struct {
	ID int32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}