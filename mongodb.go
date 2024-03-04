package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// manages all the functions
type manager struct {
	connection *mongo.Client      //connection
	Ctx        context.Context    //context
	Cancel     context.CancelFunc //cancellation
}

// the context package is a powerful toool to manage operations like timeouts,cancelation,deadlines,etc.
// Among these operations ,context with timeout is mainly used when we want to make an external request ,such as network request for a database request.
var mgr manager

func connectDb() {
	uri := "localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri))) //checking the connectivity and applying it to the client
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //cancelling the earlier parent context and making a new context
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

	mgr = manager{connection: client, Ctx: ctx, Cancel: cancel}

	fmt.Println("connected!!!!!!!")

}
func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	//CancelFunc to cancel to context
	defer cancel()

	//client provides a method to close
	//a mongoDB connection
	defer func() {
		//client.Disconnect method also has deadline
		//returns error if any
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}

func main() {
	connectDb()
}
