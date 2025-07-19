package model

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Tenant[T any] struct {
	ID        T         `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type TenantRepository interface {
	Check(name string) (*Tenant[bson.Binary], error)
	CheckByID(id bson.Binary) (*Tenant[bson.Binary], error)
	Insert(name string) error
	Detail(id bson.Binary) (*Tenant[bson.Binary], error)
	List(limit, page int, keyword, order, orderKey string) ([]Tenant[bson.Binary], error)
	Update(id bson.Binary, name string) error
	Delete(id bson.Binary) error
}

type TenantService interface {
	Check(name string) (*Tenant[uuid.UUID], error)
	CheckByID(id uuid.UUID) (*Tenant[uuid.UUID], error)
	Insert(name string) error
	Detail(id uuid.UUID) (*Tenant[uuid.UUID], error)
	List(limit, page int, keyword, order, orderKey string) ([]Tenant[uuid.UUID], error)
	Update(id uuid.UUID, name string) error
	Delete(id uuid.UUID) error
}
