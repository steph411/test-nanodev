package models

type Request struct {
	ID              string `json:"id" bson:"_id"`
	Title           string `json:"title" bson:"title"`
	Content         string `json:"content" bson:"content"`
	ExpertiseAreaId string `json:"expertiseAreaId" bson:"expertiseAreaId"`
	UserId          string `json:"userId" bson:"userId"`
	StaffMemberId   string `json:"staffMemberId" bson:"staffMemberId"`
}
