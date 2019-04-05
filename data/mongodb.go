package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type key string

const (
	hostKey         = key("hostKey")
	usernameKey     = key("usernameKey")
	passwordKey     = key("passwordKey")
	databaseNameKey = key("databaseNameKey")
)

//Repository should be used to retrieve data
var Repository *UrsaRepo

func init() {
	const (
		host         = "localhost"
		username     = "morty"
		password     = "P@ssw0rd"
		databaseName = "ursadb"
	)

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, hostKey, host)
	ctx = context.WithValue(ctx, usernameKey, username)
	ctx = context.WithValue(ctx, passwordKey, password)
	ctx = context.WithValue(ctx, databaseNameKey, databaseName)

	db, err := configDb(ctx)

	if err != nil {
		log.Fatalf("ursa: db configuration failed: %v", err)
	}

	Repository = newRepository(db)
}

func configDb(ctx context.Context) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		ctx.Value(usernameKey).(string),
		ctx.Value(passwordKey).(string),
		ctx.Value(hostKey).(string),
		ctx.Value(databaseNameKey).(string))

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("ursa: could not connect to mongo: %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("ursa: mongo client could not connect to mongo: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("ursa: ursa client could not ping from server: %v", err)
	}

	db := client.Database(ctx.Value(databaseNameKey).(string))

	return db, nil
}

func newRepository(db *mongo.Database) *UrsaRepo {
	return &UrsaRepo{Db: db}
}
