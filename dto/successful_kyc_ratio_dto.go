package dto

// SuccessfulKYCClientDTO represents the KYC ratio for a specific client
type SuccessfulKYCClientDTO struct {
	ClientId      int     `json:"clientId"`
	SuccessfulKYC float64 `json:"successfulKYC"`
}

// SuccessfulKYCRatioDayDTO represents the response for a single day
type SuccessfulKYCRatioDayDTO struct {
	Date       string                   `json:"date"`
	Statistics []SuccessfulKYCClientDTO `json:"statistics"`
}

// SuccessfulKYCRatioRangeDTO represents the response for a range (week or month)
type SuccessfulKYCRatioRangeDTO struct {
	Range      string                     `json:"range"`
	Statistics []SuccessfulKYCRatioDayDTO `json:"statistics"`
}
