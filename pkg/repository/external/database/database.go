package database

import (
	"context"
	"core-users/pkg/configs"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Credential struct {
	Host string
	Port string
}

func NewDatabaseMongo(config configs.DataBase) *mongo.Database {
	// Conectar a MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	credential := getCredential(config)
	url := fmt.Sprintf("mongodb://%s:%s", credential.Host, credential.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Seleccionar la base de datos
	db := client.Database(config.Name)
	return db
}

func getCredential(config configs.DataBase) *Credential {
	host := os.Getenv(config.Host)
	port := os.Getenv(string(config.Port))

	return &Credential{
		Host: host,
		Port: port,
	}

}
