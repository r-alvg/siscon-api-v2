package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"siscon/configs"
	"sync"
)

type MongoDatabase struct {
	con *mongo.Database
}
type Repository struct {
	*mongo.Collection
}

var (
	DB       *mongo.Database
	instance *MongoDatabase
	once     sync.Once
)

func Start() {
	DB = connect()
}

func connect() *mongo.Database {
	once.Do(func() {
		ctx := context.Background()
		clientOpt := options.Client().ApplyURI(configs.Properties.Db.Uri)
		client, err := mongo.NewClient(clientOpt)

		if err != nil {
			logrus.Error(err)
		}
		if err := client.Connect(ctx); err != nil {
			logrus.Fatal("error on connect")
		}
		if err := client.Ping(context.TODO(), nil); err != nil {
			logrus.Fatal("error on ping cause: ", err)
		}

		instance = &MongoDatabase{client.Database("siscon")}
	})
	return instance.con
}

func NewRepository(name string) Repository {
	return Repository{DB.Collection(name)}
}

