package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Parents struct {
	ID         uuid.UUID `gorm:"column:id;primaryKey;unique;<-false"`
	FullName   string    `gorm:"column:fullName;size:100;not null"`
	Email      string    `gorm:"column:email;size:50"`
	Phone      string    `gorm:"column:no_phone;size:15;not null"`
	Status     string    `gorm:"column:status;size:10;not null"`
	Employment string    `gorm:"column:employment;size:50"`
	Address    string    `gorm:"column:address;size:100"`
	StudentID  uuid.UUID `gorm:"column:studentID;not null"`
	HealthyID  uuid.UUID `gorm:"column:healthyID;not null"`
	Healthy    Healthy   `gorm:"foreignKey:HealthyID;references:ID"`
	CreatedAt  time.Time `gorm:"column:createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt"`
}

func (s *Parents) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid

	return nil
}
