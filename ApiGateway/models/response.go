package models

type Response struct {
	Error   bool        `json:"error"`
	Status  int64       `json:"status"`
	Massage string      `json:"massage"`
	Data    interface{} `json:"data"`
}
