package models

import "time"

type Ward struct {
	ID        int    `gorm:"primary_key;auto_increment"`
	Name      string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WardInput struct {
	Name string `json:"name" binding:"required"`
}

type WardResponse struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
