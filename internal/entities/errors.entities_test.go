package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignupError(t *testing.T) {
	var err SignupError = ErrSignUpUserAlreadyExists
	assert.Equal(t, "user with this email already exists.", err.Error())

	err = ErrSignUpMissingFields
	assert.Equal(t, "some required fields are missing.", err.Error())
}

func TestLoginError(t *testing.T) {
	var err LoginError = ErrLoginInvalidCredentials
	assert.Equal(t, "invalid email or password.", err.Error())

	err = ErrLoginUserNotFound
	assert.Equal(t, "user not found.", err.Error())

	err = ErrLoginMissingFields
	assert.Equal(t, "some required fields are missing.", err.Error())
}

func TestProfileError(t *testing.T) {
	var err ProfileError = ErrProfileMissingFields
	assert.Equal(t, "some required fields are missing.", err.Error())

	err = ErrProfileUserNotFound
	assert.Equal(t, "user not found.", err.Error())
}

func TestSwipeError(t *testing.T) {
	var err SwipeError = ErrSwipeMissingFields
	assert.Equal(t, "some required fields are missing.", err.Error())

	err = ErrSwipeProfileNotFound
	assert.Equal(t, "profile not found.", err.Error())

	err = ErrSwipeInvalidProfile
	assert.Equal(t, "invalid profile.", err.Error())
}

func TestFeedError(t *testing.T) {
	var err FeedError = ErrFeedProfileNotFound
	assert.Equal(t, "profile not found.", err.Error())
}
