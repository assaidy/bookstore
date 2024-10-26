package models

import "time"

type Order struct {
	Id         int          `json:"id"`
	UserId     int          `json:"userId"`
	AppliedAt  time.Time    `json:"appliedAt"`
	TotalPrice float64      `json:"totalPrice"`
	OrderBooks []*OrderBook `json:"orderBooks"`
}

type OrderBook struct {
	BookId        int     `json:"bookId"`
	Quantity      int     `json:"quantity"`
	PricePerUnite float64 `json:"pricePerUnite"`
}
