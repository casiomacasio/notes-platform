package model

type NotificationInput struct {
    Type string      `json:"type" binding:"required"`
    Data interface{} `json:"data" binding:"required"`
}
