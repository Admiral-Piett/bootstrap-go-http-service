package models

var Settings struct {
    Port string
}

var SETTINGS = Settings

type BaseResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}
