package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	fmt.Println("successfully running")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://jojogo:jojogo@cluster0.aywk9.mongodb.net/JOJOGO?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("groups").Collection("version1")
	var result bson.M // group_name The gay group
	err = coll.FindOne(context.TODO(), bson.D{{"group_name", "gay group"}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			fmt.Println("something went wrong")
			fmt.Println("error message: ", err)
			return
		}
		panic(err)
	}

	fmt.Println(result)

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)

	cursor, err := coll.Find(context.TODO(), bson.D{{"total_member", bson.D{{"$lte", 5}}}})
	if err != nil {
		fmt.Println("something went wrong")
		fmt.Println("error message: ", err)
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	fmt.Println("successfully ending")
}
