package middleware

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	db *mongo.Database
}

var mongoDB MongoDB

func GetOrInitMongo() (*MongoDB, error) {
	if mongoDB.db != nil {
		return &mongoDB, nil
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	database := client.Database(os.Getenv("DB_NAME"))

	mongoDB = MongoDB{db: database}
	return &mongoDB, nil
}

func CloseMongo() {
	mongo, _ := GetOrInitMongo()
	mongo.db.Client().Disconnect(context.TODO())
}

func CollectionMiddleware(c *fiber.Ctx) error {
	var reqBody struct {
		Collection string `json:"collection"`
	}

	if err := c.BodyParser(&reqBody); err != nil {
		return err
	}

	mongo, _ := GetOrInitMongo()

	collection := mongo.db.Collection(reqBody.Collection)

	c.Context().SetUserValue("collection", collection)

	return c.Next()
}
