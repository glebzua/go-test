package validators

import (
	"github.com/test_server/internal/domain"
)

type userRequest struct {
	Name     string `json:"name" validate:"required,gte=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Role     uint8  `json:"role_id" validate:"required,numeric"`
}

func mapUserRequestToDomain(request *userRequest) *domain.User {
	return &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     domain.Role(request.Role),
	}
}

type userLogInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func mapUserLogInRequestToDomain(request *userLogInRequest) *domain.User {
	return &domain.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

type userUpdateRequest struct {
	Name  string `json:"name" validate:"required,gte=3"`
	Email string `json:"email" validate:"required,email"`
	Role  uint8  `json:"role_id" validate:"required,numeric"`
}

func mapUserUpdateRequestToDomain(request *userUpdateRequest) *domain.User {
	return &domain.User{
		Name:  request.Name,
		Email: request.Email,
		Role:  domain.Role(request.Role),
	}
}
