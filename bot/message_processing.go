package bot

import (
	"fmt"
	"log"
	"qbot/bot/at_member"
	"qbot/bot/common/tools"
	"qbot/bot/weather"
	"qbot/pkg/utils"
)

const (
	cqMe  = "[CQ:at,qq=%s] "
	AtAll = "/艾特全体"
)

var (
	AtMe = fmt.Sprintf(cqMe, utils.GlobalConf.QqBot.QqId)
)

func MessageDistribution(messageBytes []byte) {

	message, err := tools.ParsingMessage(messageBytes)
	if err != nil {
		log.Println("parse message failed,message = ", string(messageBytes))
		return
	}

	//处理消息事件
	if message.PostType == "message" {
		//@me 功能列表
		if message.Message == AtMe {
			go at_member.AtMeFunc(message)
		}

		if message.Message == AtAll {
			go at_member.AtAllMember(message)
		}

		//天气查询
		go weather.WeatherQueryFunc(message)
	}

}
