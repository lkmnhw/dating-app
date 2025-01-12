package core

import (
	"context"
	"dating-app/app_config"
	"dating-app/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface: an abstraction of core repository
type Interface interface {
	// FindOneUserByEmail: return a user by email
	FindOneUserByEmail(ctx context.Context, email string) (entities.User, error)

	// InsertUser: insert new user
	InsertUser(ctx context.Context, user entities.User) error

	// FindOneProfileByUserID: return a profile by userID
	FindOneProfileByUserID(ctx context.Context, userID primitive.ObjectID) (entities.Profile, error)

	// FindProfilesByGenderAndAge: return list of profile by excluded ids, gender and age
	FindProfilesByGenderAndAge(ctx context.Context, excludProfileIDs []primitive.ObjectID, gender string, minAge, maxAge int, limit int64) ([]entities.Profile, error)

	// InsertProfile: insert new profile
	InsertProfile(ctx context.Context, profile entities.Profile) error

	// FindOneMatch: find match
	FindOneMatch(ctx context.Context, fromProfileID, targetProfileID primitive.ObjectID, action string) (entities.Match, error)

	// FindMatches: find matches
	FindMatchesIn24Hours(ctx context.Context, fromProfileID primitive.ObjectID) ([]entities.Match, error)

	// InsertMatch: insert new match
	InsertMatch(ctx context.Context, match entities.Match) error
}

// Repository: repository of core
type Repository struct {
	db                 *mongo.Database
	collectionUsers    string
	collectionProfiles string
	collectionMatches  string
}

// NewRepository: to initialize repository of user
func NewRepository(pdb *mongo.Database, isProd bool) *Repository {
	config := app_config.Get(isProd)
	return &Repository{
		db:                 pdb,
		collectionUsers:    config.ENV.COLLECTION_USERS,
		collectionProfiles: config.ENV.COLLECTION_PROFILES,
		collectionMatches:  config.ENV.COLLECTION_MATCHES,
	}
}
