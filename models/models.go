package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Salary float64            `json:"salary" bson:"salary"`
	Age    int16              `json:"age" bson:"age"`
}
