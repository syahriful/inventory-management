package model

import (
	"time"
)

type User struct {
	ID        int64
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
