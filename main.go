package main

import (
	"flag"
	"fmt"
	"os"
	"problems-everyday-bot/config"
	"problems-everyday-bot/pkg/lark"
	"problems-everyday-bot/pkg/leetcode"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "f", "", "配置文件")
	flag.Parse()
	config.InitConfig(configPath)

	questionDetail, err := leetcode.GetTodayQuestionDetail()
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(-1)
	}
	url := leetcode.LeetCodeURL + "/problems/" + questionDetail.Data.Question.TitleSlug
	card := lark.BuildEveryDayQuestionNotice(url, questionDetail)
	// 替换此处的bot url
	for _, bot := range *config.Conf.LarkBots {
		err = lark.SendBotMessage(bot, card)
		if err != nil {
			fmt.Printf("err: %v", err)
			os.Exit(-1)
		}
		fmt.Printf("机器人: %s 发送成功", bot)
	}
	fmt.Printf("全部机器人推送成功")
	os.Exit(0)
}
