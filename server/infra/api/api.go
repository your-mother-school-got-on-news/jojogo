package api

import (
	"context"
	"encoding/json"
	"fmt"
	"jojogo/server/config"
	"jojogo/server/infra/api/db"
	"jojogo/server/jwt"
	"jojogo/server/template"
	"jojogo/server/utils/log"
	"jojogo/server/utils/user"
	"strconv"
	"time"

	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

type group struct {
	Group_name   string `json:"group_name"`
	Total_member int32  `json:"total_member"`
	// Members      []string `json:"members"`
	// Start_time string `json:"start_time"`
	Active bool `json:"active"`
}

func GetGroups(c *gin.Context) {
	coll := db.Client.Database("groups").Collection("version1")

	cursor, err := coll.Find(context.TODO(), bson.D{{"total_member", bson.D{{"$lte", 500}}}})
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

	log.Info("result", zap.Any("result", result))

	one_group := group{
		Group_name:   result["group_name"].(string),  // result["group_name"],
		Total_member: result["total_member"].(int32), // result["total_member"],
		// Members:      result["members"].([]string),   // result["members"],
		// Start_time: result["start_time"].(string), // result["start_time"],
		Active: result["active"].(bool), // result["active"],
	}

	c.IndentedJSON(http.StatusOK, one_group)
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

	c.IndentedJSON(http.StatusOK, insertion_response)
}

func UpdateGroupName(c *gin.Context) {
	set_name := c.Param("set_name")
	search_name := c.Param("search_name")

	coll := db.Client.Database("groups").Collection("version1")

	models := []mongo.WriteModel{
		// mongo.NewReplaceOneModel().SetFilter(bson.D{{"title", "Record of a Shriveled Datum"}}).
		// 	SetReplacement(bson.D{{"title", "Dodging Greys"}, {"text", "When there're no matches, no longer need to panic. You can use upsert"}}),
		mongo.NewUpdateOneModel().SetFilter(bson.D{{"group_name", search_name}}).
			SetUpdate(bson.D{{"$set", bson.D{{"group_name", set_name}}}}),
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

	c.IndentedJSON(http.StatusOK, res)
}

func BookById(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func CheckoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := GetBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func ReturnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := GetBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func GetBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func CreateBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println("failed")
		return
	}

	fmt.Println("success")

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func LoginHandler(c *gin.Context) {
	var request template.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Bad Request, ", zap.String("error", "incorrect parameters"))
		template.BadRequest(c, template.ErrParamsCode, "incorrect parameters")
		return
	}
	user, err := user.FindUserByUsername(request.UserName)

	if err != nil {
		log.Error("Status Not Found, ", zap.String("error", fmt.Sprintf("user %s not found", request.UserName)))
		template.StatusNotFound(c, template.ErrParamsCode, fmt.Sprintf("user %s not found", request.UserName))
		return
	}

	if user.Password != request.Password {
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
