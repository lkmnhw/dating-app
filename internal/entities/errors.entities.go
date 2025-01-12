package entities

// Signup-related errors
type SignupError string

func (e SignupError) Error() string {
	return string(e)
}

const (
	ErrSignUpUserAlreadyExists SignupError = "user with this email already exists."
	ErrSignUpMissingFields     SignupError = "some required fields are missing."
)

// Login-related errors
type LoginError string

func (e LoginError) Error() string {
	return string(e)
}

const (
	ErrLoginInvalidCredentials LoginError = "invalid email or password."
	ErrLoginUserNotFound       LoginError = "user not found."
	ErrLoginMissingFields      LoginError = "some required fields are missing."
)

// Profile-related errors
type ProfileError string

func (e ProfileError) Error() string {
	return string(e)
}

const (
	ErrProfileMissingFields ProfileError = "some required fields are missing."
	ErrProfileUserNotFound  ProfileError = "user not found."
)

// Swipe-related errors
type SwipeError string

func (e SwipeError) Error() string {
	return string(e)
}

const (
	ErrSwipeMissingFields   SwipeError = "some required fields are missing."
	ErrSwipeProfileNotFound SwipeError = "profile not found."
	ErrSwipeInvalidProfile  SwipeError = "invalid profile."
)

// Feed-related errors
type FeedError string

func (e FeedError) Error() string {
	return string(e)
}

const (
	ErrFeedProfileNotFound FeedError = "profile not found."
)
