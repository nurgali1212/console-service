package model

type Object struct {
	Date    interface{} `json:"date"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}
