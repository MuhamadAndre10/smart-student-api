package http

import (
	"fmt"
	"github.com/MuhamadAndre10/student-profile-service/internal/models"
	"github.com/MuhamadAndre10/student-profile-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"path/filepath"
	"time"
)

type StudentHandler struct {
	Log     *zap.Logger
	UseCase *usecase.StudentUseCase
}

func NewStudentHandler(log *zap.Logger, useCase *usecase.StudentUseCase) *StudentHandler {
	return &StudentHandler{
		Log:     log,
		UseCase: useCase,
	}
}

func (s *StudentHandler) Insert(c *fiber.Ctx) error {

	// parsing to struct from request body
	request := new(models.InsertStudent)
	if err := c.BodyParser(request); err != nil {
		// write log
		s.Log.Error("failed to parse request body", zap.Error(err))
		return fiber.ErrBadRequest
	}

	file, err := c.FormFile("photo")
	if err != nil || file.Size > 5<<20 || !isAllowedFileType(file.Header.Get(fiber.HeaderContentType)) {
		s.Log.Error("file error", zap.String("error", "maybe file is large or not type allowed or not found"))
		return fiber.ErrBadRequest
	}

	newFile := fmt.Sprintf("photo_%d%s", time.Now().Unix(), filepath.Ext(file.Filename))

	// save file
	if err = c.SaveFile(file, fmt.Sprintf("./file_uploads/%s", newFile)); err != nil {
		s.Log.Error("failed to save file", zap.Error(err))
		return fiber.ErrBadRequest
	}

	request.Photo = newFile

	// insert
	response, err := s.UseCase.Insert(c.Context(), request)
	if err != nil {
		s.Log.Error("failed to insert student", zap.Error(err))
		return err
	}

	//return data with code
	return c.Status(fiber.StatusOK).JSON(models.WebResponse[*models.StudentsResponse]{
		Data: response,
	})

}

func isAllowedFileType(contentType string) bool {
	allowedExtension := []string{"image/jpeg", "image/png", "image/jpg"}
	for _, allowedExt := range allowedExtension {
		if contentType == allowedExt {
			return true
		}
	}
	return false
}
