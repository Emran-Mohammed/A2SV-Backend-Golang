package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
        return nil, err
    }

	if err = client.Ping(ctx, nil); err != nil{
		return nil, err
	}
	// db := client.Database("taskdb")
	// TaskCollection = db.Collection("tasks")
	return client, nil
	
}

