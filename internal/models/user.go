package models

import (
	"time"
)

type User struct {
	ID        int64
	Name      string
	AuthID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
