package db

import (
	"context"
	"fmt"
)

type AbuseStruct struct {
	ID        string   `bson:"_id,omitempty" json:"_id,omitempty"`
	Type      string    `bson:"type,omitempty" json:"type"`
}

func ReadData(query map[string]interface{}) (AbuseStruct, error) {
	active := AbuseStruct{}
	if err := Mdb.Collection(Collections["orders"]).FindOne(context.TODO(), query).Decode(&active); err != nil {
		fmt.Println("-------------err---", err)
	}
	return active, nil
}