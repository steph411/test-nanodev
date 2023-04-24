package router

import (
	"fiber-internal-api/common"
	"fiber-internal-api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddExpertiseAreaGroup(app *fiber.App) {
	ExpertiseAreaGroup := app.Group("/areas")

	ExpertiseAreaGroup.Get("/", getExpertiseAreas)
	ExpertiseAreaGroup.Get("/:id", getExpertiseArea)
	ExpertiseAreaGroup.Post("/", createExpertiseArea)
	ExpertiseAreaGroup.Put("/:id", updateExpertiseArea)
	ExpertiseAreaGroup.Delete("/:id", deleteExpertiseArea)
}

func getExpertiseAreas(c *fiber.Ctx) error {
	coll := common.GetDBCollection("expertise_areas")

	// find all ExpertiseAreas
	ExpertiseAreas := make([]models.ExpertiseArea, 0)
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		ExpertiseArea := models.ExpertiseArea{}
		err := cursor.Decode(&ExpertiseArea)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		ExpertiseAreas = append(ExpertiseAreas, ExpertiseArea)
	}

	return c.Status(200).JSON(fiber.Map{"data": ExpertiseAreas})
}

func getExpertiseArea(c *fiber.Ctx) error {
	coll := common.GetDBCollection("expertise_areas")

	// find the ExpertiseArea
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	ExpertiseArea := models.ExpertiseArea{}

	err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&ExpertiseArea)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"data": ExpertiseArea})
}

type createExpertiseAreaDTO struct {
	Name string `json:"name" bson:"name"`
}

func createExpertiseArea(c *fiber.Ctx) error {
	// validate the body
	b := new(createExpertiseAreaDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}
	// create the ExpertiseArea
	coll := common.GetDBCollection("expertise_areas")
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create ExpertiseArea",
			"message": err.Error(),
		})
	}

	// return the ExpertiseArea
	return c.Status(201).JSON(fiber.Map{
		"result": result,
	})
}

type updateExpertiseAreaDTO struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

func updateExpertiseArea(c *fiber.Ctx) error {
	// validate the body
	b := new(updateDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// get the id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	// update the ExpertiseArea
	coll := common.GetDBCollection("expertise_areas")
	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update ExpertiseArea",
			"message": err.Error(),
		})
	}

	// return the ExpertiseArea
	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

func deleteExpertiseArea(c *fiber.Ctx) error {
	// get the id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	// delete the ExpertiseArea
	coll := common.GetDBCollection("expertise_areas")
	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete ExpertiseArea",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}
