package response

type GetQuestionOfTodayResp struct {
	Data struct {
		TodayRecord []struct {
			Question struct {
				QuestionFrontendId string `json:"questionFrontendId"`
				QuestionTitleSlug  string `json:"questionTitleSlug"`
				Typename           string `json:"__typename"`
			} `json:"question"`
			LastSubmission interface{} `json:"lastSubmission"`
			Date           string      `json:"date"`
			UserStatus     interface{} `json:"userStatus"`
			Typename       string      `json:"__typename"`
		} `json:"todayRecord"`
	} `json:"data"`
}

type GetQuestionOfTodayDetailResp struct {
	Data struct {
		Question struct {
			QuestionId            string        `json:"questionId"`
			QuestionFrontendId    string        `json:"questionFrontendId"`
			BoundTopicId          int           `json:"boundTopicId"`
			Title                 string        `json:"title"`
			TitleSlug             string        `json:"titleSlug"`
			Content               string        `json:"content"`
			TranslatedTitle       string        `json:"translatedTitle"`
			TranslatedContent     string        `json:"translatedContent"`
			IsPaidOnly            bool          `json:"isPaidOnly"`
			Difficulty            string        `json:"difficulty"`
			Likes                 int           `json:"likes"`
			Dislikes              int           `json:"dislikes"`
			IsLiked               interface{}   `json:"isLiked"`
			SimilarQuestions      string        `json:"similarQuestions"`
			Contributors          []interface{} `json:"contributors"`
			LangToValidPlayground string        `json:"langToValidPlayground"`
			TopicTags             []struct {
				Name           string `json:"name"`
				Slug           string `json:"slug"`
				TranslatedName string `json:"translatedName"`
				Typename       string `json:"__typename"`
			} `json:"topicTags"`
			CompanyTagStats interface{} `json:"companyTagStats"`
			CodeSnippets    []struct {
				Lang     string `json:"lang"`
				LangSlug string `json:"langSlug"`
				Code     string `json:"code"`
				Typename string `json:"__typename"`
			} `json:"codeSnippets"`
			Stats             string        `json:"stats"`
			Hints             []string      `json:"hints"`
			Solution          interface{}   `json:"solution"`
			Status            interface{}   `json:"status"`
			SampleTestCase    string        `json:"sampleTestCase"`
			MetaData          string        `json:"metaData"`
			JudgerAvailable   bool          `json:"judgerAvailable"`
			JudgeType         string        `json:"judgeType"`
			MysqlSchemas      []interface{} `json:"mysqlSchemas"`
			EnableRunCode     bool          `json:"enableRunCode"`
			EnvInfo           string        `json:"envInfo"`
			Book              interface{}   `json:"book"`
			IsSubscribed      bool          `json:"isSubscribed"`
			IsDailyQuestion   bool          `json:"isDailyQuestion"`
			DailyRecordStatus interface{}   `json:"dailyRecordStatus"`
			EditorType        string        `json:"editorType"`
			UgcQuestionId     interface{}   `json:"ugcQuestionId"`
			Style             string        `json:"style"`
			Typename          string        `json:"__typename"`
		} `json:"question"`
	} `json:"data"`
}
