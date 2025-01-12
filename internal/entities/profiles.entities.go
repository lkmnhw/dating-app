package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID             *primitive.ObjectID `bson:"_id,omitempty"`
	UserID         *primitive.ObjectID `bson:"user_id"`
	Name           string              `bson:"name"`
	Description    string              `bson:"description"`
	Gender         string              `bson:"gender"`
	DateOfBirth    time.Time           `bson:"date_of_birth"`
	Preference     ProfilePreference   `bson:"preference"`
	PremiumPackage *PremiumPackage     `bson:"premium_package,omitempty"`
	CreatedAt      *time.Time          `bson:"created_at,omitempty"`
	UpdatedAt      *time.Time          `bson:"updated_at"`
}

type ProfilePreference struct {
	Gender     string `bson:"gender"`
	MinimumAge int    `bson:"minimum_age"`
	MaximumAge int    `bson:"maximum_age"`
}

type PremiumPackage struct {
	PurchaseDate time.Time `bson:"purchase_date"`
	ExpireDate   time.Time `bson:"expire_date"`
}
