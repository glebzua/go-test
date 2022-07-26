package validators

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/test_server/internal/domain"
)

type UserValidator struct {
	validator *validator.Validate
}

func NewUserValidator() *UserValidator {
	return &UserValidator{
		validator: validator.New(),
	}
}

func (t UserValidator) ValidateAndMap(r *http.Request) (*domain.User, error) {
	var userResource userRequest
	err := json.NewDecoder(r.Body).Decode(&userResource)
	if err != nil {
		return nil, fmt.Errorf("userValidator ValidateAndMap: %w", err)
	}

	err = t.validator.Struct(userResource)
	if err != nil {
		return nil, fmt.Errorf("userValidator ValidateAndMap: %w", err)
	}

	return mapUserRequestToDomain(&userResource), nil
}

type UserLogInValidator struct {
	validator *validator.Validate
}

func NewUserLogInValidator() *UserLogInValidator {
	return &UserLogInValidator{
		validator: validator.New(),
	}
}

func (t UserLogInValidator) ValidateAndMap(r *http.Request) (*domain.User, error) {
	var userResource userLogInRequest
	err := json.NewDecoder(r.Body).Decode(&userResource)
	if err != nil {
		return nil, fmt.Errorf("loginValidator ValidateAndMap: %w", err)
	}

	err = t.validator.Struct(userResource)
	if err != nil {
		return nil, fmt.Errorf("loginValidator ValidateAndMap: %w", err)
	}

	return mapUserLogInRequestToDomain(&userResource), nil
}

type UserUpdateValidator struct {
	validator *validator.Validate
}

func NewUserUpdateValidator() *UserUpdateValidator {
	return &UserUpdateValidator{
		validator: validator.New(),
	}
}

func (t UserUpdateValidator) ValidateAndMap(r *http.Request) (*domain.User, error) {
	var userResource userUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&userResource)
	if err != nil {
		return nil, fmt.Errorf("userValidator ValidateAndMap: %w", err)
	}
	err = t.validator.Struct(userResource)
	if err != nil {
		return nil, fmt.Errorf("userValidator ValidateAndMap: %w", err)
	}

	return mapUserUpdateRequestToDomain(&userResource), nil
}
