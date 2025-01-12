package core

import (
	"context"
	"dating-app/internal/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindOneMatch: find match
func (r *Repository) FindOneMatch(ctx context.Context, fromProfileID, targetProfileID primitive.ObjectID, action string) (entities.Match, error) {
	data := entities.Match{}
	filter := bson.M{"from_profile_id": fromProfileID, "target_profile_id": targetProfileID, "action": action}
	err := r.db.Collection(r.collectionMatches).FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// FindMatches: find matches
func (r *Repository) FindMatchesIn24Hours(ctx context.Context, fromProfileID primitive.ObjectID) ([]entities.Match, error) {
	now := time.Now()
	yesteday := now.Add(-24 * time.Hour)
	filter := bson.M{
		"from_profile_id": fromProfileID,
		"created_at": bson.M{
			"$gte": yesteday,
			"$lte": now,
		},
	}

	cursor, err := r.db.Collection(r.collectionMatches).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	datas := []entities.Match{}
	if err := cursor.All(ctx, &datas); err != nil {
		return nil, err
	}

	return datas, nil
}

// InsertMatch: insert new match
func (r *Repository) InsertMatch(ctx context.Context, match entities.Match) error {
	now := time.Now()
	match.CreatedAt = now
	_, err := r.db.Collection(r.collectionMatches).InsertOne(ctx, match)
	if err != nil {
		return err
	}
	return nil
}
