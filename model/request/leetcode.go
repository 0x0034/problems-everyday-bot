package request

type GetQuestionOfTodayReq struct {
	OperationName string `json:"operationName"`
	Variables     struct {
	} `json:"variables"`
	Query string `json:"query"`
}

type GetQuestionOfTodayDetailReq struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		TitleSlug string `json:"titleSlug"`
	} `json:"variables"`
	Query string `json:"query"`
}
