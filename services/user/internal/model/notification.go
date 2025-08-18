package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notification struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId    int                `bson:"userId" json:"userId"`
	Type      string             `bson:"type" json:"type"`
	Title     string             `bson:"title" json:"title"`
	Message   string             `bson:"message" json:"message"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	ReadAt    *time.Time         `bson:"readAt,omitempty" json:"readAt,omitempty"`
}
