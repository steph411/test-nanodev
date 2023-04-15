package router

import (
	"fiber-internal-api/common"
	"fiber-internal-api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddRequestGroup(app *fiber.App) {
	requestGroup := app.Group("/requests")

	requestGroup.Get("/", getRequests)
	requestGroup.Get("/:id", getRequest)
	requestGroup.Post("/", createRequest)
	requestGroup.Put("/:id", updateRequest)
	requestGroup.Delete("/:id", deleteRequest)
}

func getRequests(c *fiber.Ctx) error {
	coll := common.GetDBCollection("requests")

	// find all requests
	requests := make([]models.Request, 0)
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		request := models.Request{}
		err := cursor.Decode(&request)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		requests = append(requests, request)
	}

	return c.Status(200).JSON(fiber.Map{"data": requests})
}

func getRequest(c *fiber.Ctx) error {
	coll := common.GetDBCollection("requests")

	// find the request
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

	request := models.Request{}

	err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&request)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"data": request})
}

type createDTO struct {
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Year   string `json:"year" bson:"year"`
}

func createRequest(c *fiber.Ctx) error {
	// validate the body
	b := new(createDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// create the request
	coll := common.GetDBCollection("requests")
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create request",
			"message": err.Error(),
		})
	}

	// return the request
	return c.Status(201).JSON(fiber.Map{
		"result": result,
	})
}

type updateDTO struct {
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
	Author string `json:"author,omitempty" bson:"author,omitempty"`
	Year   string `json:"year,omitempty" bson:"year,omitempty"`
}

func updateRequest(c *fiber.Ctx) error {
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

	// update the request
	coll := common.GetDBCollection("requests")
	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update request",
			"message": err.Error(),
		})
	}

	// return the request
	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

func deleteRequest(c *fiber.Ctx) error {
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

	// delete the request
	coll := common.GetDBCollection("requests")
	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete request",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}
