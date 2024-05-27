package test

import (
	"github.com/MuhamadAndre10/student-profile-service/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestLog(t *testing.T) {
	log := config.NewLogger()

	assert.NotNil(t, log)
	log.Info("Test Info", zap.String("key", "value"))
}
