package core

import (
	"context"
	"dating-app/internal/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOneProfileByUserID: return a profile by userID
func (r *Repository) FindOneProfileByUserID(ctx context.Context, userID primitive.ObjectID) (entities.Profile, error) {
	data := entities.Profile{}
	filter := bson.M{"user_id": userID}
	err := r.db.Collection(r.collectionProfiles).FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// FindProfilesByGenderAndAge: return list of profile by gender and age
func (r *Repository) FindProfilesByGenderAndAge(ctx context.Context, excludProfileIDs []primitive.ObjectID, gender string, minAge, maxAge int, limit int64) ([]entities.Profile, error) {
	now := time.Now()
	youngestDate := now.AddDate(-minAge, 0, 0)
	oldestDate := now.AddDate(-maxAge, 0, 0)
	filter := bson.M{
		"gender": gender,
		"date_of_birth": bson.M{
			"$gte": oldestDate,
			"$lte": youngestDate,
		},
	}
	findOptions := options.Find()
	findOptions.SetLimit(limit)

	cursor, err := r.db.Collection(r.collectionProfiles).Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	datas := []entities.Profile{}
	if err := cursor.All(ctx, &datas); err != nil {
		return nil, err
	}

	return datas, nil
}

// InsertProfile: insert new profile
func (r *Repository) InsertProfile(ctx context.Context, profile entities.Profile) error {
	filter := bson.M{"user_id": profile.UserID}
	now := time.Now()
	profile.UpdatedAt = &now
	profile.CreatedAt = nil
	update := bson.M{
		"$set": profile,
		"$setOnInsert": bson.M{
			"created_at": now,
		},
	}
	options := options.FindOneAndUpdate().SetUpsert(true)
	err := r.db.Collection(r.collectionProfiles).FindOneAndUpdate(ctx, filter, update, options).Err()
	if err != nil {
		return err
	}
	return nil
}
