package models

type CartBook struct {
	UserId        int     `json:"userId"`
	BookId        int     `json:"bookId"`
	Quantity      int     `json:"quantity"`
	PricePerUnite float64 `json:"pricePerUnite"`
}

type CartAddBookReq struct {
	BookId   int `json:"bookId" validate:"required,number"`
	Quantity int `json:"quantity" validate:"required,number,gt=0"`
}
