package models

import "time"

// KYCRequest represents the KYCRequest table
type KYCRequest struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	ClientId           uint      `json:"client_id"`
	RequesterClientId  uint      `json:"requester_client_id"`
	Code               string    `json:"code"`
	MobileNumber       string    `json:"mobile_number"`
	CreateDate         time.Time `json:"create_date"`
	KYCRequestStatusId uint8     `json:"kyc_request_status_id"`
}

// KYCRequestStatus represents the KYCRequestStatus table
type KYCRequestStatus struct {
	ID    uint8  `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

// TableName sets the insert table name for this struct type
func (KYCRequest) TableName() string {
	return "KYCRequest" // Set the correct table name here
}
