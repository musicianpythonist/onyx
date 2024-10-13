package dto

// NewClientsDayResponseDTO represents the response for new clients for a single day
type NewClientsDayResponseDTO struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

// NewClientsRangeResponseDTO represents the response for new clients for a range (week, month, etc.)
type NewClientsRangeResponseDTO struct {
	Range      string                     `json:"range"`
	NewClients []NewClientsDayResponseDTO `json:"newClients"`
}
