package dto

type KYCRequestsDayResponseDTO struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

type KYCRequestsRangeResponseDTO struct {
	Range       string                      `json:"range"`
	KYCRequests []KYCRequestsDayResponseDTO `json:"kyc_requests"`
}
