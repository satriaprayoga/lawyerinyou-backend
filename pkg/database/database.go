package database

import (
	"fmt"
	"lawyerinyou-backend/pkg/settings"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

func Setup() {
	now := time.Now()
	var err error

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		settings.AppConfigSetting.Database.Host,
		settings.AppConfigSetting.Database.User,
		settings.AppConfigSetting.Database.Password,
		settings.AppConfigSetting.Database.Name,
		settings.AppConfigSetting.Database.Port)
	fmt.Printf("%s", connectionString)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	Conn, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   settings.AppConfigSetting.Database.TablePrefix,
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		log.Printf("connection.setup err : %v", err)
		panic(err)
	}
	sqlDB, err := Conn.DB()
	if err != nil {
		log.Printf("connection.setup DB err : %v", err)
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	//go migrate()

	timeSpent := time.Since(now)
	log.Printf("Config database is ready in %v", timeSpent)
}
