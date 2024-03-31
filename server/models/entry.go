package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
    Entry is a struct that represents a single
	entry in the database.
*/
type Entry struct {
	ID          primitive.ObjectID `json:"id"`
	Dish        *string            `json:"dish"`
	Fat         *float64           `json:"fat"`
	Ingredients *string            `json:"ingredients"`
	Calories    *string            `json:"calories"`
}
