package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty"`
	Email     string              `bson:"email"`
	Password  string              `bson:"password"`
	CreatedAt time.Time           `bson:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at"`
}
