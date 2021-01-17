package orders

type Order struct {
	ID        string `json:"id"`
	UserId    int64  `json:"user_id"`
	Status    string `json:"status"`
	ExpiresAt string `json:"expires_at"`
	ItemId    string `json:"item_id"`
}
