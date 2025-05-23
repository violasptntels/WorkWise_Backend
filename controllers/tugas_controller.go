package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/violasptntels/WorkWise_Backend/config"
	"github.com/violasptntels/WorkWise_Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getTugasCollection() *mongo.Collection {
	return config.DB.Collection("tugas")
}

func GetAllTugas(c *fiber.Ctx) error {
	collection := getTugasCollection()
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data tugas"})
	}
	var tugas []models.Tugas
	if err := cursor.All(context.Background(), &tugas); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal parsing data"})
	}
	return c.JSON(tugas)
}

func GetTugasByID(c *fiber.Ctx) error {
	collection := getTugasCollection()
	id := c.Params("id")
	var tugas models.Tugas
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&tugas)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tugas tidak ditemukan"})
	}
	return c.JSON(tugas)
}

// func GetTugasByKaryawanID(c *fiber.Ctx) error {
//     karyawanID := c.Params("id")
//     collection := getTugasCollection()

//     // Ambil semua tugas yang memiliki karyawan_id sesuai
//     cursor, err := collection.Find(context.Background(), bson.M{"karyawan_id": karyawanID})
//     if err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil tugas untuk karyawan ini"})
//     }

//     var tugas []models.Tugas
//     if err := cursor.All(context.Background(), &tugas); err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Gagal parsing data tugas"})
//     }

//     return c.JSON(tugas)
// }

func CreateTugas(c *fiber.Ctx) error {
	tugasCollection := getTugasCollection()

	var tugas models.Tugas
	if err := c.BodyParser(&tugas); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	// Validasi ID tugas unik
	count, err := tugasCollection.CountDocuments(context.Background(), bson.M{"_id": tugas.ID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengecek ID tugas"})
	}
	if count > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ID tugas sudah digunakan"})
	}

	// Validasi dan format tanggal deadline
	parsedDate, err := time.Parse("2006-01-02", tugas.Deadline)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format tanggal tidak valid. Gunakan format yyyy-mm-dd"})
	}
	tugas.Deadline = parsedDate.Format("2006-01-02")

	// Simpan tugas ke database
	_, err = tugasCollection.InsertOne(context.Background(), tugas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan tugas"})
	}

	return c.JSON(fiber.Map{"message": "Tugas berhasil dibuat"})
}

func UpdateTugas(c *fiber.Ctx) error {
	collection := getTugasCollection()
	id := c.Params("id")

	var update models.Tugas
	if err := c.BodyParser(&update); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	update.ID = id // penting: pastikan ID di struct ikut diset

	filter := bson.M{"_id": id} // perbaikan disini
	updateDoc := bson.M{"$set": update}

	_, err := collection.UpdateOne(context.Background(), filter, updateDoc)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengupdate tugas"})
	}

	return c.JSON(fiber.Map{"message": "Tugas berhasil diperbarui"})
}

func DeleteTugas(c *fiber.Ctx) error {
	collection := getTugasCollection()
	id := c.Params("id")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus tugas"})
	}
	return c.JSON(fiber.Map{"message": "Tugas berhasil dihapus"})
}