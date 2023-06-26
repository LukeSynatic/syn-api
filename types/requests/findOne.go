package requests

type FindOneReq struct {
	MongoReq
	Filter     interface{} `json:"filter"`
	Projection interface{} `json:"projection"`
}

// func (req *FindReq) Opts() *options.FindOneOptions {
// 	opts := new(options.FindOneOptions)

// 	if req.Projection != nil {
// 		opts.Projection = req.Projection
// 	}

// 	if req.Sort != nil {
// 		opts.Sort = req.Sort
// 	}

// 	if req.Limit > 0 {
// 		opts.Sort = req.Limit
// 	}

// 	if req.Skip > 0 {
// 		opts.Sort = req.Skip
// 	}

// 	return opts
// }
