package model

import "time"

type Notification struct {
	Id        string     `bson:"_id"`
	UserId    int        `bson:"userId" json:"userId"`
	Type      string     `bson:"type" json:"type"`
	Title     string     `bson:"title" json:"title"`
	Message   string     `bson:"message" json:"message"`
	Status    string     `bson:"status" json:"status"`
	CreatedAt time.Time  `bson:"createdAt" json:"createdAt"`
	ReadAt    *time.Time `bson:"readAt,omitempty" json:"readAt,omitempty"`
}
