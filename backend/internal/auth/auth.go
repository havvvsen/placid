package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterSignUpHandler(mux *http.ServeMux, logger *slog.Logger) {
	mux.HandleFunc("POST /api/register", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var req RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Register Request received"))
	})

	logger.Info("Registered /api/register")
}

func RegisterSignInHandler(mux *http.ServeMux, logger *slog.Logger) {
	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login Request received"))
	})
	logger.Info("Registered /api/login")
}
