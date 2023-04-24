package router

import (
	"fiber-internal-api/common"
	"fiber-internal-api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddStaffMembersGroup(app *fiber.App) {
	staffMemberGroup := app.Group("/members")

	staffMemberGroup.Get("/", getStaffMembers)
	staffMemberGroup.Get("/:id", getStaffMember)
	staffMemberGroup.Post("/", createStaffMember)
	staffMemberGroup.Put("/:id", updateStaffMember)
	staffMemberGroup.Delete("/:id", deleteStaffMember)
}

func getStaffMembers(c *fiber.Ctx) error {
	coll := common.GetDBCollection("staff_members")

	// find all staffMembers
	staffMembers := make([]models.StaffMember, 0)
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		staffMember := models.StaffMember{}
		err := cursor.Decode(&staffMember)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		staffMembers = append(staffMembers, staffMember)
	}

	return c.Status(200).JSON(fiber.Map{"data": staffMembers})
}

func getStaffMember(c *fiber.Ctx) error {
	coll := common.GetDBCollection("staff_members")

	// find the staffMember
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

	staffMember := models.StaffMember{}

	err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&staffMember)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"data": staffMember})
}

type createStaffMemberDTO struct {
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	AreaId string `json:"areaId,omitempty" bson:"areaId,omitempty"`
}

func createStaffMember(c *fiber.Ctx) error {
	// validate the body
	b := new(createStaffMemberDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// create the staffMember
	coll := common.GetDBCollection("staff_members")
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create staffMember",
			"message": err.Error(),
		})
	}

	// return the staffMember
	return c.Status(201).JSON(fiber.Map{
		"result": result,
	})
}

type updateStaffMemberDTO struct {
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	AreaId string `json:"areaId,omitempty" bson:"areaId,omitempty"`
}

func updateStaffMember(c *fiber.Ctx) error {
	// validate the body
	b := new(updateStaffMemberDTO)
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

	// update the staffMember
	coll := common.GetDBCollection("staff_members")
	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update staffMember",
			"message": err.Error(),
		})
	}

	// return the staffMember
	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

func deleteStaffMember(c *fiber.Ctx) error {
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

	// delete the staffMember
	coll := common.GetDBCollection("staff_members")
	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete staffMember",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}
