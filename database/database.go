package database

import (
	"context"
	"fmt"
	"log"

	"github.com/bedLad/go-fiber-mongo-hrms/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

func Connect(mongoURI string, dbName string) {
	var err error
	mg.Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Successful")

	// connecting with database
	mg.Db = mg.Client.Database(dbName)
	fmt.Println("Connection instance is ready")

}

func GetCollections() []primitive.M {
	cursor, err := mg.Db.Collection("employees").Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var result bson.M
		if err = cursor.Decode(&results); err != nil {
			log.Fatal(err)
		}

		results = append(results, result)
	}

	fmt.Println("Records successfully fetched")
	defer cursor.Close(context.TODO())
	return results
}

func CreateCollection(employee models.Employee) {
	inserted, err := mg.Db.Collection("employees").InsertOne(context.TODO(), employee)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 employee with id : ", inserted.InsertedID)
}

func GetCollectionByID(employeeId string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(employeeId)
	filter := bson.M{"_id": id}

	var result bson.M
	if err := mg.Db.Collection("employees").FindOne(context.TODO(), filter).Decode(&result); err != nil {
		log.Fatal(err)
	}

	return result
}

func DeleteCollection(employeeId string) {
	id, _ := primitive.ObjectIDFromHex(employeeId)
	filter := bson.M{"_id": id}

	deleted, err := mg.Db.Collection("employees").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Record deleted with id : ", deleted.DeletedCount)
}
