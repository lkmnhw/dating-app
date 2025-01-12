package services

import (
	"context"
	"dating-app/api/request"
	"dating-app/api/response"
	"dating-app/internal/entities"
	"dating-app/internal/repositories/core"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SwipeInterface interface {
	Execute(ctx context.Context, email string, input request.Swipe) (response.Swipe, error)
}

type Swipe struct {
	coreRepo core.Interface
}

func NewSwipe(coreRepo core.Interface) *Swipe {
	return &Swipe{coreRepo: coreRepo}
}

func (s *Swipe) Execute(ctx context.Context, email string, input request.Swipe) (response.Swipe, error) {
	output := response.Swipe{Message: "success"}
	// validation
	err := s.validate(input)
	if err != nil {
		output.Message = err.Error()
		return output, err
	}

	// get user's profile
	user, err := s.coreRepo.FindOneUserByEmail(ctx, email)
	if err != nil || user.ID == nil {
		if err == mongo.ErrNoDocuments {
			err = entities.ErrSwipeProfileNotFound
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
			err = entities.ErrSwipeProfileNotFound
			output.Message = err.Error()
			return output, err
		}
		output.Message = err.Error()
		return output, err
	}

	targetProfileID, err := primitive.ObjectIDFromHex(input.TargetProfileID)
	if err != nil {
		err = entities.ErrSwipeInvalidProfile
		output.Message = err.Error()
		return output, err
	}

	// store match
	match := entities.Match{
		FromProfileID:   *profile.ID,
		TargetProfileID: targetProfileID,
		Action:          input.Action,
	}

	err = s.coreRepo.InsertMatch(ctx, match)
	if err != nil {
		output.Message = err.Error()
		return output, err
	}

	// check if match
	isMatch, err := s.coreRepo.FindOneMatch(ctx, targetProfileID, *profile.ID, "like")
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return output, err
		}
		output.Message = err.Error()
		return output, err
	}

	if isMatch.TargetProfileID == *profile.ID && isMatch.FromProfileID == targetProfileID {
		output.Match = true
	}

	return output, nil
}

func (s *Swipe) validate(input request.Swipe) error {
	// validate input
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return entities.ErrLoginMissingFields
	}

	return nil
}
