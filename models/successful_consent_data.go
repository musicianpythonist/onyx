package models

import "time"

// ConsentToServiceProvider represents the structure of the ConsentToServiceProvider table
type ConsentToServiceProvider struct {
	ID                               int64     `json:"id" gorm:"primaryKey;column:Id"`
	ConsentToServiceProviderStatusId int8      `json:"consentToServiceProviderStatusId" gorm:"column:ConsentToServiceProviderStatusId"`
	ClientId                         int       `json:"clientId" gorm:"column:ClientId"`
	MerchantId                       int       `json:"merchantId" gorm:"column:MerchantId"`
	CreateDate                       time.Time `json:"createDate" gorm:"column:CreateDate"`
}

// TableName sets the default table name for GORM
func (ConsentToServiceProvider) TableName() string {
	return "ConsentToServiceProvider"
}
