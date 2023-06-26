package requests

type AggregateReq struct {
	MongoReq
	Pipeline                 interface{} `json:"pipeline"`
	BypassDocumentValidation bool        `json:"bypassDocumentValidation"`
	BatchSize                int32       `json:"batchSize"`
}
