package requests

type InsertOneReq struct {
	MongoReq
	Document                 interface{} `json:"document"`
	BypassDocumentValidation bool        `json:"bypassDocumentValidation"`
}
