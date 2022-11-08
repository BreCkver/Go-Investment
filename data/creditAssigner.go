package data

import (
	"context"
	"log"
	"time"

	"github.com/BreCkver/Go-Investment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	baseName       = "transaction"
	collectionName = "creditAssigner"
)

type CreditAssignerData struct {
}

func NewCreditAssignerData() *CreditAssignerData {
	return &CreditAssignerData{}
}

func (d *CreditAssignerData) CreditAssignmentSummarySave(summary *models.CreditAssignmentSummary) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientDB := Conexion()

	db := clientDB.Database(baseName)
	col := db.Collection(collectionName)
	result, err := col.InsertOne(ctx, summary)
	if err != nil {
		log.Printf("Error guardando info %v", err.Error())
		return "", err
	}

	objIdentifier, _ := result.InsertedID.(primitive.ObjectID)
	return objIdentifier.Hex(), nil
}

func (d *CreditAssignerData) GetLastCreditAssignmentSummary() (models.CreditAssignmentSummary, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientDB := Conexion()

	db := clientDB.Database(baseName)
	col := db.Collection(collectionName)

	var results []models.CreditAssignmentSummary
	var summary models.CreditAssignmentSummary

	options := options.Find()
	condition := bson.M{"is_active": true}
	cursor, err := col.Find(ctx, condition, options)

	if err != nil {
		log.Fatal(err.Error())
		return summary, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&summary)
		if err != nil {
			last := results[len(results)-1]
			return last, err
		}
		results = append(results, summary)
	}

	if len(results) >= 1 {
		return results[len(results)-1], nil
	} else {
		return summary, nil
	}

}

func (d *CreditAssignerData) UpdateLastCreditAssignmentSummary() error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientDB := Conexion()

	db := clientDB.Database(baseName)
	col := db.Collection(collectionName)

	row := make(map[string]interface{})
	row["is_active"] = false
	updString := bson.M{
		"$set": row,
	}

	filter := bson.M{"is_active": bson.M{"$eq": true}}
	_, err := col.UpdateOne(ctx, filter, updString)

	if err != nil {
		return err
	}

	return nil

}
