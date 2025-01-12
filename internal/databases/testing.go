package databases

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectTesting: connect to databases in testing mode.
func (di *Instance) ConnectTesting() {
	var err error
	if di.TestingClient == nil {
		di.TestingClient, err = mongo.Connect(context.TODO(), options.Client(), options.Client().ApplyURI(di.config.ENV.TESTING_DATABASE_SOURCE))
		if err != nil {
			log.Panicf(" (x) database error (connect): cannot connect to test database: %v", err)
		}
	}
	db := di.TestingClient.Database(di.config.ENV.TESTING_DATABASE_NAME)
	di.CoreDB = db
}

// CloseTesting: close connection to database instance of test mongodb database.
func (di *Instance) CloseTesting() {
	err := di.TestingClient.Disconnect(context.TODO())
	if err != nil {
		log.Println(" (x) error closing connection to test database")
	}
}
