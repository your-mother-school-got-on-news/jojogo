package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"jojogo/server/utils/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func Connect() {
	log.Info("successfully running")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://jojogo:jojogo@cluster0.aywk9.mongodb.net/JOJOGO?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error(fmt.Sprintf("%v", err))
	}

	coll := client.Database("groups").Collection("version1")
	var result bson.M // group_name The gay group
	err = coll.FindOne(context.TODO(), bson.D{{"group_name", "gay group"}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			log.Error("something went wrong", zap.Error(err))
			panic(err)
		}
	}

	log.Info("result", zap.Any("result", result))

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	log.Info("output", zap.Any("output", output))

	cursor, err := coll.Find(context.TODO(), bson.D{{"total_member", bson.D{{"$lte", 5}}}})
	if err != nil {
		log.Error("something went wrong", zap.Error(err))
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Error("something went wrong", zap.Error(err))
		panic(err)
	}
	for _, result := range results {
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Error("something went wrong", zap.Error(err))
			panic(err)
		}
		log.Info(string(output))
	}

	log.Info("successfully ending")
}
