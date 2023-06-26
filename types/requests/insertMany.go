package requests

type InsertManyReq struct {
	MongoReq
	Documents                []interface{} `json:"documents"`
	BypassDocumentValidation bool          `json:"bypassDocumentValidation"`
	Ordered                  bool          `json:"ordered"`
}
