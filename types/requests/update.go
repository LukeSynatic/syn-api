package requests

import "go.mongodb.org/mongo-driver/mongo/options"

type UpdateReq struct {
	MongoReq
	Filter                   interface{}          `json:"filter"`
	Document                 interface{}          `json:"document"`
	BypassDocumentValidation bool                 `json:"bypassDocumentValidation"`
	Upsert                   bool                 `json:"upsert"`
	ArrayFilters             options.ArrayFilters `json:"arrayFilters"`
}
