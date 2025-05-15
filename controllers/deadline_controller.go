package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/violasptntels/WorkWise_Backend/config"
	"github.com/violasptntels/WorkWise_Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getDeadlineCollection() *mongo.Collection {
	return config.DB.Collection("deadline")
}

func GetAllDeadline(c *fiber.Ctx) error {
	collection := getDeadlineCollection()
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data deadline"})
	}
	var deadlines []models.Deadline
	if err := cursor.All(context.Background(), &deadlines); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal parsing data deadline"})
	}
	return c.JSON(deadlines)
}

func GetDeadlineByID(c *fiber.Ctx) error {
	collection := getDeadlineCollection()
	id := c.Params("id")
	var deadline models.Deadline
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&deadline)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Deadline tidak ditemukan"})
	}
	return c.JSON(deadline)
}

func CreateDeadline(c *fiber.Ctx) error {
	collection := getDeadlineCollection()
	var deadline models.Deadline
	if err := c.BodyParser(&deadline); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	count, _ := collection.CountDocuments(context.Background(), bson.M{"_id": deadline.ID})
	if count > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ID deadline sudah digunakan"})
	}

	_, err := collection.InsertOne(context.Background(), deadline)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan data deadline"})
	}
	return c.Status(201).JSON(deadline)
}

func UpdateDeadline(c *fiber.Ctx) error {
	collection := getDeadlineCollection()
	id := c.Params("id")
	var update models.Deadline
	if err := c.BodyParser(&update); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}
	_, err := collection.UpdateOne(context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": update})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal memperbarui deadline"})
	}
	return c.JSON(fiber.Map{"message": "Deadline berhasil diperbarui"})
}

func DeleteDeadline(c *fiber.Ctx) error {
	collection := getDeadlineCollection()
	id := c.Params("id")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus deadline"})
	}
	return c.JSON(fiber.Map{"message": "Deadline berhasil dihapus"})
}
