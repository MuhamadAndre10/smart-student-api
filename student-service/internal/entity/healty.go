package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Healthy struct {
	ID                    uuid.UUID `gorm:"column:id;primaryKey;unique;<-false"`
	BloodType             string    `gorm:"column:bloodType;size:5"`
	Allergy               string    `gorm:"column:allergy;size:100"`
	Height                int       `gorm:"column:height"`
	Weight                int       `gorm:"column:weight"`
	EmergencyContactName  string    `gorm:"column:emergencyContactName;size:100"`
	EmergencyContactPhone string    `gorm:"column:emergencyContactPhone;size:15"`
	CreatedAt             time.Time `gorm:"column:createdAt"`
	UpdatedAt             time.Time `gorm:"column:updatedAt"`
}

func (s *Healthy) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid

	return nil
}
