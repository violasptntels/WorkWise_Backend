package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/violasptntels/WorkWise-Backend/config"
	"github.com/violasptntels/WorkWise-Backend/models"
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

func CreateTugas(c *fiber.Ctx) error {
	collection := getTugasCollection()
	var tugas models.Tugas
	if err := c.BodyParser(&tugas); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	count, _ := collection.CountDocuments(context.Background(), bson.M{"_id": tugas.ID})
	if count > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ID tugas sudah digunakan"})
	}

	_, err := collection.InsertOne(context.Background(), tugas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan data tugas"})
	}
	return c.Status(201).JSON(tugas)
}

func UpdateTugas(c *fiber.Ctx) error {
	collection := getTugasCollection()
	id := c.Params("id")
	var update models.Tugas
	if err := c.BodyParser(&update); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	_, err := collection.UpdateOne(context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": update})
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
