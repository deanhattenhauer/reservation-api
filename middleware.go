package main

import (
	"net/http"

	"github.com/deanhattenhauer/reservation-api/internal/auth"
)

func (cfg *apiConfig) middlewareVerifyAdmin(next http.Handler) http.Handler{	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request headers
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Unable to get token header", err)
			return
		}

		_, role, err := auth.ValidateJWT(token, cfg.jwtSecret)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't validate token", err)
			return
		}

		if role != "admin"{
			respondWithError(w, http.StatusForbidden, "Unauthorized access", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}