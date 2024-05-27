package models

import (
	"github.com/google/uuid"
	"time"
)

type StudentsResponse struct {
	ID         uuid.UUID `json:"id"`
	StudentsID string    `json:"students_id"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Gender     string    `json:"gender"`
	Religion   string    `json:"religion"`
	BrithDate  time.Time `json:"brith_date"`
	Address    string    `json:"address"`
	Photo      string    `json:"photo"`
	ClassID    uuid.UUID `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type InsertStudent struct {
	StudentsID string    `json:"students_id"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Gender     string    `json:"gender"`
	Religion   string    `json:"religion"`
	BrithDate  time.Time `json:"brith_date"`
	Address    string    `json:"address"`
	Photo      string    `json:"photo"`
	ClassID    uuid.UUID `json:"-"`
}
