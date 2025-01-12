package services

import (
	"context"
	"dating-app/api/request"
	"dating-app/internal/entities"
	"dating-app/internal/repositories/core"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileInterface interface {
	Execute(ctx context.Context, email string, input request.Profile) error
}

type Profile struct {
	coreRepo core.Interface
}

func NewProfile(coreRepo core.Interface) *Profile {
	return &Profile{coreRepo: coreRepo}
}

func (s *Profile) Execute(ctx context.Context, email string, input request.Profile) error {
	// validation
	err := s.validate(input)
	if err != nil {
		return err
	}

	// check if user exists
	user, err := s.coreRepo.FindOneUserByEmail(ctx, email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entities.ErrProfileUserNotFound
		}
		return err
	}

	// parse from string into date
	layout := "2006-01-02"
	dateOfBirth, err := time.Parse(layout, input.DateOfBirth)
	if err != nil {
		return entities.ErrProfileMissingFields
	}

	profile := entities.Profile{
		UserID:      user.ID,
		Name:        input.Name,
		Description: input.Description,
		Gender:      input.Gender,
		DateOfBirth: dateOfBirth,
		Preference: entities.ProfilePreference{
			Gender:     input.Preference.Gender,
			MinimumAge: input.Preference.MinimumAge,
			MaximumAge: input.Preference.MaximumAge,
		},
		PremiumPackage: nil,
	}

	if input.PremiumPackage.PurchaseDate != "" && input.PremiumPackage.ExpireDate != "" {
		profile.PremiumPackage = &entities.PremiumPackage{}
		profile.PremiumPackage.PurchaseDate, err = time.Parse(layout, input.PremiumPackage.PurchaseDate)
		if err != nil {
			return entities.ErrProfileMissingFields
		}

		profile.PremiumPackage.ExpireDate, err = time.Parse(layout, input.PremiumPackage.ExpireDate)
		if err != nil {
			return entities.ErrProfileMissingFields
		}
	}

	err = s.coreRepo.InsertProfile(ctx, profile)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	return nil
}

func (s *Profile) validate(input request.Profile) error {
	// validate input
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return entities.ErrProfileMissingFields
	}

	return nil
}
