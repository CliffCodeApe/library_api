package dto

type LendingRequest struct {
	BookURL string `json:"book_url"`
	UserID  uint64 `json:"user_id"`
}

type LendingResponse struct {
	StatusCode int           `json:"status_code"`
	Message    string        `json:"message"`
	Data       []LendingData `json:"data"`
}

type MakeLendingResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       LendingData `json:"data"`
}

type LendingData struct {
	ID     uint64 `json:"id"`
	BookID uint64 `json:"book_id"`
	UserID uint64 `json:"user_id"`
	Status string `json:"status"`
}
