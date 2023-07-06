package middleware

import (
	"bufio"
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Context struct {
	*fiber.Ctx
}

type customStreamWriter struct {
	*bufio.Writer
	*mongo.Cursor
	*fiber.Ctx
}

type cursorWriter struct {
	*mongo.Cursor
	*bufio.Writer
}

func (c *cursorWriter) Write() error {
	defer c.Cursor.Close(context.TODO())

	for c.Cursor.Next(context.Background()) {
		_, err := c.Writer.Write(c.Cursor.Current)
		if err != nil {
			return err
		}
		c.Writer.Flush()
	}

	return nil
}

// func (w *customStreamWriter) StreamData() {
// 	log.Printf("CURSOR NIL: %v", w.cursor == nil)

// 	defer w.cursor.Close(context.Background())

// 	// This method will be executed in a separate goroutine
// 	for w.cursor.Next(context.Background()) {
// 		document := make(map[string]interface{})
// 		if err := w.cursor.Decode(&document); err != nil {
// 			// Handle decoding error
// 			fmt.Println("Error decoding document:", err.Error())
// 			break
// 		}

// 		data, err := json.Marshal(document)
// 		if err != nil {
// 			// Handle JSON encoding error
// 			fmt.Println("Error encoding JSON:", err.Error())
// 			break
// 		}

// 		// Stream the data
// 		if _, err := w.c.Write(data); err != nil {
// 			// Handle any error that may occur during streaming
// 			fmt.Println("Error streaming data:", err.Error())
// 			break
// 		}
// 	}
// }

type MongoDB struct {
	*mongo.Database
}

var mongoDB MongoDB

func GetOrInitMongo() (*MongoDB, error) {
	if mongoDB.Database != nil {
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

	mongoDB = MongoDB{database}
	return &mongoDB, nil
}

func CloseMongo() {
	mongo, _ := GetOrInitMongo()
	mongo.Database.Client().Disconnect(context.TODO())
}

func Collection(c *fiber.Ctx) error {
	var reqBody struct {
		Collection string `json:"collection"`
	}

	if err := c.BodyParser(&reqBody); err != nil {
		return err
	}

	mongo, _ := GetOrInitMongo()

	collection := mongo.Database.Collection(reqBody.Collection)

	c.Context().SetUserValue("collection", collection)

	return c.Next()
}

func (c *Context) SimpleRes(res interface{}, err *error) error {
	if *err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(res)
}

func (c *Context) SingleRes(bytes []byte, err *error) error {
	if *err != nil {
		return c.Status(500).JSON(err)
	}

	res := map[string]interface{}{}
	bsonErr := bson.Unmarshal(bytes, res)

	if bsonErr != nil {
		return c.Status(500).JSON(bsonErr)
	}

	return c.Status(200).JSON(res)
}

func (c *Context) BatchRes(cursor *mongo.Cursor, err *error) error {
	if *err != nil {
		c.Status(500).JSON(err)
	}
	defer cursor.Close(context.TODO())

	res := make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var item interface{}
		if err := cursor.Decode(&item); err != nil {
			return c.Status(500).JSON(err)
		}
		res = append(res, item)
	}

	return c.Status(200).JSON(res)
}

func (c *Context) StreamRes(cursor *mongo.Cursor, err *error) error {
	if *err != nil {
		c.Status(500).JSON(err)
	}

	writer := bufio.NewWriter(c.Context().Response.BodyWriter())

	cWriter := cursorWriter{cursor, writer}
	cWriter.Write()

	return nil
}
