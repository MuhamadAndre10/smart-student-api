package config

import (
	"fmt"
	"github.com/MuhamadAndre10/student-profile-service/internal/entity"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewDatabase(viper *viper.Viper, log *zap.Logger) *gorm.DB {
	username := viper.GetString("PG_USERNAME")
	password := viper.GetString("PG_PASSWORD")
	host := viper.GetString("PG_HOST")
	port := viper.GetInt("PG_PORT")
	database := viper.GetString("PG_DATABASE")
	idleConnection := viper.GetInt("PG_DB_POOL_IDLE")
	maxConnection := viper.GetInt("PG_DB_POOL_MAX")
	maxLifeTimeConnection := viper.GetInt("PG_DB_POOL_MAX_LIFETIME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, username, password, database, port)

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

	log.Info("Connecting to database", zap.String("status", "success"))

	err = migrate(db)
	if err != nil {
		log.Fatal("failed to migrate", zap.Error(err))

	}

	log.Info("Migrating database", zap.String("status", "success"))

	connection, err := db.DB()
	if err != nil {
		log.Fatal("failed to connect database", zap.Error(err))
	}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}

func migrate(db *gorm.DB) error {
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
