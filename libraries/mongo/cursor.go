package mongo

import (
	"context"
	"github.com/R3l3ntl3ss/CarJacked/models"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func OfficerCursor(cur *mongo.Cursor) (officers []models.Officer ,err error) {

	var results []models.Officer

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		// Create a value into which the single document can be decoded
		var elem models.Officer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	if err := cur.Close(context.TODO()); err != nil {
		return nil, err
	}

	return results,nil
}

func CaseCursor(cur *mongo.Cursor) (cases []models.Case ,err error) {

	var results []models.Case

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		// Create a value into which the single document can be decoded
		var elem models.Case
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	if err := cur.Close(context.TODO()); err != nil {
		return nil, err
	}

	return results,nil
}