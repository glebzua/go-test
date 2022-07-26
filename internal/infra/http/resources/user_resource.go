package resources

import (
	"time"

	"github.com/test_server/internal/domain"
)

type UserDto struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Role        uint8     `json:"role"`
	CreatedDate time.Time `json:"created_date"`
}

func MapDomainToUserDto(user *domain.User) *UserDto {
	return &UserDto{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Role:        uint8(user.Role),
		CreatedDate: user.CreatedDate,
	}
}

type TokenDto struct {
	Token string `json:"token"`
}

func MapDomainToTokenDto(token string) *TokenDto {
	return &TokenDto{
		Token: token,
	}
}

func MapDomainToUserDtoCollection(users []domain.User) []UserDto {
	var result []UserDto
	for _, t := range users {
		dto := MapDomainToUserDto(&t)
		result = append(result, *dto)
	}
	return result
}
