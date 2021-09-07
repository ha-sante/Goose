package goose

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database is a type for structure and predictability of the variables in the connection pool
type Database struct {
	client *mongo.Client
	collections map[string]*mongo.Collection // an object of mongodb collections
}

// DatabaseConnection is for bringing all the required parts of the module to one place
var DatabaseConnection Database;



/*
 databaseURL string = "mongodb://username:password@d1908.mlab.com:39195/database"
 databaseConnection := goose.ConnectDatabase(databaseURL);
*/
func ConnectDatabase(databaseURL string) (*mongo.Client, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))

	if err != nil {
		return nil, err
	}

	defer client.Disconnect(ctx)

	DatabaseConnection.client = client; // Set the client of the connection

	return client, nil
}

/*
 databaseConnection.RegisterDataModel( models.UserAccountModel,  "userAccounts" );
 databaseConnection.RegisterDataModel( models.UserMessagesModel, "userMessages" );
*/
func RegisterDataModel(  bsonData, collectionName string) (*mongo.Collection, error) {

	// Add the collection to map of collections
	err, value := DatabaseConnection.client.Database("testing").Collection(collectionName);
	
	if err!= nil {
		return nil, err
	}

	DatabaseConnection.collections[collectionName] = value;

	log.Print(DatabaseConnection);

	return value, nil
}
