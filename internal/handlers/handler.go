package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"dating-app/api/request"
	"dating-app/api/response"
	"dating-app/internal/auth"
	"dating-app/internal/entities"
	"dating-app/internal/helpers"
	"dating-app/internal/services"
)

type ConnectionHandler struct {
	Ctx            context.Context
	SignUpService  services.SignUpInterface
	LogInService   services.LogInInterface
	ProfileService services.ProfileInterface
	SwipeService   services.SwipeInterface
	FeedService    services.FeedInterface
}

// Ping: for health check
func (c *ConnectionHandler) Ping(w http.ResponseWriter, r *http.Request) {
	stop := benchmark("Ping")
	defer stop()
	helpers.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "success"})
}

// SignUp: to create new user
func (c *ConnectionHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	stop := benchmark("SignUp")
	defer stop()

	var input request.SignUp
	var output response.SignUp

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	err = c.SignUpService.Execute(c.Ctx, input)
	if err != nil {
		output.Message = err.Error()
		if _, ok := err.(entities.SignupError); ok {
			helpers.WriteJSONResponse(w, http.StatusConflict, output)
			return
		}
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, output)
		return
	}

	output.Message = "success"
	helpers.WriteJSONResponse(w, http.StatusOK, output)
}

// LogIn: to login and create token
func (c *ConnectionHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	stop := benchmark("LogIn")
	defer stop()

	var input request.LogIn
	var output response.LogIn

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	output, err = c.LogInService.Execute(c.Ctx, input)
	if err != nil {
		output.Message = err.Error()
		if _, ok := err.(entities.LoginError); ok {
			helpers.WriteJSONResponse(w, http.StatusUnauthorized, output)
			return
		}
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, output)
		return
	}

	helpers.WriteJSONResponse(w, http.StatusOK, output)
}

func (c *ConnectionHandler) Profile(w http.ResponseWriter, r *http.Request) {
	stop := benchmark("Profile")
	defer stop()

	var input request.Profile
	var output response.Profile

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	email, ok := r.Context().Value(auth.EmailKey).(string)
	if !ok {
		helpers.WriteJSONResponse(w, http.StatusUnauthorized, "Unauthorized: email not found in context")
		return
	}

	err = c.ProfileService.Execute(c.Ctx, email, input)
	if err != nil {
		output.Message = err.Error()
		if _, ok := err.(entities.ProfileError); ok {
			helpers.WriteJSONResponse(w, http.StatusBadRequest, output)
			return
		}
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, output)
		return
	}

	output.Message = "success"
	helpers.WriteJSONResponse(w, http.StatusOK, output)
}

func (c *ConnectionHandler) Swipe(w http.ResponseWriter, r *http.Request) {
	stop := benchmark("Swipe")
	defer stop()

	var input request.Swipe
	var output response.Swipe

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	email, ok := r.Context().Value(auth.EmailKey).(string)
	if !ok {
		helpers.WriteJSONResponse(w, http.StatusUnauthorized, "Unauthorized: email not found in context")
		return
	}

	output, err = c.SwipeService.Execute(c.Ctx, email, input)
	if err != nil {
		output.Message = err.Error()
		if _, ok := err.(entities.SwipeError); ok {
			helpers.WriteJSONResponse(w, http.StatusBadRequest, output)
			return
		}
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, output)
		return
	}

	output.Message = "success"
	helpers.WriteJSONResponse(w, http.StatusOK, output)
}

func (c *ConnectionHandler) Feed(w http.ResponseWriter, r *http.Request) {
	stop := benchmark("Feed")
	defer stop()

	var output response.Feed

	email, ok := r.Context().Value(auth.EmailKey).(string)
	if !ok {
		helpers.WriteJSONResponse(w, http.StatusUnauthorized, "Unauthorized: email not found in context")
		return
	}

	output, err := c.FeedService.Execute(c.Ctx, email)
	if err != nil {
		output.Message = err.Error()
		if _, ok := err.(entities.SwipeError); ok {
			helpers.WriteJSONResponse(w, http.StatusBadRequest, output)
			return
		}
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, output)
		return
	}

	output.Message = "success"
	helpers.WriteJSONResponse(w, http.StatusOK, output)
}

func benchmark(endpoint string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		log.Printf("%s: %s\n", endpoint, elapsed)
	}
}
