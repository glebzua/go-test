package app

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/test_server/internal/domain"
)

type TokenService interface {
	CreateToken(user *domain.User) (string, error)
	VerifyToken(tokenString string) (*domain.Token, error)
}

type sessionService struct {
	secretAccess []byte
}

type userAccessClaims struct {
	UserId   int64       `json:"user_id"`
	UserRole domain.Role `json:"user_role"`
	jwt.StandardClaims
}

func NewTokenService(secretAccess []byte) TokenService {
	return &sessionService{
		secretAccess: secretAccess,
	}
}

func (s *sessionService) VerifyToken(tokenString string) (*domain.Token, error) {
	claims, err := parseJWT(tokenString, &userAccessClaims{}, s.secretAccess)
	if err != nil {
		return nil, fmt.Errorf("sessionService VerifyToken: %w", err)
	}

	accessClaims := claims.(*userAccessClaims)
	return &domain.Token{
		UserId:   accessClaims.UserId,
		UserRole: accessClaims.UserRole,
	}, nil
}

func (s *sessionService) CreateToken(user *domain.User) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userAccessClaims{
		UserId:   user.Id,
		UserRole: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: getNewAccessExpireUnixTime(),
		},
	})

	return jwtToken.SignedString(s.secretAccess)
}

func parseJWT(tokenString string, claims jwt.Claims, secret []byte) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("tokenService parseJWT: parse error: %w", err)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, fmt.Errorf("sessionService parseJWT: validation error: %w", err)
	}

	return claims, nil
}

func getNewAccessExpireUnixTime() int64 {
	return time.Now().Add(time.Hour).Unix()
}
