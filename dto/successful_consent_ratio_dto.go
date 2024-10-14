package dto

// SuccessfulConsentMerchantDTO represents the consent ratio for a specific merchant
type SuccessfulConsentMerchantDTO struct {
	MerchantId        int     `json:"merchantId"`
	SuccessfulConsent float64 `json:"successfulConsent"`
}

// SuccessfulConsentRatioDayDTO represents the response for a single day
type SuccessfulConsentRatioDayDTO struct {
	Date       string                         `json:"date"`
	Statistics []SuccessfulConsentMerchantDTO `json:"statistics"`
}

// SuccessfulConsentRatioRangeDTO represents the response for a range (week or month)
type SuccessfulConsentRatioRangeDTO struct {
	Range      string                         `json:"range"`
	Statistics []SuccessfulConsentRatioDayDTO `json:"statistics"`
}
