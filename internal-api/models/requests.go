package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
	InProgress Status = "IN_PROGRESS"
	Treated    Status = "TREATED"
	Rejected   Status = "REJECTED"
)

func (status Status) IsStatusValid() error {
	switch status {
	case InProgress, Treated, Rejected:
		return nil
	}
	return errors.New("Invalid status type")
}

type Request struct {
	ID              string             `json:"id" bson:"_id"`
	Title           string             `json:"title" bson:"title"`
	Content         string             `json:"content" bson:"content"`
	ExpertiseAreaId primitive.ObjectID `json:"expertiseAreaId" bson:"expertiseAreaId"`
	ExpertiseArea   ExpertiseArea      `json:"expertiseArea" bson:"expertiseArea"`
	UserId          string             `json:"userId" bson:"userId"`
	StaffMemberId   primitive.ObjectID `json:"staffMemberId" bson:"staffMemberId"`
	StaffMember     StaffMember        `json:"staffMember" bson:"staffMember,omitempty"`
	CreatedAt       primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt       primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	Status          Status             `json:"status" bson:"status"` // either IN_PROGRESS, TREATED or REJECTED
}
