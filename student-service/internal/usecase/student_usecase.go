package usecase

import (
	"context"
	"github.com/MuhamadAndre10/student-profile-service/internal/entity"
	"github.com/MuhamadAndre10/student-profile-service/internal/models"
	"github.com/MuhamadAndre10/student-profile-service/internal/models/converter"
	"github.com/MuhamadAndre10/student-profile-service/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type StudentContract interface {
	Create(db *gorm.DB, entity *entity.Students) error
}

type StudentUseCase struct {
	DB                *gorm.DB
	Log               *zap.Logger
	Validate          *validator.Validate
	StudentRepository StudentContract
}

func NewStudentUseCase(db *gorm.DB, log *zap.Logger, validate *validator.Validate, studentsRepo *repository.StudentRepository) *StudentUseCase {
	return &StudentUseCase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		StudentRepository: studentsRepo,
	}
}

func (c *StudentUseCase) Insert(ctx context.Context, request *models.InsertStudent) (*models.StudentsResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Error("failed to validate request", zap.Error(err))
		return nil, err
	}

	student := &entity.Students{
		StudentID: request.StudentID,
		FullName:  request.FullName,
		Address:   request.Address,
		BrithDate: request.BrithDate,
		Gender:    request.Gender,
		Photo:     request.Photo,
		Religion:  request.Religion,
		ClassID:   request.ClassID,
		Email:     request.Email,
		Phone:     request.Phone,
		CreatedAt: time.Now().Local().UTC(),
	}

	err := c.StudentRepository.Create(tx, student)
	if err != nil {
		c.Log.Error("failed to insert student data", zap.Error(err))
		return nil, fiber.ErrInternalServerError
	}

	c.Log.Info("inserting student data", zap.Any("student", student))

	if err = tx.Commit().Error; err != nil {
		c.Log.Error("failed to commit transaction", zap.Error(err))
		return nil, fiber.ErrInternalServerError
	}

	return converter.StudentToResponse(student), nil

}
