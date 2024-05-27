package converter

import (
	"github.com/MuhamadAndre10/student-profile-service/internal/entity"
	"github.com/MuhamadAndre10/student-profile-service/internal/models"
)

func StudentToResponse(students *entity.Students) *models.StudentsResponse {
	return &models.StudentsResponse{
		ID:         students.ID,
		StudentsID: students.StudentsID,
		FullName:   students.FullName,
		Address:    students.Address,
		BrithDate:  students.BrithDate,
		Gender:     students.Gender,
		Photo:      students.Photo,
		Religion:   students.Religion,
		ClassID:    students.ClassID,
		Email:      students.Email,
		Phone:      students.Phone,
		CreatedAt:  students.CreatedAt,
		UpdatedAt:  students.UpdatedAt,
	}
}
