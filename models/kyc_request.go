package models

import (
	"time"
)

// KYCRequest represents a KYC request in the system.
type KYCRequest struct {
	Id                 int       `json:"id" gorm:"primary_key"`
	ClientId           int       `json:"clientId"`
	RequesterClientId  int       `json:"requesterClientId"`
	Code               string    `json:"code"`
	MobileNumber       string    `json:"mobileNumber"`
	CreateDate         time.Time `json:"createDate"`
	KYCRequestStatusId int       `json:"kycRequestStatusId"`
}

// TableName sets the default table name for GORM
func (KYCRequest) TableName() string {
	return "KYCRequest"
}
