package models

import (
	"github.com/google/uuid"
	"time"
)

type StudentsResponse struct {
	ID        uuid.UUID `json:"id"`
	StudentID string    `json:"student_id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Gender    string    `json:"gender"`
	Religion  string    `json:"religion"`
	BrithDate string    `json:"brith_date"`
	Address   string    `json:"address"`
	Photo     string    `json:"photo"`
	ClassID   uuid.UUID `json:"class_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InsertStudent struct {
	StudentID string    `json:"student_id" form:"studentID"  validate:"required"`
	FullName  string    `json:"full_name" form:"fullName"  validate:"required"`
	Email     string    `json:"email" form:"email"  validate:"required"`
	Phone     string    `json:"phone" form:"phone"  validate:"required"`
	Gender    string    `json:"gender" form:"gender" validate:"required"`
	Religion  string    `json:"religion" form:"religion" validate:"required"`
	BrithDate string    `json:"brith_date" form:"birthDate" validate:"required"`
	Address   string    `json:"address" form:"address" validate:"required"`
	Photo     string    `json:"photo" form:"photo" validate:"required"`
	ClassID   uuid.UUID `json:"class_id" form:"classID"  validate:"required"`
}
