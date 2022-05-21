package template

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name     string             `json:"name" bson:"name"`
	Password string             `json:"password" bson:"password"`
	ID       primitive.ObjectID `bson:"_id,omitempty"`
}
