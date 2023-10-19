package Database

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	ParseJson "Rain/core/functions/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Attack struct {
	API      string
	Method   string
	Target   string
	Port     int
	Duration int

	Username string
	End      int64
	Created  int64
}

type Running struct {
	Running int
}

var _ = reflect.TypeOf(Attack{})
var _ = reflect.TypeOf(Running{})

func GetAPI_RunningMethod(API string, Method string) (int, error) {
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}

	collection := MCXT.Database("Rain").Collection("Attacks")
	filter := bson.M{"API": API, "Method": Method}
	var result Running
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		//log.Fatal(err)
		return 0, nil
	}
	return result.Running, nil
}

func GetAPI_Running(API string) (int, error) {
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}
	collection := MCXT.Database("Rain").Collection("Attacks")
	filter := bson.M{"API": API, "End": bson.M{"$gt": time.Now().Unix()}}
	count, err := collection.CountDocuments(context.TODO(), filter)
	s := int(count)
	if err != nil {
		//log.Fatal(err)
		return 0, nil
	}
	return s, nil
}

func LogAttack(Attacks *Attack) error {
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}
	collection := MCXT.Database("Rain").Collection("Attacks")
	doc := bson.M{"API": Attacks.API, "Method": Attacks.Method, "Target": Attacks.Target, "Port": Attacks.Port, "Duration": Attacks.Duration, "Username": Attacks.Username, "End": Attacks.End, "Created": Attacks.Created}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

func GetRunningUser(User string) (int, error) {
	error := ParseJson.Configuration_Parse()
	User = strings.ToLower(User)
	if error != nil {
		fmt.Println(error)
	}
	collection := MCXT.Database("Rain").Collection("Attacks")
	filter := bson.M{"Username": User, "End": bson.M{"$gt": time.Now().Unix()}}
	count, err := collection.CountDocuments(context.TODO(), filter)
	s := int(count)
	if err != nil {
		return 0, nil
	}
	return s, nil
}

func AlreadyUnderAttack(User string, Target string) (*Attack, error) {
	User = strings.ToLower(User)
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}
	collection := MCXT.Database("Rain").Collection("Attacks")
	filter := bson.M{"Target": Target, "End": bson.M{"$gt": time.Now().Unix()}}
	var result Attack
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		//log.Fatal(err)
		return nil, nil
	}
	if time.Now().Unix() > result.End {
		return nil, nil
	}
	return &result, nil
}

func OngoingUser(User string) ([]*Attack, error) {
	User = strings.ToLower(User)
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}
	collection := MCXT.Database("Rain").Collection("Attacks")
	filter := bson.M{"Username": User, "End": bson.M{"$gt": time.Now().Unix()}}
	var result []*Attack
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var attack Attack
		err := cursor.Decode(&attack)
		if err != nil {
			//log.Fatal(err)
			return nil, err
		}
		result = append(result, &attack)
	}
	if err := cursor.Err(); err != nil {
		//log.Fatal(err)
		return nil, err
	}
	cursor.Close(context.TODO())
	return result, nil
}

func Ongoing(User string) ([]*Attack, error) {
	User = strings.ToLower(User)
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}
	collection := MCXT.Database("Rain").Collection("Attacks")
	filter := bson.M{"End": bson.M{"$gt": time.Now().Unix()}}
	var result []*Attack
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var attack Attack
		err := cursor.Decode(&attack)
		if err != nil {
			//log.Fatal(err)
			return nil, err
		}
		result = append(result, &attack)
	}
	if err := cursor.Err(); err != nil {
		//log.Fatal(err)
		return nil, err
	}
	cursor.Close(context.TODO())
	return result, nil
}

func MyAttacking(User string) ([]*Attack, error) {
	User = strings.ToLower(User)
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}
	collection := MCXT.Database("Rain").Collection("Attacks")
	opts := options.Find().SetSort(bson.D{{"_id", -1}})
	filter := bson.M{"Username": User}
	var result []*Attack
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var attack Attack
		err := cursor.Decode(&attack)
		if err != nil {
			return nil, err
		}
		result = append(result, &attack)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.TODO())
	return result, nil
}
