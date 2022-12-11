package lark

import (
	"encoding/json"
	"fmt"
	"github.com/chyroc/lark"
	"net/http"
	"problems-everyday-bot/model/request"
	"problems-everyday-bot/model/response"
	"strings"
)

// BuildEveryDayQuestionNotice 构建每日一题通知
func BuildEveryDayQuestionNotice(url string, q *response.GetQuestionOfTodayDetailResp) *lark.MessageContentCard {
	return &lark.MessageContentCard{
		Header: &lark.MessageContentCardHeader{
			Title: &lark.MessageContentCardObjectText{
				Tag:     "plain_text",
				Content: " 📅 LeetCode 每日一题",
			},
			Template: "green",
		},
		Config: &lark.MessageContentCardConfig{
			EnableForward: true,
		},
		I18NModules: &lark.MessageContentCardI18NModule{
			ZhCn: []lark.MessageContentCardModule{
				lark.MessageContentCardModuleDIV{
					Text: &lark.MessageContentCardObjectText{
						Tag:     "lark_md",
						Content: fmt.Sprintf(`**标题: %s**`, q.Data.Question.TranslatedTitle),
					},
				},
				lark.MessageContentCardModuleDIV{
					Text: &lark.MessageContentCardObjectText{
						Tag:     "lark_md",
						Content: fmt.Sprintf(`**难度: %s**`, q.Data.Question.Difficulty),
					},
				},
				lark.MessageContentCardModuleDIV{
					Text: &lark.MessageContentCardObjectText{
						Tag:     "lark_md",
						Content: fmt.Sprintf(`%s`, q.Data.Question.TranslatedContent),
					},
				},
				lark.MessageContentCardModuleDIV{
					Text: &lark.MessageContentCardObjectText{
						Tag:     "lark_md",
						Content: fmt.Sprintf(`**原题链接: %s**`, url),
					},
				},
			},
		},
	}
}

func SendBotMessage(u string, c *lark.MessageContentCard) error {
	req := request.SendBotMessageReq{
		MsgType: "interactive",
		Card:    c,
	}
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.Post(u, "application/json", strings.NewReader(string(payload)))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("发送消息失败, 消息ID: %s", resp.Status)
	}

	return nil // 发送成功
}

//func SendMessage(o string, c *lark.MessageContentCard) error {
//
//	s, err := json.Marshal(&c)
//	fmt.Println(string(s))
//	rawResp, resp, err := lark.New(lark.WithAppCredential("", "")).Message.Send().ToOpenID(o).SendCard(context.TODO(), c.String())
//
//	if err != nil {
//		return err
//	}
//	if &resp.StatusCode != nil && resp.StatusCode != 0 {
//		return fmt.Errorf("发送消息失败, 消息ID: %s", rawResp.MessageID)
//	}
//	return nil // 发送成功
//}
