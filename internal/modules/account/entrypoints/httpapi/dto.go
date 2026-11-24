package httpapi

type transferRequest struct {
	FromID string `json:"fromId"`
	ToID   string `json:"toId"`
	Amount int64  `json:"amount"`
}
