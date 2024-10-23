package models

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryCreateOrUpdateReq struct {
	Name string `json:"name" validate:"required,min=3,max=32,notBlank"`
}
