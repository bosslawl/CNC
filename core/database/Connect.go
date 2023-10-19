package Database

import (
	"context"
	"fmt"
	"log"

	ParseJson "Rain/core/functions/json"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MCXT = Connect()

func Connect() *mongo.Client {
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}
	//clientOptions := options.Client().ApplyURI("" + Build_Json.Meta_build_Configuration.Database.MongoURL + "")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(""+ParseJson.ConfigParse.Database.MongoURL+""))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	return client
}
