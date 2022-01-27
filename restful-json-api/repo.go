package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var currentId int

//var todos Todos

var ctx context.Context
var client *mongo.Client
var collection *mongo.Collection

// Give us some seed data
func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx = context.TODO()

	var err error
	// Connect to MongoDB
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection = client.Database("test").Collection("todos")

	RepoCreateTodo(Todo{Name: "Write presentation", Id: 11})
	RepoCreateTodo(Todo{Name: "Host meetup", Id: 12})
}

/* func deinit() {
	// Close the connection once no longer needed
	err := client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
} */
func RepoFindTodo(id int) Todo {
	return RepoFindTodoMongoDB(ctx, collection, id)
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t Todo) Todo {
	return RepoCreateTodoMongoDB(ctx, collection, t)
}

//this is bad, I don't think it passes race condtions
func RepoListTodo() Todos {
	return RepoListTodoMongoDB(ctx, collection)
}

func RepoDestroyTodo(id int) error {
	return RepoDestroyTodoMongoDB(ctx, collection, id)
}

func RepoFindTodoMongoDB(ctx context.Context, collection *mongo.Collection, id int) Todo {
	filter := bson.D{{"id", id}}

	// Find a single document
	var result Todo
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		// return empty Todo if not found
		return Todo{}
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result
}

//this is bad, I don't think it passes race condtions
func RepoListTodoMongoDB(ctx context.Context, collection *mongo.Collection) Todos {

	findOptions := options.Find()
	//findOptions.SetLimit(2)

	var results Todos

	// Finding multiple documents returns a cursor
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(ctx) {
		var elem Todo
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(ctx)

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results

}

//this is bad, I don't think it passes race condtions
func RepoCreateTodoMongoDB(ctx context.Context, collection *mongo.Collection, t Todo) Todo {
	//currentId += 1
	//t.Id = currentId

	// Insert a single document
	insertResult, err := collection.InsertOne(ctx, t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	//id := fmt.Sprintf("%v", insertResult.InsertedID)
	return t
}

func RepoDestroyTodoMongoDB(ctx context.Context, collection *mongo.Collection, id int) error {
	filter := bson.D{{"id", id}}

	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("could not find todo with id of %d to delete", id)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	return nil
}
