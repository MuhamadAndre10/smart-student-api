package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Students struct {
	ID         uuid.UUID `gorm:"column:id;primaryKey;unique;<-false"`
	StudentsID string    `gorm:"column:schoolNoID;unique;not null"`
	FullName   string    `gorm:"column:fullName;size:100;not null"`
	Email      string    `gorm:"column:email;size:50"`
	Phone      string    `gorm:"column:no_phone;size:15;not null"`
	Gender     string    `gorm:"column:gender;type:gender_type;not null"`
	Religion   string    `gorm:"column:religion;not null;size:10"`
	BrithDate  time.Time `gorm:"column:birthDate"`
	Address    string    `gorm:"column:noPhone;size:100"`
	Photo      string    `gorm:"column:photo;size:255"`
	ClassID    uuid.UUID `gorm:"column:classID;unique;not null"`
	Parents    []Parents `gorm:"foreignKey:StudentID;references:ID"`
	CreatedAt  time.Time `gorm:"column:createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt"`
}

func (s *Students) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid

	return nil
}
