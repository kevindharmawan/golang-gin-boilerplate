package models

import (
	"time"
)

type Profile struct {
	ID        string    `json:"user_id" gorm:"type:varchar(30);unique;primary_key"`
	Name      string    `json:"name" gorm:"not null"`
	BirthDate time.Time `json:"birth_date" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;autoUpdateTime"`
}

type ProfileCreate struct {
	Name      string    `json:"name" binding:"required"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
}

type ProfileUpdate struct {
	Name      string    `json:"name" binding:"-"`
	BirthDate time.Time `json:"birth_date" binding:"-"`
}
