package models

type Karyawan struct {
	ID            string `json:"id" bson:"_id"`
	NamaLengkap   string `json:"nama_lengkap" bson:"nama_lengkap"`
	TanggalLahir  string `json:"tanggal_lahir" bson:"tanggal_lahir"`     // Format: yyyy-mm-dd
	JenisKelamin  string `json:"jenis_kelamin" bson:"jenis_kelamin"`
	NomorTelepon  string `json:"nomor_telepon" bson:"nomor_telepon"`
	Jabatan       string `json:"jabatan" bson:"jabatan"`
	Posisi        string `json:"posisi" bson:"posisi"`
	Email         string `json:"email" bson:"email"`
}
