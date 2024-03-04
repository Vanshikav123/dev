package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*go get go.mongodb.org/mongo-driver/x/mongo/driver/topology@v1.14.0
go get go.mongodb.org/mongo-driver/x/mongo/driver/auth@v1.14.0
go get go.mongodb.org/mongo-driver/x/mongo/driver/auth@v1.14.0
*/

// Topic is simple crud mongo

type manager struct {
	connection *mongo.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

var Mgr Manager

type Manager interface {
	Insert(interface{}) error
	GetAll() ([]User, error)
	DeleteData(primitive.ObjectID) error
	UpdateData(User) error
}

func connectDb() {
	uri := "localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected!!!")
	Mgr = &manager{connection: client, ctx: ctx, cancel: cancel}
}

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {
	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func init() {
	connectDb()
}

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}

func main() {
	// insert record to mongodb
	u := User{Name: "go guru", Email: "goguru@gmail.com"}
	err := Mgr.Insert(u)
	fmt.Println(err)

	// get all records in db
	data, err := Mgr.GetAll()

	fmt.Println(data, err)

	// delete record from db
	id := "641e08889d85ada518e83ed1"
	// objectId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// err = Mgr.DeleteData(objectId)
	// fmt.Println(err)

	// update
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	u.ID = objectId
	u.Name = "test"
	u.Email = "test@gmail.com"
	err = Mgr.UpdateData(u)
	fmt.Println(err)
}

func (mgr *manager) Insert(data interface{}) error {
	orgCollection := mgr.connection.Database("goguru123").Collection("collectiongoguru")
	result, err := orgCollection.InsertOne(context.TODO(), data)
	fmt.Println(result.InsertedID)
	return err
}

func (mgr *manager) GetAll() (data []User, err error) {

	orgCollection := mgr.connection.Database("goguru123").Collection("collectiongoguru")

	// Pass these options to the Find method
	findOptions := options.Find()

	cur, err := orgCollection.Find(context.TODO(), bson.M{}, findOptions)
	for cur.Next(context.TODO()) {
		var d User
		err := cur.Decode(&d)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, d)
	} // close for

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return data, nil
}

func (mgr *manager) DeleteData(id primitive.ObjectID) error {
	orgCollection := mgr.connection.Database("goguru123").Collection("collectiongoguru")

	filter := bson.D{{Key: "_id", Value: id}}
	_, err := orgCollection.DeleteOne(context.TODO(), filter)
	return err
}

func (mgr *manager) UpdateData(data User) error {
	orgCollection := mgr.connection.Database("goguru123").Collection("collectiongoguru")

	filter := bson.D{{Key: "_id", Value: data.ID}}
	update := bson.D{{Key: "$set", Value: data}}

	_, err := orgCollection.UpdateOne(context.TODO(), filter, update)

	return err
}

//run => go run main.go
//go to terminal
//mongosh
//show dbs
//use goguru123
//show collections
//db.collectiongoguru.find().pretty()
