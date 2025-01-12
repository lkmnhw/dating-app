package databases

import (
	"context"
	"dating-app/app_config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Instance of databases used.
type Instance struct {
	config *app_config.AppConfig

	// core database
	CoreClient    *mongo.Client
	CoreDB        *mongo.Database
	TestingClient *mongo.Client
}

// New database instance constructor.
func New(isProd bool) *Instance {
	di := &Instance{}
	di.config = app_config.Get(isProd)
	return di
}

// ConnectCoreDB: connect database instance of core mysql database.
func (di *Instance) ConnectCoreDB() {
	var err error
	if di.CoreClient == nil {
		di.CoreClient, err = mongo.Connect(context.TODO(), options.Client(), options.Client().ApplyURI(di.config.ENV.DATABASE_SOURCE))
		if err != nil {
			log.Panicf(" (x) database error (connect): cannot connect to core database: %v", err)
		}
	}
	di.CoreDB = di.CoreClient.Database(di.config.ENV.DATABASE_NAME)
}

// CloseCoreDB: close connection to database instance of core mysql database.
func (di *Instance) CloseCoreDB() {
	if di.CoreClient == nil {
		return
	}

	err := di.CoreClient.Disconnect(context.TODO())
	if err != nil {
		log.Println(" (x) error closing connection to core database")
	}
}
