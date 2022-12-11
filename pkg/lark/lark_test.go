package lark

import (
	"problems-everyday-bot/pkg/leetcode"
	"testing"
)

func TestSendMessage(t *testing.T) {
	questionDetail, err := leetcode.GetTodayQuestionDetail()
	if err != nil {
		t.Error(err)
	}
	url := leetcode.LeetCodeURL + "/problems/" + questionDetail.Data.Question.TitleSlug
	card := BuildEveryDayQuestionNotice(url, questionDetail)
	err = SendBotMessage("", card)
	if err != nil {
		t.Error(err)
	}
	//err = SendMessage("", card)
	//if err != nil {
	//	t.Error(err)
	//}

}
