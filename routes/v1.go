package routes

import (
	"context"
	"syn-api/middleware"
	"syn-api/types/requests"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func BindV1(app *fiber.App) {
	// Create router group
	v1 := app.Group("/api/v1")

	// Bind middleware
	v1.Use(middleware.CollectionMiddleware)

	// Bind routes
	v1.Post("/find", find)
	v1.Post("/findOne", findOne)
	v1.Post("/insertOne", insertOne)
	v1.Post("/insertMany", insertMany)
	v1.Post("/updateOne", updateOne)
	v1.Post("/updateMany", updateMany)
	v1.Post("/replaceOne", replaceOne)
	v1.Post("/deleteOne", deleteOne)
	v1.Post("/aggregate", aggregate)
}

func find(ctx *fiber.Ctx) error {
	collection := ctx.Context().UserValue("collection").(*mongo.Collection)

	find := new(requests.FindReq)

	if err := ctx.BodyParser(find); err != nil {
		return fiber.ErrBadRequest
	}

	cursor, err := collection.Find(context.TODO(), find.Filter, find.Opts())
	if err != nil {
		ctx.Status(500).JSON(err)
	}
	defer cursor.Close(context.TODO())

	res := make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var item interface{}
		if err := cursor.Decode(&item); err != nil {
			return fiber.ErrBadRequest
		}
		res = append(res, item)
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(200).JSON(res)
}

func findOne(ctx *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func insertOne(ctx *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func insertMany(ctx *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func updateOne(ctx *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func updateMany(ctx *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func replaceOne(ctx *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func deleteOne(ctx *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func aggregate(ctx *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}
