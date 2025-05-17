package models

type Deadline struct {
	ID      string `json:"id" bson:"_id"`
	Tanggal string `json:"tanggal" bson:"tanggal"` // Format: yyyy-mm-dd
	

}
