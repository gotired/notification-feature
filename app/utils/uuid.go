package utils

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ConvertUUIDToBinary(id uuid.UUID) bson.Binary {
	return bson.Binary{
		Subtype: 4,
		Data:    id[:],
	}
}

func ConvertBinarytoUUID(id bson.Binary) (uuid.UUID, error) {
	return uuid.FromBytes(id.Data)
}
