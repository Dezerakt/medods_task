package configs

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"medods_task/internal/app/models"
	"time"
)

var (
	DbObject    *gorm.DB
	MongoObject *mongo.Client

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

func (d *DbConfig) ConnectToMongo(config *EnvConfig) error {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", config.MongoUser, config.MongoPassword, config.MongoHost, config.MongoPort))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	MongoObject = client

	log.Println(color.GreenString("Соединение с Mongo успешно"))

	return nil
}
