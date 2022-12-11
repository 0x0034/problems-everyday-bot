package leetcode

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"problems-everyday-bot/model/request"
	"problems-everyday-bot/model/response"
	"strings"
)

const LeetCodeURL = "https://leetcode-cn.com"
const GetTodayQuestionGraphQL = "query questionOfToday { todayRecord {   question {     questionFrontendId     questionTitleSlug     __typename   }   lastSubmission {     id     __typename   }   date   userStatus   __typename }}"
const GetTodayQuestionDetailGraphQL = "query questionData($titleSlug: String!) {  question(titleSlug: $titleSlug) {    questionId    questionFrontendId    boundTopicId    title    titleSlug    content    translatedTitle    translatedContent    isPaidOnly    difficulty    likes    dislikes    isLiked    similarQuestions    contributors {      username      profileUrl      avatarUrl      __typename    }    langToValidPlayground    topicTags {      name      slug      translatedName      __typename    }    companyTagStats    codeSnippets {      lang      langSlug      code      __typename    }    stats    hints    solution {      id      canSeeDetail      __typename    }    status    sampleTestCase    metaData    judgerAvailable    judgeType    mysqlSchemas    enableRunCode    envInfo    book {      id      bookName      pressName      source      shortDescription      fullDescription      bookImgUrl      pressImgUrl      productUrl      __typename    }    isSubscribed    isDailyQuestion    dailyRecordStatus    editorType    ugcQuestionId    style    __typename  }}"

// GetTodayQuestion 获取今日一题
func GetTodayQuestion() (*response.GetQuestionOfTodayResp, error) {
	payload, err := json.Marshal(request.GetQuestionOfTodayReq{
		OperationName: "questionOfToday",
		Variables:     struct{}{},
		Query:         GetTodayQuestionGraphQL,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.Post(LeetCodeURL+"/graphql", "application/json", strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	resp := &response.GetQuestionOfTodayResp{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// GetTodayQuestionDetail 获取今日一题详情
func GetTodayQuestionDetail() (*response.GetQuestionOfTodayDetailResp, error) {
	q, err := GetTodayQuestion()
	if err != nil {
		return nil, err
	}
	if len(q.Data.TodayRecord) <= 0 {
		return nil, errors.New("get today question failed")
	}
	payload, err := json.Marshal(request.GetQuestionOfTodayDetailReq{
		OperationName: "questionData",
		Variables: struct {
			TitleSlug string `json:"titleSlug"`
		}{
			TitleSlug: q.Data.TodayRecord[0].Question.QuestionTitleSlug,
		},
		Query: GetTodayQuestionDetailGraphQL,
	})
	req, err := http.Post(LeetCodeURL+"/graphql", "application/json", strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	resp := &response.GetQuestionOfTodayDetailResp{}
	err = json.Unmarshal(body, resp)
	return resp, err
}
