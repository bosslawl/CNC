package Database

import (
	ParseJson "Rain/core/functions/json"
	Util "Rain/core/functions/util"
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User_Struct struct {
	Username string
	Password string

	Admin           bool
	Powersaving     bool
	Bypassblacklist bool
	Reseller        bool
	Newuser         bool
	Banned          bool
	VIP             bool
	Raw             bool
	Holder          bool
	MaxSessions     int

	MaxTime     int
	Cooldown    int
	Concurrents int
	MFA         bool
	MFAToken    string

	Expiry int64
}

var _ = reflect.TypeOf(User_Struct{})

func GetUser(username string) (*User_Struct, bool) {
	error := ParseJson.Configuration_Parse()
	username = strings.ToLower(username)
	if error != nil {
		fmt.Println(error)
	}
	collection := MCXT.Database("Rain").Collection("Users")
	filter := bson.M{"Username": username}
	var result User_Struct
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		//log.Fatal(err)
		return nil, false
	}
	return &result, true
}

func CreateNewUser(User *User_Struct) error {
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}
	User.Username = strings.ToLower(User.Username)
	collection := MCXT.Database("Rain").Collection("Users")
	doc := bson.M{"Username": User.Username, "Password": User.Password, "Admin": User.Admin, "Powersaving": User.Powersaving, "Bypassblacklist": User.Bypassblacklist, "Reseller": User.Reseller, "Newuser": User.Newuser, "Banned": User.Banned, "VIP": User.VIP, "Raw": User.Raw, "Holder": User.Holder, "MaxTime": User.MaxTime, "Cooldown": User.Cooldown, "Concurrents": User.Concurrents, "Expiry": User.Expiry, "MFA": User.MFA, "MFAToken": User.MFAToken, "MaxSessions": User.MaxSessions}

	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		//log.Fatal(err)
		return nil
	}
	if result != nil {
		return nil
	}
	return nil
}

func AddTime(user string, TimeUnit time.Duration) bool {
	User, error := GetUser(user)
	if !error {
		return false
	}

	Plan_time := time.Unix(User.Expiry, 0)
	Time_Asending := Plan_time.Add(TimeUnit).Unix()

	collection := MCXT.Database("Rain").Collection("Users")
	filter := bson.M{"Username": user}
	update := bson.M{"$set": bson.M{"Expiry": Time_Asending}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		//log.Fatal(err)
		return false
	}
	return true
}

func ChangePassword(User string, Password string, Hash bool) (int, error) {
	User = strings.ToLower(User)
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}
	var LenPassword int = len(Password)
	Pwd := Util.PasswordHash(Password)
	User = strings.ToLower(User)
	collection := MCXT.Database("Rain").Collection("Users")
	filter := bson.M{"Username": User}
	update := bson.M{"$set": bson.M{"Password": Pwd}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}
	return LenPassword, nil
}

func EditFeild(user, feild, replace string, isint bool) bool {
	error := ParseJson.Configuration_Parse()
	user = strings.ToLower(user)
	if error != nil {
		fmt.Println(error)
	}

	if isint {
		i, err := strconv.Atoi(replace)
		if err != nil {
			return false
		}
		collection := MCXT.Database("Rain").Collection("Users")
		filter := bson.M{"Username": user}
		update := bson.M{"$set": bson.M{feild: i}}
		_, err = collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			//log.Fatal(err)
			return false
		}
		return true
	} else {
		if replace == "0" {
			replace = "false"
			replace1, _ := strconv.ParseBool(replace)
			collection := MCXT.Database("Rain").Collection("Users")
			filter := bson.M{"Username": user}
			update := bson.M{"$set": bson.M{feild: replace1}}
			_, err := collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				//log.Fatal(err)
				return false
			}
		} else if replace == "1" {
			replace = "true"
			replace1, _ := strconv.ParseBool(replace)
			collection := MCXT.Database("Rain").Collection("Users")
			filter := bson.M{"Username": user}
			update := bson.M{"$set": bson.M{feild: replace1}}
			_, err := collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				//log.Fatal(err)
				return false
			}
		} else {
			collection := MCXT.Database("Rain").Collection("Users")
			filter := bson.M{"Username": user}
			update := bson.M{"$set": bson.M{feild: replace}}
			_, err := collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				//log.Fatal(err)
				return false
			}
		}
	}
	return true
}

func Remove(username string) bool {
	error := ParseJson.Configuration_Parse()
	username = strings.ToLower(username)
	if error != nil {
		fmt.Println(error)
	}
	collection := MCXT.Database("Rain").Collection("Users")
	filter := bson.M{"Username": username}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		//log.Fatal(err)
		return false
	}
	return true
}

func GetUsers() ([]User_Struct, error) {
	error := ParseJson.Configuration_Parse()
	if error != nil {
		fmt.Println(error)
	}
	collection := MCXT.Database("Rain").Collection("Users")
	var results []User_Struct
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		//log.Fatal(err)
	}
	for cur.Next(context.Background()) {
		var elem User_Struct
		err := cur.Decode(&elem)
		if err != nil {
			//log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	cur.Close(context.Background())
	return results, nil
}
