package servers

import (
	"dating-app/internal/auth"
	"dating-app/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// api routes
func router(hndlr *handlers.ConnectionHandler) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/ping", http.HandlerFunc(hndlr.Ping)).Methods(http.MethodGet)
	r.Handle("/signup", http.HandlerFunc(hndlr.SignUp)).Methods(http.MethodPost)
	r.Handle("/login", http.HandlerFunc(hndlr.LogIn)).Methods(http.MethodPost)
	r.Handle("/profile", auth.Middleware(http.HandlerFunc(hndlr.Profile))).Methods(http.MethodPost)
	r.Handle("/swipe", auth.Middleware(http.HandlerFunc(hndlr.Swipe))).Methods(http.MethodPost)
	r.Handle("/feed", auth.Middleware(http.HandlerFunc(hndlr.Feed))).Methods(http.MethodGet)
	return r
}
