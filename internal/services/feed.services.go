package services

import (
	"context"
	"dating-app/api/response"
	"dating-app/internal/entities"
	"dating-app/internal/repositories/core"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FeedInterface interface {
	Execute(ctx context.Context, email string) (response.Feed, error)
}

type Feed struct {
	coreRepo core.Interface
}

func NewFeed(coreRepo core.Interface) *Feed {
	return &Feed{coreRepo: coreRepo}
}

func (s *Feed) Execute(ctx context.Context, email string) (response.Feed, error) {
	output := response.Feed{Message: "success"}
	// get user's profile
	user, err := s.coreRepo.FindOneUserByEmail(ctx, email)
	if err != nil || user.ID == nil {
		if err == mongo.ErrNoDocuments {
			err = entities.ErrFeedProfileNotFound
			output.Message = err.Error()
			return output, err
		}
		output.Message = err.Error()
		return output, err
	}

	// find profile
	profile, err := s.coreRepo.FindOneProfileByUserID(ctx, *user.ID)
	if err != nil || profile.ID == nil {
		if err == mongo.ErrNoDocuments {
			err = entities.ErrFeedProfileNotFound
			output.Message = err.Error()
			return output, err
		}
		output.Message = err.Error()
		return output, err
	}

	// get matches in 24 hours
	matcheds, err := s.coreRepo.FindMatchesIn24Hours(ctx, *profile.ID)
	if err != nil && err != mongo.ErrNoDocuments {
		output.Message = err.Error()
		return output, err
	}

	excludProfileIDs := []primitive.ObjectID{}
	for _, matched := range matcheds {
		excludProfileIDs = append(excludProfileIDs, matched.TargetProfileID)
	}

	limit := 10
	now := time.Now()
	if profile.PremiumPackage == nil || now.Before(profile.PremiumPackage.PurchaseDate) || now.After(profile.PremiumPackage.ExpireDate) {
		limit = limit - len(matcheds)
	}

	// get profiles for feed
	feedProfiles, err := s.coreRepo.FindProfilesByGenderAndAge(ctx, excludProfileIDs, profile.Preference.Gender, profile.Preference.MinimumAge, profile.Preference.MaximumAge, int64(limit))
	if err != nil && err != mongo.ErrNoDocuments {
		output.Message = err.Error()
		return output, err
	}

	for _, feedProfile := range feedProfiles {
		age := now.Year() - feedProfile.DateOfBirth.Year()
		if now.YearDay() < feedProfile.DateOfBirth.YearDay() {
			age--
		}

		output.Data = append(output.Data, response.DataFeedProfile{
			Name:        feedProfile.Name,
			Description: feedProfile.Description,
			Gender:      feedProfile.Gender,
			Age:         age,
		})
	}
	return output, nil
}
