package domain

import "time"

type Token struct {
	UserId     int64
	Token      string
	UserRole   Role
	ExpireDate *time.Time
}
