package models

import "time"

// KYCRequest represents a KYC request in the database
type KYCRequest struct {
	ID                 int64     `gorm:"primaryKey;column:Id"`
	ClientID           int64     `gorm:"column:ClientId"`
	RequesterClientID  int64     `gorm:"column:RequesterClientId"`
	Code               string    `gorm:"column:Code"`
	MobileNumber       string    `gorm:"column:MobileNumber"`
	CreateDate         time.Time `gorm:"column:CreateDate"`
	KYCRequestStatusID int8      `gorm:"column:KYCRequestStatusId"`
}

// TableName sets the default table name for GORM
func (KYCRequest) TableName() string {
	return "KYCRequest"
}
