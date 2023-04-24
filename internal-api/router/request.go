package router

import (
	"fiber-internal-api/common"
	"fiber-internal-api/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	userId := c.Query("userId")
	// matchStage := bson.D{{userId: userId}}
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "userId", Value: userId}}}}
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "expertise_areas"}, {Key: "localField", Value: "expertiseAreaId"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "expertiseArea"}}}}
	// unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$expertiseArea"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$expertiseArea"}}}}

	stages := mongo.Pipeline{lookupStage, unwindStage}
	if userId != "" {
		stages = mongo.Pipeline{matchStage, lookupStage, unwindStage}
	}
	cursor, err := coll.Aggregate(c.Context(), stages)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		request := models.Request{}
		// expertiseArea := models.ExpertiseArea{}
		// request.ExpertiseArea = expertiseArea
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

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: objectId}}}}
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "expertise_area"}, {Key: "localField", Value: "expertiseAreaId"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "expertiseArea"}}}}
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$expertiseArea"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}

	// err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&request)
	cursor, err := coll.Aggregate(c.Context(), mongo.Pipeline{matchStage, lookupStage, unwindStage})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = cursor.Decode(&request)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"data": request})
}

type createDTO struct {
	Title           string             `json:"title" bson:"title"`
	Content         string             `json:"content" bson:"content"`
	ExpertiseAreaId primitive.ObjectID `json:"expertiseAreaId" bson:"expertiseAreaId"`
	UserId          string             `json:"userId,omitempty" bson:"userId,omitempty"`
	CreatedAt       time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
	Status          models.Status      `json:"status" bson:"status"`
}

func createRequest(c *fiber.Ctx) error {
	// validate the body
	b := new(createDTO)
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	b.Status = models.InProgress
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	fmt.Println(b)
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
	Title           string             `json:"title" bson:"title"`
	Content         string             `json:"content" bson:"content"`
	ExpertiseAreaId primitive.ObjectID `json:"expertiseAreaId" bson:"expertiseAreaId"`
	UserId          string             `json:"userId" bson:"userId"`
	Status          models.Status      `json:"status" bson:"status"`
}

func updateRequest(c *fiber.Ctx) error {
	// validate the body
	b := new(updateDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	if statusValidateErr := b.Status.IsStatusValid(); statusValidateErr != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid status provided",
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
	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b, "$currentDate": bson.M{"updatedAt": true}})
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
