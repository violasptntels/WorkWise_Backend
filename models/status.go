package models

type Status struct {
	ID    string `json:"id" bson:"_id"`
	Nama  string `json:"nama" bson:"nama"`
	Color string `json:"color" bson:"color"`
}
