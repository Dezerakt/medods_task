package configs

import (
	"fmt"
	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"medods_task/internal/app/models"
	"time"
)

var (
	DbObject *gorm.DB

	TablesList = []interface{}{
		models.User{},
		models.Token{},
	}
)

type DbConfig struct {
}

func NewDbConfig() *DbConfig {
	return &DbConfig{}
}

func (d *DbConfig) ConnectToDatabase(config *EnvConfig) error {
	var err error
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)
	DbObject, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Add(6 * time.Hour)
		},
	})
	if err != nil {
		return err
	}

	log.Println(color.GreenString("Соединение с базой данных успешно"))

	return nil
}
