package manager

import (
	"fmt"
	"log"

	"github.com/rafialariq/digital-bank/config"
	"github.com/rafialariq/digital-bank/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	ConnectDb() *gorm.DB
}

type infraManager struct {
	db     *gorm.DB
	config config.AppConfig
}

func (i *infraManager) initDb() {
	// dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", i.config.User, i.config.Password, i.config.Host, i.config.Port, i.config.Name, i.config.SslMode)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", i.config.Host, i.config.User, i.config.Password, i.config.Name, i.config.Port, i.config.SslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// logging here
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Merchant{})
	if err != nil {
		// logging here
		log.Fatal(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		// logging here
		log.Fatal(err)
	}

	defer func() {
		if err := recover(); err != nil {
			// logging here
			log.Println("Application filed to run", err)
			sqlDb.Close()
		}
	}()

	i.db = db
	fmt.Println("Connected to DB")
}

func (i *infraManager) ConnectDb() *gorm.DB {
	return i.db
}

func NewInfraManager(config config.AppConfig) InfraManager {
	infra := infraManager{
		config: config,
	}
	infra.initDb()
	return &infra
}
