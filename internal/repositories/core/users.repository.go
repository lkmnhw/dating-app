package core

import (
	"context"
	"dating-app/internal/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// FindOneUserByEmail: return a user by email
func (r *Repository) FindOneUserByEmail(ctx context.Context, email string) (entities.User, error) {
	data := entities.User{}
	filter := bson.M{"email": email}
	err := r.db.Collection(r.collectionUsers).FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// InsertUser: insert new user
func (r *Repository) InsertUser(ctx context.Context, user entities.User) error {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err := r.db.Collection(r.collectionUsers).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
