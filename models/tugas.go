package models

type Tugas struct {
	ID         string `json:"id" bson:"_id"`
	Nama       string `json:"nama" bson:"nama"`
	Deskripsi  string `json:"deskripsi" bson:"deskripsi"`
	KaryawanID string `json:"karyawan_id" bson:"karyawan_id"`
	StatusID   string `json:"status_id" bson:"status_id"`
	DeadlineID string `json:"deadline_id" bson:"deadline_id"`
}
