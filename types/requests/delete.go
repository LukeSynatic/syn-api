package requests

type DeleteReq struct {
	MongoReq
	Filter interface{} `json:"filter"`
}
