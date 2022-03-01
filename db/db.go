package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mdb *mongo.Database
var Mconn *mongo.Client
var mongoCtx context.Context

type FindOptions struct {
	Limit      int64
	Skip       int64
	Sort       primitive.D
	Projection primitive.M
	Hint primitive.M
}

type CountOptions struct {
	Limit      int64
	Skip       int64
	Hint primitive.M
}

type UpdateOptions struct {
	Upsert bool
}

type UpdateOneResponseStruct struct {
	MatchedCount  int64
	ModifiedCount int64
	UpsertedCount int64
}

func Connect() {

	mongoCtx = context.Background()
	db := os.Getenv("db")
	fmt.Println("-----url=", db)
	Mconn, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(db))
	if err != nil {
		log.Fatal(err)
	}

	err = Mconn.Ping(mongoCtx, nil)
	fmt.Println("-----err=", err)

	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
	Mdb = Mconn.Database("oms")
	
}
