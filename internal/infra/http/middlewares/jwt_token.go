package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/test_server/internal/app"
	"github.com/test_server/internal/domain"
)

type contextUserKey string

const contextUserIdKey contextUserKey = "1"
const BEARER_SCHEMA = "Bearer "

func AuthMiddleware(s app.TokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if user, err := authorizeWithAccessToken(r, s); err == nil {
				r = r.WithContext(context.WithValue(r.Context(), contextUserIdKey, user))
				next.ServeHTTP(w, r)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
		})
	}
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetAuthorizedUser(r)
		//if user == nil {
		//	log.Println("Warning! User authorization check is turned off!")
		//	next.ServeHTTP(w, r)
		//	return
		//} else
		if user.UserRole == domain.ROLE_ADMIN {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	})
}
func ModeratorOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetAuthorizedUser(r)
		if user == nil {
			log.Println("Warning! User authorization check is turned off!")
			next.ServeHTTP(w, r)
			return
		} else if user.UserRole == domain.ROLE_MODERATOR {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	})
}
func authorizeWithAccessToken(r *http.Request, s app.TokenService) (*domain.Token, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("failed to get token from Authorization header")
	}
	token := authHeader[len(BEARER_SCHEMA):]
	user, err := s.VerifyToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to get token from Authorization header")
	}
	return user, nil
}

func GetAuthorizedUser(r *http.Request) *domain.Token {
	object := r.Context().Value(contextUserIdKey)
	if user, ok := object.(*domain.Token); ok {
		return user
	}
	return nil
}
