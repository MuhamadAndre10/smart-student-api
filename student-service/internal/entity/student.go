package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Students struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;unique;<-false"`
	StudentID string    `gorm:"column:student_id;unique;not null"`
	FullName  string    `gorm:"column:full_name;size:100;not null"`
	Email     string    `gorm:"column:email;size:50"`
	Phone     string    `gorm:"column:no_phone;size:15;not null"`
	Gender    string    `gorm:"column:gender;type:gender_type;not null"`
	Religion  string    `gorm:"column:religion;not null;size:10"`
	BrithDate string    `gorm:"column:birth_date"`
	Address   string    `gorm:"column:no_phone;size:100"`
	Photo     string    `gorm:"column:photo;size:255"`
	ClassID   uuid.UUID `gorm:"column:class_id;unique;not null"`
	Parents   []Parents `gorm:"foreignKey:StudentID;references:ID"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (s *Students) BeforeCreate(tx *gorm.DB) error {
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	s.ID = uid

	return nil
}
