package models

import (
	"time"
)

type FeePackage struct {
	ID          int       `gorm:"primary_key;auto_increment"`
	Description string    `gorm:"type:varchar(255);not null"`
	Amount      int       `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FeePackageInput struct {
	Description string    `json:"description" binding:"required"`
	Amount      int       `json:"amount" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
}

type FeePackageResponse struct {
	ID          int       `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	Amount      int       `json:"amount,omitempty"`
	StartDate   time.Time `json:"start_date,omitempty"`
	EndDate     time.Time `json:"end_date,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
