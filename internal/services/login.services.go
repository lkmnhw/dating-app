package services

import (
	"context"
	"dating-app/api/request"
	"dating-app/api/response"
	"dating-app/internal/auth"
	"dating-app/internal/entities"
	"dating-app/internal/repositories/core"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type LogInInterface interface {
	Execute(ctx context.Context, input request.LogIn) (response.LogIn, error)
}

type LogIn struct {
	coreRepo core.Interface
}

func NewLogIn(coreRepo core.Interface) *LogIn {
	return &LogIn{coreRepo: coreRepo}
}

func (s *LogIn) Execute(ctx context.Context, input request.LogIn) (response.LogIn, error) {
	output := response.LogIn{Message: "success"}
	// validation
	err := s.validate(input)
	if err != nil {
		output.Message = err.Error()
		return output, err
	}

	// validate password
	user, err := s.coreRepo.FindOneUserByEmail(ctx, input.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = entities.ErrLoginUserNotFound
			output.Message = err.Error()
			return output, err
		}
		output.Message = err.Error()
		return output, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		err = entities.ErrLoginInvalidCredentials
		output.Message = err.Error()
		return output, err
	}

	// generate token
	token, err := auth.GenerateToken(user.Email)
	if err != nil {
		output.Message = err.Error()
		return output, err
	}

	output.Token = token
	return output, nil
}

func (s *LogIn) validate(input request.LogIn) error {
	// validate input
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return entities.ErrLoginMissingFields
	}

	return nil
}
