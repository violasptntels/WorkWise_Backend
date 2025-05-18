package models

type Tugas struct {
	ID         string `json:"id" bson:"_id"`
	Judul      string `json:"judul" bson:"judul"`
	Deskripsi  string `json:"deskripsi" bson:"deskripsi"`
	KaryawanID string `json:"karyawan_id" bson:"karyawan_id"`
	Status     string `json:"status" bson:"status"`
	Deadline   string `json:"deadline" bson:"deadline"` // Format yyyy-mm-dd
}
