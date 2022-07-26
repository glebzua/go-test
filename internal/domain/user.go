package domain

import (
	"time"
)

type Role int

const (
	ROLE_ADMIN     Role = 1
	ROLE_MODERATOR Role = 2
)

type User struct {
	Id          int64
	Name        string
	Email       string
	Password    string
	Passhash    []byte
	Role        Role
	CreatedDate time.Time
	DeletedDate time.Time
}
