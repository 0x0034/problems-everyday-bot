package request

import "github.com/chyroc/lark"

type SendBotMessageReq struct {
	MsgType string                   `json:"msg_type"`
	Card    *lark.MessageContentCard `json:"card"`
}
