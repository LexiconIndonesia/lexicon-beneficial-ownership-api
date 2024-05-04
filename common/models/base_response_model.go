package models

type BasePaginationResponse struct {
	Data interface{}  `json:"data"`
	Meta MetaResponse `json:"meta"`
}

type BaseResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Msg   string `json:"message"`
}
