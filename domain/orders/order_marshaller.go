package orders

type PublicOrder struct {
	Id        string `json:"id"`
	UserId    int64  `json:"user_id"`
	Status    string `json:"status"`
	ExpiresAt string `json:"expires_at"`
	ItemId    int64  `json:"item_id"`
}

// func (order *Order) Marshall() interface{} {
// 	orderJson, _ := json.Marshal(order)
// 	var order
// }
