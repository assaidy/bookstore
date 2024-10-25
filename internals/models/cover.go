package models

type Cover struct {
	Id       int    `json:"id"`
	Encoding string `json:"-"`
	Content  string `json:"-"`
}

type CoverCreateOrUpdateReq struct {
	Encoding string `json:"encoding" validate:"required,imgEncoding"`
	Content  string `json:"content" validate:"required,base64,notBlank"`
}
