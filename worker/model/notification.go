package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Notification[T any] struct {
	ID        T         `json:"id" bson:"_id,omitempty"`
	TenantID  T         `json:"tenant_id" bson:"tenant_id"`
	Message   string    `json:"message" bson:"message"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UserID    T         `json:"user_id" bson:"user_id"`
}

type NotificationRepo interface {
	StoreNotifications([]Notification[bson.Binary]) error
	GetUserIDsByTenantID(id bson.Binary) ([]bson.Binary, error)
}

type NotificationPayload struct {
	TenantID string `json:"tenant_id"`
	Message  string `json:"message"`
}
