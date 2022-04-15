package db

import (
	"context"
	"fmt"
	"time"

	"jojogo/server/utils/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	log.Info("start to connet the db...")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://jojogo:jojogo@cluster0.aywk9.mongodb.net/JOJOGO?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error(fmt.Sprintf("%v", err))
		panic(err)
	}

	Client = client

	log.Info("successful!")
}
