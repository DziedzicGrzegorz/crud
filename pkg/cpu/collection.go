package cpu

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbHandler interface {
	CreateConnection(context.Context) error
	CloseConnection(context.Context) error
	GetClient() *mongo.Client
	ListDatabases(context.Context) (mongo.ListDatabasesResult, error)
	// CreateCollection(context.Context) error
	CreateDatabase(ctx context.Context) error
}

type dbHandler struct {
	client *mongo.Client
}

func NewDbHandler() DbHandler {
	return &dbHandler{}
}

func (d *dbHandler) CreateConnection(ctx context.Context) error {
	pass := os.Getenv("MONGO_PASS")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://GrzegorzDziedzic:%s@restauth.q17o9nt.mongodb.net/?retryWrites=true&w=majority", pass)).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}
	d.client = client
	return nil
}

func (d *dbHandler) CloseConnection(ctx context.Context) error {
	if d.client != nil {
		return d.client.Disconnect(ctx)
	}
	return nil
}

func (d *dbHandler) GetClient() *mongo.Client {
	return d.client
}

func (d *dbHandler) ListDatabases(ctx context.Context) (mongo.ListDatabasesResult, error) {
	dbList, err := d.client.ListDatabases(ctx, bson.D{})
	if err != nil {
		fmt.Printf("error listing databases: %s", err)
	}
	fmt.Println(dbList)

	return dbList, nil
}

// create database in cloud
func (d *dbHandler) CreateDatabase(ctx context.Context) error {
	err := d.client.Database("cpu").CreateCollection(ctx, "cpu2")

	fmt.Println(err)

	if err != nil {
		return fmt.Errorf("error creating collection: %s", err)
	}
	return nil
}
