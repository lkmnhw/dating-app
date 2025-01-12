package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Match struct {
	ID              *primitive.ObjectID `bson:"_id,omitempty"`
	FromProfileID   primitive.ObjectID  `bson:"from_profile_id"`
	TargetProfileID primitive.ObjectID  `bson:"target_profile_id"`
	Action          string              `bson:"action"`
	CreatedAt       time.Time           `bson:"created_at"`
}
