package models

import (
	"context"
	"encoding/json"
	"github.com/Ad3bay0c/WebTesting/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Person struct {
	ID			primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname	string				`json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname	string				`json:"lastname,omitempty" bson:"lastname,omitempty"`
	CreatedAt	time.Time			`json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt	time.Time			`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
type Message struct {
	error	string	`json:"error"`
}
var collection = db.Client.Database("testing").Collection("people")

func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)

	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	person.CreatedAt  = time.Now()
	person.UpdatedAt  = time.Now()
	result, _ := collection.InsertOne(ctx, person)

	json.NewEncoder(w).Encode(result)
}

func GetPeopleEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var people []Person
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	result, _ := collection.Find(ctx, bson.M{})

	for result.Next(ctx) {
		var person Person
		_ = result.Decode(&person)
		people = append(people, person)
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(people)
}

func GetPersonEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(Message{error: "ID does not exist"})
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	result := collection.FindOne(ctx, Person{ID: id})
	var person Person
	err = result.Decode(&person)
	if err != nil {
		json.NewEncoder(w).Encode(Message{error: "ID does not exist"})
		return
	}

	json.NewEncoder(w).Encode(person)
}

func UpdatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	json.NewDecoder(r.Body).Decode(&person)

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(Message{error: "ID does not exist"})
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	person.UpdatedAt = time.Now().Local()

	result:= collection.FindOneAndUpdate(ctx, Person{ID: id}, bson.M{"$set": person})
	err = result.Decode(&person)
	if err != nil {
		json.NewEncoder(w).Encode(Message{error: "ID does not exist"})
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(person)
}

func DeletePersonEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(Message{error: "ID does not exist"})
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	res := collection.FindOneAndDelete(ctx, Person{ID: id})
	var person Person

	err = res.Decode(&person)
	if err != nil {
		json.NewEncoder(w).Encode(Message{error: "ID does not exist"})
		return
	}
	json.NewEncoder(w).Encode(Person{ID: person.ID, Firstname: person.Firstname, Lastname: person.Lastname})
}