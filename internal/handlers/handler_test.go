package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"dating-app/api/request"
	"dating-app/api/response"
	"dating-app/internal/auth"
	"dating-app/internal/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock services
type MockSignUpService struct {
	mock.Mock
}

func (m *MockSignUpService) Execute(ctx context.Context, input request.SignUp) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

type MockLogInService struct {
	mock.Mock
}

func (m *MockLogInService) Execute(ctx context.Context, input request.LogIn) (response.LogIn, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(response.LogIn), args.Error(1)
}

type MockProfileService struct {
	mock.Mock
}

func (m *MockProfileService) Execute(ctx context.Context, email string, input request.Profile) error {
	args := m.Called(ctx, email, input)
	return args.Error(0)
}

type MockSwipeService struct {
	mock.Mock
}

func (m *MockSwipeService) Execute(ctx context.Context, email string, input request.Swipe) (response.Swipe, error) {
	args := m.Called(ctx, email, input)
	return args.Get(0).(response.Swipe), args.Error(1)
}

type MockFeedService struct {
	mock.Mock
}

func (m *MockFeedService) Execute(ctx context.Context, email string) (response.Feed, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(response.Feed), args.Error(1)
}

// Test Ping
func TestPing(t *testing.T) {
	handler := &ConnectionHandler{}
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	rr := httptest.NewRecorder()

	handler.Ping(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	assert.Equal(t, "success", response["message"])
}

// Test SignUp
func TestSignUp_Success(t *testing.T) {
	mockSignUpService := new(MockSignUpService)
	handler := &ConnectionHandler{
		Ctx:           context.Background(),
		SignUpService: mockSignUpService,
	}

	input := request.SignUp{Email: "test@example.com", Password: "password"}
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	mockSignUpService.On("Execute", handler.Ctx, input).Return(nil)

	handler.SignUp(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var output response.SignUp
	json.NewDecoder(rr.Body).Decode(&output)
	assert.Equal(t, "success", output.Message)
}

func TestSignUp_InvalidRequest(t *testing.T) {
	handler := &ConnectionHandler{}
	body := []byte("invalid json")
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.SignUp(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestSignUp_UserAlreadyExists(t *testing.T) {
	mockSignUpService := new(MockSignUpService)
	handler := &ConnectionHandler{
		Ctx:           context.Background(),
		SignUpService: mockSignUpService,
	}

	input := request.SignUp{Email: "test@example.com", Password: "password"}
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	mockSignUpService.On("Execute", handler.Ctx, input).Return(entities.ErrSignUpUserAlreadyExists)

	handler.SignUp(rr, req)

	assert.Equal(t, http.StatusConflict, rr.Code)
	var output response.SignUp
	json.NewDecoder(rr.Body).Decode(&output)
	assert.Equal(t, "user with this email already exists.", output.Message)
}

// Test LogIn
func TestLogIn_Success(t *testing.T) {
	mockLogInService := new(MockLogInService)
	handler := &ConnectionHandler{
		Ctx:          context.Background(),
		LogInService: mockLogInService,
	}

	input := request.LogIn{Email: "test@example.com", Password: "password"}
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	mockLogInService.On("Execute", handler.Ctx, input).Return(response.LogIn{Token: "token"}, nil)

	handler.LogIn(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var output response.LogIn
	json.NewDecoder(rr.Body).Decode(&output)
	assert.Equal(t, "token", output.Token)
}

func TestLogIn_InvalidRequest(t *testing.T) {
	handler := &ConnectionHandler{}
	body := []byte("invalid json")
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.LogIn(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestLogIn_InvalidCredentials(t *testing.T) {
	mockLogInService := new(MockLogInService)
	handler := &ConnectionHandler{
		Ctx:          context.Background(),
		LogInService: mockLogInService,
	}

	input := request.LogIn{Email: "test@example.com", Password: "wrongpassword"}
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	mockLogInService.On("Execute", handler.Ctx, input).Return(response.LogIn{}, entities.ErrLoginInvalidCredentials)

	handler.LogIn(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	var output response.LogIn
	json.NewDecoder(rr.Body).Decode(&output)
	assert.Equal(t, "invalid email or password.", output.Message)
}

// Test Profile
func TestProfile_Success(t *testing.T) {
	mockProfileService := new(MockProfileService)
	handler := &ConnectionHandler{
		Ctx:            context.Background(),
		ProfileService: mockProfileService,
	}

	input := request.Profile{Description: "Hello World"}
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/profile", bytes.NewBuffer(body))
	req = req.WithContext(context.WithValue(req.Context(), auth.EmailKey, "test@example.com"))
	rr := httptest.NewRecorder()

	mockProfileService.On("Execute", handler.Ctx, "test@example.com", input).Return(nil)

	handler.Profile(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var output response.Profile
	json.NewDecoder(rr.Body).Decode(&output)
	assert.Equal(t, "success", output.Message)
}

func TestProfile_Unauthorized(t *testing.T) {
	handler := &ConnectionHandler{}
	body := []byte(`{"bio": "Hello World"}`)
	req, _ := http.NewRequest(http.MethodPost, "/profile", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.Profile(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	var output string
	json.NewDecoder(rr.Body).Decode(&output)
	assert.Equal(t, "Unauthorized: email not found in context", output)
}

// Test Swipe
func TestSwipe_Success(t *testing.T) {
	mockSwipeService := new(MockSwipeService)
	handler := &ConnectionHandler{
		Ctx:          context.Background(),
		SwipeService: mockSwipeService,
	}

	input := request.Swipe{TargetProfileID: "profile123"}
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/swipe", bytes.NewBuffer(body))
	req = req.WithContext(context.WithValue(req.Context(), auth.EmailKey, "test@example.com"))
	rr := httptest.NewRecorder()

	mockSwipeService.On("Execute", handler.Ctx, "test@example.com", input).Return(response.Swipe{}, nil)

	handler.Swipe(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var output response.Swipe
	json.NewDecoder(rr.Body).Decode(&output)
	assert.Equal(t, "success", output.Message)
}

func TestSwipe_Unauthorized(t *testing.T) {
	handler := &ConnectionHandler{}
	body := []byte(`{"profile_id": "profile123"}`)
	req, _ := http.NewRequest(http.MethodPost, "/swipe", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.Swipe(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	var output string
	json.NewDecoder(rr.Body).Decode(&output)
	assert.Equal(t, "Unauthorized: email not found in context", output)
}

// Test Feed
func TestFeed_Success(t *testing.T) {
	mockFeedService := new(MockFeedService)
	handler := &ConnectionHandler{
		Ctx:         context.Background(),
		FeedService: mockFeedService,
	}

	req, _ := http.NewRequest(http.MethodGet, "/feed", nil)
	req = req.WithContext(context.WithValue(req.Context(), auth.EmailKey, "test@example.com"))
	rr := httptest.NewRecorder()

	mockFeedService.On("Execute", handler.Ctx, "test@example.com").Return(response.Feed{}, nil)

	handler.Feed(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var output response.Feed
	json.NewDecoder(rr.Body).Decode(&output)
	assert.Equal(t, "success", output.Message)
}

func TestFeed_Unauthorized(t *testing.T) {
	handler := &ConnectionHandler{}
	req, _ := http.NewRequest(http.MethodGet, "/feed", nil)
	rr := httptest.NewRecorder()

	handler.Feed(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	var output string
	json.NewDecoder(rr.Body).Decode(&output)
	assert.Equal(t, "Unauthorized: email not found in context", output)
}
