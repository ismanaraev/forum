package models

type CategoryID int

type Category struct {
	ID   CategoryID `json:"id"`
	Name string     `json:"name"`
}
