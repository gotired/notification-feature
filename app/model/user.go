package model

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User[T any] struct {
	ID        T         `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Tenant    T         `json:"tenant" bson:"tenant"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type CreateUser struct {
	Name   string    `json:"name"`
	Tenant uuid.UUID `json:"tenant"`
}

type UserRepository interface {
	Check(name string) (*User[bson.Binary], error)
	Insert(name string, tenant bson.Binary) error
	Detail(id bson.Binary) (*User[bson.Binary], error)
	List(limit, page int, keyword, order, orderKey string) ([]User[bson.Binary], error)
	Update(id bson.Binary, name string) error
	Delete(id bson.Binary) error
}

type UserService interface {
	Check(name string) (*User[uuid.UUID], error)
	Insert(name string, tenant uuid.UUID) error
	Detail(id uuid.UUID) (*User[uuid.UUID], error)
	List(limit, page int, keyword, order, orderKey string) ([]User[uuid.UUID], error)
	Update(id uuid.UUID, name string) error
	Delete(id uuid.UUID) error
}
