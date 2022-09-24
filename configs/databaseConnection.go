package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBName of the database.
const (
	///////////// Dev Mongo Environemnt Start////////////////////////////////////

	URI    = "mongodb+srv://winston:Win12345@cluster0.ggapjbv.mongodb.net/?retryWrites=true&w=majority"
	DBName = "Entertainment"

	///////////// Dev Mongo Environemnt End////////////////////////////////////

	///////////////// Production Mongo Environemnt Start//////////////////

	// URI    = "mongodb+srv://ecologital-prod:5qyr8iSFtzpUnB6c@production-platform-vwre8.gcp.mongodb.net/test?retryWrites=true&w=majority"
	// DBName = "demandforecastingdata"
	///////////////// Production Mongo Environemnt End //////////////////

)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
		log.Println("[(ConnectDB): Cannot create a client to connect to mongoDB (err):]", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err.Error())
		log.Println("[(ConnectDB): Cannot connect to mongoDB (err):]", err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err.Error())
		log.Println("[(ConnectDB): Cannot ping to mongodb (err):]", err)
	}
	log.Println("[(ConnectDB): Connection Established to MongoDB]")
	return client

}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, DBName string, collectionName string) *mongo.Collection {
	collection := client.Database(DBName).Collection(collectionName)
	return collection
}
