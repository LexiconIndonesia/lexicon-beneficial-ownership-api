package models

type BaseResponse struct {
	Data interface{}  `json:"data"`
	Meta MetaResponse `json:"meta"`
}
