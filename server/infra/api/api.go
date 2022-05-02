package api

import (
	"context"
	"crypto/sha256"
	"fmt"
	"jojogo/server/config"
	"jojogo/server/infra/api/db"
	"jojogo/server/jwt"
	"jojogo/server/template"
	"jojogo/server/utils/log"
	_user "jojogo/server/utils/user"
	"reflect"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type group struct {
	Group_name   string             `json:"group_name"`
	Total_member int32              `json:"total_member"`
	Members      primitive.A        `json:"members"`
	Start_time   primitive.DateTime `json:"start_time"`
	Active       bool               `json:"active"`
}

func GetGroups(c *gin.Context) {
	coll := db.Client.Database("groups").Collection("version1")

	cursor, err := coll.Find(context.TODO(), bson.D{}) // {"total_member", bson.D{{"$lte", 500}}}
	if err != nil {
		log.Error("something went wrong", zap.Error(err))
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Error("something went wrong", zap.Error(err))
		panic(err)
	}
	// c.JSON(http.StatusOK, results)

	var groups []group
	for _, result := range results {
		fmt.Println("result['members'] = ", reflect.TypeOf(result["members"]))
		fmt.Println("result['total_member'] = ", reflect.TypeOf(result["total_member"]))

		// one_group := struct {
		// 	Group_name   string
		// 	Total_member int32
		// 	Members      primitive.A
		// 	Start_time   primitive.DateTime
		// 	Active       bool
		// }{
		// 	result["group_name"].(string),
		// 	result["total_member"].(int32),
		// 	result["members"].(primitive.A),
		// 	result["start_time"].(primitive.DateTime),
		// 	result["active"].(bool),
		// }
		// c.JSON(http.StatusOK, one_group)
		one_group := group{
			result["group_name"].(string),
			result["total_member"].(int32),
			result["members"].(primitive.A),
			result["start_time"].(primitive.DateTime),
			result["active"].(bool),
		}
		groups = append(groups, one_group)
	}

	c.JSON(http.StatusOK, groups)
}

func GetGroupByName(c *gin.Context) {
	group_name := c.Param("group_name")

	coll := db.Client.Database("groups").Collection("version1")
	var result bson.M // group_name The gay group
	err := coll.FindOne(context.TODO(), bson.D{{"group_name", group_name}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			log.Error("something went wrong", zap.Error(err))
			panic(err)
		}
	}

	var one_group group

	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &one_group)

	c.JSON(http.StatusOK, one_group)
}

type response struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func CreateGroup(c *gin.Context) {
	group_name := c.Param("group_name")

	coll := db.Client.Database("groups").Collection("version1")
	doc := bson.D{
		{"group_name", group_name},
		{"total_member", 0},
		{"members", []string{}},
		{"start_time", time.Now()},
		{"active", true},
	}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}

	insertion_response := response{
		ID:      result.InsertedID.(primitive.ObjectID).Hex(),
		Message: "success",
	}

	c.JSON(http.StatusOK, insertion_response)
}

func UpdateGroupName(c *gin.Context) {
	set_name := c.Param("set_name")
	search_name := c.Param("search_name")

	coll := db.Client.Database("groups").Collection("version1")

	// id, _ := primitive.ObjectIDFromHex("5eb3d668b31de5d588f42a7a")
	filter := bson.D{{"group_name", search_name}}
	update := bson.D{{"$set", bson.D{{"group_name", set_name}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	type update_response struct {
		Affected string `json:"affcted"`
		Message  string `json:"message"`
	}

	res := update_response{
		Affected: strconv.FormatInt(result.ModifiedCount, 10),
		Message:  "Succesful",
	}

	fmt.Printf("Number of documents replaced or modified: %d", result.ModifiedCount)

	c.JSON(http.StatusOK, res)
}

func AddToGroup(c *gin.Context) {
	person_name := c.Param("person_name")
	search_name := c.Param("search_name")

	coll := db.Client.Database("groups").Collection("version1")

	// id, _ := primitive.ObjectIDFromHex("5eb3d668b31de5d588f42a7a")
	filter := bson.D{{"group_name", search_name}}
	update := bson.D{{"$push", bson.D{{"members", person_name}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	type update_response struct {
		Affected string `json:"affcted"`
		Message  string `json:"message"`
	}

	res := update_response{
		Affected: strconv.FormatInt(result.ModifiedCount, 10),
		Message:  "Succesful",
	}

	fmt.Printf("Number of documents replaced or modified: %d", result.ModifiedCount)

	c.JSON(http.StatusOK, res)
}

func DelFromGroup(c *gin.Context) {
	person_name := c.Param("person_name")
	search_name := c.Param("search_name")

	coll := db.Client.Database("groups").Collection("version1")

	// id, _ := primitive.ObjectIDFromHex("5eb3d668b31de5d588f42a7a")
	filter := bson.D{{"group_name", search_name}}
	update := bson.D{{"$pull", bson.D{{"members", person_name}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	type update_response struct {
		Affected string `json:"affcted"`
		Message  string `json:"message"`
	}

	res := update_response{
		Affected: strconv.FormatInt(result.ModifiedCount, 10),
		Message:  "Succesful",
	}

	fmt.Printf("Number of documents replaced or modified: %d", result.ModifiedCount)

	c.JSON(http.StatusOK, res)
}

func ChangeGroupState(c *gin.Context) {
	rcv_state := c.Param("state")
	search_name := c.Param("search_name")
	state, _ := strconv.ParseBool(rcv_state)

	coll := db.Client.Database("groups").Collection("version1")

	models := []mongo.WriteModel{
		// mongo.NewReplaceOneModel().SetFilter(bson.D{{"title", "Record of a Shriveled Datum"}}).
		// 	SetReplacement(bson.D{{"title", "Dodging Greys"}, {"text", "When there're no matches, no longer need to panic. You can use upsert"}}),
		mongo.NewUpdateOneModel().SetFilter(bson.D{{"group_name", search_name}}).
			SetUpdate(bson.D{{"$set", bson.D{{"active", state}}}}),
	}
	opts := options.BulkWrite().SetOrdered(true)
	results, err := coll.BulkWrite(context.TODO(), models, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			log.Error("something went wrong", zap.Error(err))
			panic(err)
		}
	}

	type update_response struct {
		Affected string `json:"affcted"`
		Message  string `json:"message"`
	}

	res := update_response{
		Affected: strconv.FormatInt(results.ModifiedCount, 10),
		Message:  "Succesful",
	}

	fmt.Printf("Number of documents replaced or modified: %d", results.ModifiedCount)

	c.JSON(http.StatusOK, res)
}

func Register(c *gin.Context) {
	var request template.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Bad Request, ", zap.String("error", "incorrect parameters"))
		template.BadRequest(c, template.ErrParamsCode, "incorrect parameters")
		return
	}

	coll := db.Client.Database("User").Collection("user")
	var check_exist bson.M // group_name The gay group
	err := coll.FindOne(context.TODO(), bson.D{{"name", request.UserName}}).Decode(&check_exist)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			panic(err)
		}
	}
	if len(check_exist) > 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"Message": "The given name is used already",
		})
		return
	}

	h := sha256.New()
	h.Write([]byte(request.Password))

	doc := bson.D{
		{"name", request.UserName},
		{"password", string(h.Sum(nil)[:])},
	}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}

	insertion_response := response{
		ID:      result.InsertedID.(primitive.ObjectID).Hex(),
		Message: "success",
	}

	c.JSON(http.StatusOK, insertion_response)
}

func LoginHandler(c *gin.Context) {
	var request template.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Bad Request, ", zap.String("error", "incorrect parameters"))
		template.BadRequest(c, template.ErrParamsCode, "incorrect parameters")
		return
	}
	user, err := _user.FindUserByUsername(request.UserName)

	if err != nil {
		log.Error("Status Not Found, ", zap.String("error", fmt.Sprintf("user %s not found", request.UserName)))
		template.StatusNotFound(c, template.ErrParamsCode, fmt.Sprintf("user %s not found", request.UserName))
		return
	}

	if !_user.Verify(request.Password, user.Password) {
		log.Error("Incorrect Password")
		template.UnauthorityError(c, template.ErrUnauthorizedCode, "Incorrect Password")
		return
	}

	// // create jwt token
	jwtToken, err := jwt.GenerateToken(user)
	if err != nil {
		log.Error("Cannot generate the Token", zap.Any("error", err))
		template.UnauthorityError(c, template.ErrUnauthorizedCode, "Cannot generate the Token")
		return
	}

	// 測試domain先寫localhost secure先寫false
	c.SetCookie(jwt.Key, jwtToken, config.Val.JWTTokenLife, "/", "localhost", false, true)
	template.Success(c, zap.Any("token", jwtToken))
}

func Init() {
	log.Info("Connecting to database...")
	db.Connect()
	log.Info("Connection to database established.")
}
