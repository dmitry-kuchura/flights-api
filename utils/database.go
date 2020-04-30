package utils

import (
	"context"
	"fmt"
	"log"

	m "flights-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateOne(collection *mongo.Collection, value string, data m.Flight) {
	filter := bson.D{{"number", value}}

	update := bson.D{
		{"$set", bson.D{
			{"boardstatus", data.BoardStatus},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func InsertMany(collection *mongo.Collection, list m.Flight) {
	flights := []interface{}{list}

	insertManyResult, err := collection.InsertMany(context.TODO(), flights)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func InsertOne(collection *mongo.Collection, flight m.Flight) {
	insertResult, err := collection.InsertOne(context.TODO(), flight)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func FindOne(collection *mongo.Collection, value string) (flight m.Flight, err error) {
	filter := bson.D{{"number", value}}

	err = collection.FindOne(context.TODO(), filter).Decode(&flight)

	return flight, err
}

func FindMany(collection *mongo.Collection, value string) ([]*m.Flight, error) {
	if value == "" {
		return FindAll(collection)
	}

	params := options.Find()
	params.SetLimit(100)
	filter := []bson.M{bson.M{"departuretraffichub.code": "(" + value + ")"}, bson.M{"arrivaltraffichub.code": "(" + value + ")"}}

	if value == "KBP" {
		filter = []bson.M{
			bson.M{"departuretraffichub.code": "(" + value + ")"},
			bson.M{"arrivaltraffichub.code": "(" + value + " F)"},
			bson.M{"departuretraffichub.code": "(" + value + ")"},
			bson.M{"arrivaltraffichub.code": "(" + value + " F)"},
		}
	}

	var results []*m.Flight

	cur, err := collection.Find(context.Background(), bson.M{"$or": filter})

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem m.Flight
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, err
}

func FindDeparted(collection *mongo.Collection, value string) ([]*m.Flight, error) {
	if value == "" {
		return FindAll(collection)
	}

	params := options.Find()
	params.SetLimit(100)
	params.SetSort(bson.D{{"departuretime", -1}})
	filter := []bson.M{bson.M{"departuretraffichub.code": "(" + value + ")"}}

	if value == "KBP" {
		filter = []bson.M{
			bson.M{"departuretraffichub.code": "(" + value + ")"},
			bson.M{"departuretraffichub.code": "(" + value + " F)"},
		}
	}

	var results []*m.Flight

	cur, err := collection.Find(context.Background(), bson.M{"$or": filter})

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem m.Flight
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, err
}

func FindArriving(collection *mongo.Collection, value string) ([]*m.Flight, error) {
	if value == "" {
		return FindAll(collection)
	}

	params := options.Find()
	params.SetLimit(100)
	filter := []bson.M{bson.M{"arrivaltraffichub.code": "(" + value + ")"}}

	if value == "KBP" {
		filter = []bson.M{
			bson.M{"arrivaltraffichub.code": "(" + value + ")"},
			bson.M{"arrivaltraffichub.code": "(" + value + " F)"},
		}
	}

	var results []*m.Flight

	cur, err := collection.Find(context.Background(), bson.M{"$and": filter})

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem m.Flight
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, err
}

func FindAll(collection *mongo.Collection) ([]*m.Flight, error) {
	params := options.Find()
	params.SetLimit(100)

	filter := bson.M{}

	var results []*m.Flight

	cur, err := collection.Find(context.Background(), filter, params)

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem m.Flight
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, err
}
