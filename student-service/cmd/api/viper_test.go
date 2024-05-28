package main

import (
	"github.com/MuhamadAndre10/student-profile-service/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewViper(t *testing.T) {
	config.InitEnvConfigs(".")

	assert.Equal(t, "student-service", config.Env.AppName)

}
