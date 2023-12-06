package models

import "time"

type Member struct {
	ID            int    `gorm:"primary_key;auto_increment"`
	Name          string `gorm:"type:varchar(255);not null"`
	Email         string `gorm:"uniqueIndex;not null"`
	Phone         string `gorm:"uniqueIndex;not null"`
	Photo         string `gorm:"not null"`
	DepartmentID  int    `gorm:"not null"` // masjidId,classNo,
	Department    string `gorm:"not null"` //masjid,class,
	DateOfBirth   string `gorm:"not null"`
	Age           int    `gorm:"not null"`
	Gender        string `gorm:"not null"`
	Address       string `gorm:"type:text;not null"`
	WardID        int    `gorm:"not null"` // ward("edackode","mamom"),class(5,6,7)
	BloodGroup    string `gorm:"not null"`
	ParentId      int    `gorm:"not null"`
	MarrigeStatus string `gorm:"not null"`
	Occupation    string `gorm:"not null"`
	Education     string `gorm:"not null"`
	IsActive      bool   `gorm:"not null"`
	PendingAmount int    `gorm:"not null"`
	FeePackageId  int    `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type MemberInput struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" `
	Phone         string `json:"phone" binding:"required"`
	Photo         string `json:"photo"`
	DepartmentID  int    `json:"department_id"`
	Department    string `json:"department" binding:"required"`
	DateOfBirth   string `json:"date_of_birth" binding:"required"`
	Age           int    `json:"age"`
	Gender        string `json:"gender" binding:"required"`
	Address       string `json:"address" binding:"required"`
	WardID        int    `json:"ward_id" binding:"required"`
	BloodGroup    string `json:"blood_group"`
	ParentId      int    `json:"parent_id"`
	MarrigeStatus string `json:"marrige_status"`
	Occupation    string `json:"occupation"`
	Education     string `json:"education"`
	IsActive      bool   `json:"is_active"`
	PendingAmount int    `json:"pending_amount"`
	FeePackageId  int    `json:"fee_package_id"`
}

type MemberResponse struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Email         string    `json:"email,omitempty"`
	Phone         string    `json:"phone,omitempty"`
	Photo         string    `json:"photo,omitempty"`
	DepartmentID  int       `json:"department_id,omitempty"`
	Department    string    `json:"department,omitempty"`
	DateOfBirth   string    `json:"date_of_birth,omitempty"`
	Age           int       `json:"age,omitempty"`
	Gender        string    `json:"gender,omitempty"`
	Address       string    `json:"address,omitempty"`
	WardID        int       `json:"ward_id,omitempty"`
	BloodGroup    string    `json:"blood_group,omitempty"`
	ParentId      int       `json:"parent_id,omitempty"`
	MarrigeStatus string    `json:"marrige_status,omitempty"`
	Occupation    string    `json:"occupation,omitempty"`
	Education     string    `json:"education,omitempty"`
	IsActive      bool      `json:"is_active,omitempty"`
	PendingAmount int       `json:"pending_amount,omitempty"`
	FeePackageId  int       `json:"fee_package_id,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
