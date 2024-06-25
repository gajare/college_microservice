package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gajare/college_api/db"
	"github.com/gajare/college_api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	// Initialize MongoDB collection once
	client := db.ConnectDB()
	collection = client.Database("my_db").Collection("my_colleges")
}

// GetColleges fetches all colleges
func GetColleges(w http.ResponseWriter, r *http.Request) {
	var colleges []models.College

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var college models.College
		cursor.Decode(&college)
		colleges = append(colleges, college)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(colleges)
}

// GetCollege fetches a single college by ID
func GetCollege(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var college models.College
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&college)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(college)
}

// CreateCollege creates a new college
func CreateCollege(w http.ResponseWriter, r *http.Request) {
	var college models.College
	json.NewDecoder(r.Body).Decode(&college)
	fmt.Println("hello post")
	result, err := collection.InsertOne(context.Background(), college)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

// UpdateCollege updates a college by ID
func UpdateCollege(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var college models.College
	json.NewDecoder(r.Body).Decode(&college)

	update := bson.M{
		"$set": college,
	}

	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("College updated successfully")
}

// DeleteCollege deletes a college by ID
func DeleteCollege(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("College deleted successfully")
}
