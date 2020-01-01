package models

type OrderModel struct {
	Cards       []StoreCardModel `json:"cards"`
	FinalAmount uint             `json:"final_amount"`
}
