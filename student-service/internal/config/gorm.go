package config

import (
	"fmt"
	"github.com/MuhamadAndre10/student-profile-service/internal/entity"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

func NewDatabase(v *viper.Viper) *gorm.DB {
	logFile := "log/db.log"
	file, _ := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0222)
	defer file.Close()

	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{logFile}
	log, err := config.Build()
	if err != nil {
		panic(fmt.Errorf("Fatal error zap logger: %w \n", err))
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		v.GetString("PG_HOST"), v.GetString("PG_USERNAME"), v.GetString("PG_PASSWORD"), v.GetString("PG_DATABASE"), v.GetInt("PG_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(&logWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})

	if err != nil {
		log.Fatal("failed to connect database", zap.Error(err))
	}

	migrate(db)

	log.Info("Connecting to database", zap.String("status", "success"))

	conn, err := db.DB()
	if err != nil {
		log.Fatal("failed to connect database", zap.Error(err))
	}

	conn.SetMaxIdleConns(v.GetInt("PG_DB_POOL_IDLE"))
	conn.SetMaxOpenConns(v.GetInt("PG_DB_POOL_MAX"))
	conn.SetConnMaxLifetime(time.Second * time.Duration(v.GetInt("PG_DB_POOL_MAX_LIFETIME")))

	return db
}

func migrate(db *gorm.DB) error {
	_ = db.Migrator().DropTable(&entity.Students{}, &entity.Parents{}, &entity.Healthy{})

	err := db.AutoMigrate(&entity.Students{}, &entity.Parents{}, &entity.Healthy{})
	if err != nil {
		return err
	}

	return nil
}

type logWriter struct {
	Logger *zap.Logger
}

func (l *logWriter) Printf(message string, args ...interface{}) {
	l.Logger.Info(fmt.Sprintf(message, args...) + "\n")
}
