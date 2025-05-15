package models

type Karyawan struct {
    ID    string `json:"id" bson:"_id"`
    Nama  string `json:"nama" bson:"nama"`
    Email string `json:"email" bson:"email"`
}
