package models

type CartBook struct {
	UserId   int
	BookId   int
	Quantity int
}

type CartAddBookReq struct {
	BookId   int `json:"bookId" validate:"required,number"`
	Quantity int `json:"quantity" validate:"required,number,gt=0"`
}
