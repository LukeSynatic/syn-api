package requests

import "go.mongodb.org/mongo-driver/mongo/options"

type FindReq struct {
	MongoReq
	Filter     interface{} `json:"filter"`
	Projection interface{} `json:"projection"`
	Sort       interface{} `json:"sort"`
	Limit      int64       `json:"limit"`
	Skip       int64       `json:"skip"`
}

func (req *FindReq) Opts() *options.FindOptions {
	opts := new(options.FindOptions)

	if req.Projection != nil {
		opts.Projection = req.Projection
	}

	if req.Sort != nil {
		opts.Sort = req.Sort
	}

	if req.Limit > 0 {
		opts.Sort = req.Limit
	}

	if req.Skip > 0 {
		opts.Sort = req.Skip
	}

	return opts
}
