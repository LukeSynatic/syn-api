package requests

import "go.mongodb.org/mongo-driver/mongo/options"

type MongoReq struct {
	DataSource string `json:"dataSource"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

type Find struct {
	MongoReq
	options.FindOptions
	Filter interface{} `json:"filter"`
}

type FindOne struct {
	MongoReq
	options.FindOneOptions
	Filter interface{} `json:"filter"`
}

type InsertOne struct {
	MongoReq
	options.InsertOneOptions
	Document interface{} `json:"document"`
}

type InsertMany struct {
	MongoReq
	options.InsertManyOptions
	Documents []interface{} `json:"documents"`
}

type Update struct {
	MongoReq
	options.UpdateOptions
	Filter   interface{} `json:"filter"`
	Document interface{} `json:"document"`
}

type Delete struct {
	MongoReq
	options.DeleteOptions
	Filter interface{} `json:"filter"`
}

type Replace struct {
	MongoReq
	options.ReplaceOptions
	Filter   interface{} `json:"filter"`
	Document interface{} `json:"document"`
}

type Aggregate struct {
	MongoReq
	options.AggregateOptions
	Pipeline interface{} `json:"pipeline"`
}
