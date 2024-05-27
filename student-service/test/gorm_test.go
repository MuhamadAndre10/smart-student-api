package test

import (
	"github.com/MuhamadAndre10/student-profile-service/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGorm(t *testing.T) {

	cfg := config.NewViper()
	log := config.NewLogger()

	db := config.NewDatabase(cfg, log)

	assert.NotNil(t, db)

}
