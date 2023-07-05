package routes

import (
	"context"
	"fmt"
	"syn-api/middleware"
	"syn-api/types/requests"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func BindV1(app *fiber.App) {
	// Create router group
	v1 := app.Group("/v1")

	// Bind pre-process middleware
	v1.Use(middleware.Collection)

	// Bind routes
	v1.Post("/find", find)
	v1.Post("/find/stream", findStream)
	v1.Post("/findOne", findOne)
	v1.Post("/insertOne", insertOne)
	v1.Post("/insertMany", insertMany)
	v1.Post("/updateOne", updateOne)
	v1.Post("/updateMany", updateMany)
	v1.Post("/replaceOne", replaceOne)
	v1.Post("/deleteOne", deleteOne)
	v1.Post("/aggregate", aggregate)

	// Bind post-process middleware
	v1.Use(middleware.ReturnsJSON)
}

func find(ctx *fiber.Ctx) error {
	c := &middleware.Context{Ctx: ctx}
	collection := c.Context().UserValue("collection").(*mongo.Collection)

	req := new(requests.Find)
	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	cursor, err := collection.Find(context.TODO(), req.Filter, &req.FindOptions)
	return c.BatchRes(cursor, &err)
}

func findStream(ctx *fiber.Ctx) error {
	return fiber.ErrNotImplemented

	// c := &middleware.Context{Ctx: ctx}
	// collection := c.Context().UserValue("collection").(*mongo.Collection)

	// req := new(requests.Find)
	// if err := c.BodyParser(req); err != nil {
	// 	return fiber.ErrBadRequest
	// }

	// cursor, err := collection.Find(context.TODO(), req.Filter, &req.FindOptions)
	// return c.StreamRes(cursor, &err)
}

func findOne(ctx *fiber.Ctx) error {
	c := &middleware.Context{Ctx: ctx}

	collection := c.Context().UserValue("collection").(*mongo.Collection)

	req := new(requests.FindOne)
	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	fmt.Printf("%v", req.Filter)
	fmt.Printf("%v", req.FindOneOptions)

	var res map[string]interface{}
	err := collection.FindOne(context.TODO(), req.Filter, &req.FindOneOptions).Decode(res)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(res)
}

func insertOne(ctx *fiber.Ctx) error {
	c := &middleware.Context{Ctx: ctx}
	collection := c.Context().UserValue("collection").(*mongo.Collection)

	req := new(requests.InsertOne)
	if err := c.BodyParser(req); err != nil {
		fmt.Printf("%v", err)
		return fiber.ErrBadRequest
	}

	res, err := collection.InsertOne(context.TODO(), req.Document, &req.InsertOneOptions)
	return c.SingleRes(res, &err)
}

func insertMany(ctx *fiber.Ctx) error {
	c := &middleware.Context{Ctx: ctx}
	collection := c.Context().UserValue("collection").(*mongo.Collection)

	req := new(requests.InsertMany)
	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := collection.InsertMany(context.TODO(), req.Documents, &req.InsertManyOptions)
	return c.SingleRes(res, &err)
}

func updateOne(ctx *fiber.Ctx) error {
	c := &middleware.Context{Ctx: ctx}
	collection := c.Context().UserValue("collection").(*mongo.Collection)

	req := new(requests.Update)
	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := collection.UpdateOne(context.TODO(), req.Filter, req.Document, &req.UpdateOptions)
	return c.SingleRes(res, &err)
}

func updateMany(ctx *fiber.Ctx) error {
	c := &middleware.Context{Ctx: ctx}
	collection := c.Context().UserValue("collection").(*mongo.Collection)

	req := new(requests.Update)
	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := collection.UpdateMany(context.TODO(), req.Filter, req.Document, &req.UpdateOptions)
	return c.SingleRes(res, &err)
}

func replaceOne(ctx *fiber.Ctx) error {
	c := &middleware.Context{Ctx: ctx}
	collection := c.Context().UserValue("collection").(*mongo.Collection)

	req := new(requests.Replace)
	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := collection.ReplaceOne(context.TODO(), req.Filter, req.Document, &req.ReplaceOptions)
	return c.SingleRes(res, &err)
}

func deleteOne(ctx *fiber.Ctx) error {
	c := &middleware.Context{Ctx: ctx}
	collection := c.Context().UserValue("collection").(*mongo.Collection)

	req := new(requests.Delete)
	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := collection.DeleteOne(context.TODO(), req.Filter, &req.DeleteOptions)
	return c.SingleRes(res, &err)
}

func aggregate(ctx *fiber.Ctx) error {
	c := &middleware.Context{Ctx: ctx}
	collection := c.Context().UserValue("collection").(*mongo.Collection)

	req := new(requests.Aggregate)
	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := collection.Aggregate(context.TODO(), req.Pipeline, &req.AggregateOptions)
	return c.SingleRes(res, &err)
}
