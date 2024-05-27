package test

import (
	"github.com/MuhamadAndre10/student-profile-service/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewViper(t *testing.T) {
	viper := config.NewViper()

	assert.NotNil(t, viper)
	assert.Equal(t, "student-service", viper.GetString("APP_NAME"))

}
