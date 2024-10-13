package models

import "time"

// Client represents the Client table in the database
type Client struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ClientTypeId int       `json:"client_type_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	CardNumber   string    `json:"card_number"`
	CardPassword string    `json:"card_password"`
	NationalCode string    `json:"national_code"`
	Code         string    `json:"code"`
	CreateDate   time.Time `json:"create_date"`
}

// TableName sets the insert table name for this struct type
func (Client) TableName() string {
	return "Client" // Set the correct table name here
}
