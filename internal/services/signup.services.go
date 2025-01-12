package services

import (
	"context"
	"dating-app/api/request"
	"dating-app/internal/entities"
	"dating-app/internal/repositories/core"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type SignUpInterface interface {
	Execute(ctx context.Context, input request.SignUp) error
}

type SignUp struct {
	coreRepo core.Interface
}

func NewSignUp(coreRepo core.Interface) *SignUp {
	return &SignUp{coreRepo: coreRepo}
}

func (s *SignUp) Execute(ctx context.Context, input request.SignUp) error {
	// validation
	err := s.validate(ctx, input)
	if err != nil {
		return err
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// insert user
	user := entities.User{Email: input.Email, Password: string(hashedPassword)}
	err = s.coreRepo.InsertUser(ctx, user)
	if mongo.IsDuplicateKeyError(err) {
		return entities.ErrSignUpUserAlreadyExists
	}

	return err
}

func (s *SignUp) validate(ctx context.Context, input request.SignUp) error {
	// validate input
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return entities.ErrSignUpMissingFields
	}

	// validate new user
	_, err = s.coreRepo.FindOneUserByEmail(ctx, input.Email)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil
	}
	return err
}
