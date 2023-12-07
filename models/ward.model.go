package models

import (
	"time"

	"github.com/google/uuid"
)

type Ward struct {
	ID         int    `gorm:"primary_key;auto_increment"`
	Name       string `gorm:"type:varchar(255);not null"`
	WardNumber int    `gorm:"type:int;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CreatedBy  uuid.UUID
	UpdatedBy  uuid.UUID
}

type WardRequest struct {
	Name       string `json:"name" binding:"required"`
	WardNumber int    `json:"ward_number" binding:"required"`
}

type WardResponse struct {
	ID         int       `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	WardNumber int       `json:"ward_number,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedBy  uuid.UUID `json:"created_by"`
	UpdatedBy  uuid.UUID `json:"updated_by"`
}
