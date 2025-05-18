package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/violasptntels/WorkWise_Backend/config"
	"github.com/violasptntels/WorkWise_Backend/models"
	"github.com/violasptntels/WorkWise_Backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getKaryawanCollection() *mongo.Collection {
	return config.DB.Collection("karyawan")
}

func GetAllKaryawan(c *fiber.Ctx) error {
	collection := getKaryawanCollection()
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data karyawan"})
	}
	var result []models.Karyawan
	if err := cursor.All(context.Background(), &result); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal parsing data"})
	}
	return c.JSON(result)
}

func GetKaryawanByID(c *fiber.Ctx) error {
	id := c.Params("id")
	collection := getKaryawanCollection()
	var karyawan models.Karyawan
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&karyawan)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Karyawan tidak ditemukan"})
	}
	return c.JSON(karyawan)
}

func CreateKaryawan(c *fiber.Ctx) error {
	collection := getKaryawanCollection()
	var karyawan models.Karyawan
	if err := c.BodyParser(&karyawan); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format data tidak valid"})
	}

	if !utils.IsValidEmail(karyawan.Email) {
		return c.Status(400).JSON(fiber.Map{"error": "Format email tidak valid"})
	}

	count, _ := collection.CountDocuments(context.Background(), bson.M{"_id": karyawan.ID})
	if count > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ID sudah digunakan"})
	}

	_, err := collection.InsertOne(context.Background(), karyawan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan data"})
	}

	return c.Status(201).JSON(karyawan)
}

func UpdateKaryawan(c *fiber.Ctx) error {
	collection := getKaryawanCollection()
	id := c.Params("id")
	var update models.Karyawan
	if err := c.BodyParser(&update); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"nama_lengkap":   update.NamaLengkap,
			"tanggal_lahir":  update.TanggalLahir,
			"jenis_kelamin":  update.JenisKelamin,
			"nomor_telepon":  update.NomorTelepon,
			"jabatan":        update.Jabatan,
			"posisi":         update.Posisi,
			"email":          update.Email,
		}},
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengubah data"})
	}

	return c.JSON(fiber.Map{"message": "Karyawan diperbarui"})
}

func DeleteKaryawan(c *fiber.Ctx) error {
	collection := getKaryawanCollection()
	id := c.Params("id")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus data"})
	}
	return c.JSON(fiber.Map{"message": "Karyawan berhasil dihapus"})
}
