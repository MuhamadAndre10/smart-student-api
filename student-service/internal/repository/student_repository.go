package repository

import "github.com/MuhamadAndre10/student-profile-service/internal/entity"

type StudentRepository struct {
	Repository[entity.Students]
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}
