package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

func dbConfig() *mongo.Client {
	var config *options.ClientOptions
	if os.Getenv("MONGO_USERNAME") != "" && os.Getenv("MONGO_PASSWORD") != "" {
		config = options.Client().ApplyURI(
			"mongodb://" + os.Getenv("MONGO_USERNAME") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"),
		)
	} else {
		config = options.Client().ApplyURI(
			"mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"),
		)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, config)
	if err != nil {
		log.Println(err.Error())
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err.Error())
	}
	return client
}

//DB config
func DB() *mongo.Database {
	return dbConfig().Database(os.Getenv("MONGO_DATABASE"))
}
