package models

import "time"

// Client represents the Client table in the database
type Client struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ClientTypeId int       `gorm:"column:ClientTypeId" json:"client_type_id"`
	FirstName    string    `gorm:"column:FirstName" json:"first_name"`
	LastName     string    `gorm:"column:LastName" json:"last_name"`
	CardNumber   string    `gorm:"column:CardNumber" json:"card_number"`
	CardPassword string    `gorm:"column:CardPassword" json:"card_password"`
	NationalCode string    `gorm:"column:NationalCode" json:"national_code"`
	Code         string    `gorm:"column:Code" json:"code"`
	CreateDate   time.Time `gorm:"column:CreateDate" json:"create_date"`
}

// TableName sets the insert table name for this struct type
func (Client) TableName() string {
	return "Client" // Set the correct table name here
}
