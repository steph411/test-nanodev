package models

type StaffMember struct {
	ID     string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	AreaId string `json:"author" bson:"area"`
}
