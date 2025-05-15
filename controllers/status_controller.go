package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/violasptntels/WorkWise-Backend/config"
	"github.com/violasptntels/WorkWise-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getStatusCollection() *mongo.Collection {
	return config.DB.Collection("status")
}

func GetAllStatus(c *fiber.Ctx) error {
	collection := getStatusCollection()
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data status"})
	}
	var status []models.Status
	if err := cursor.All(context.Background(), &status); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal parsing data status"})
	}
	return c.JSON(status)
}

func GetStatusByID(c *fiber.Ctx) error {
	collection := getStatusCollection()
	id := c.Params("id")
	var status models.Status
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&status)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Status tidak ditemukan"})
	}
	return c.JSON(status)
}

func CreateStatus(c *fiber.Ctx) error {
	collection := getStatusCollection()
	var status models.Status
	if err := c.BodyParser(&status); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	count, _ := collection.CountDocuments(context.Background(), bson.M{"_id": status.ID})
	if count > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ID status sudah digunakan"})
	}

	_, err := collection.InsertOne(context.Background(), status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan data status"})
	}
	return c.Status(201).JSON(status)
}

func UpdateStatus(c *fiber.Ctx) error {
	collection := getStatusCollection()
	id := c.Params("id")
	var update models.Status
	if err := c.BodyParser(&update); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}
	_, err := collection.UpdateOne(context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": update})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal memperbarui status"})
	}
	return c.JSON(fiber.Map{"message": "Status berhasil diperbarui"})
}

func DeleteStatus(c *fiber.Ctx) error {
	collection := getStatusCollection()
	id := c.Params("id")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus status"})
	}
	return c.JSON(fiber.Map{"message": "Status berhasil dihapus"})
}
