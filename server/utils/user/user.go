package user

import (
	"context"
	"fmt"
	"jojogo/server/infra/api/db"
	"jojogo/server/template"
	"jojogo/server/utils/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func FindUserByUsername(username string) (template.User, error) {
	var user template.User
	coll := db.Client.Database("User").Collection("user")
	var result bson.M
	log.Info(fmt.Sprintf("find name %v", username))
	err := coll.FindOne(context.TODO(), bson.D{{"name", username}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			log.Error("something went wrong(User Name)", zap.Error(err))
			panic(err)
		}
	}

	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &user)
	if err != nil {
		panic(err)
	}
	log.Info("find user successfully, ", zap.Any("result", user))
	return user, nil
}

func FindUserByID(userid string) (template.User, error) {
	var user template.User
	coll := db.Client.Database("User").Collection("user")
	var result bson.M
	log.Info(fmt.Sprintf("find id %v", userid))
	obj_id, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		log.Error("be failed from string to objectID", zap.Error(err))
	}
	err = coll.FindOne(context.TODO(), bson.D{{"_id", obj_id}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			log.Error("something went wrong(User ID)", zap.Error(err))
			return user, err
		}
	}

	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &user)
	if err != nil {
		return user, err
	}
	log.Info("find user successfully, ", zap.Any("result", user))
	return user, nil
}
