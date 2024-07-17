package wechat

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

var (
	bot  *openwechat.Bot
	Self *openwechat.Self
)

func Login() string {
	if Self != nil {
		return "已经登陆了"
	}
	bot = openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return "登陆失败....请联系lixun "
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return "登陆失败....请联系lixun "
	}
	Self = self
	fmt.Println("Login successful")
	bot.Block()
	return "Login successful"
}

func GetAllFriends() []string {
	var str []string

	if Self == nil {
		return append(str, "先登陆咯")
	}
	friends, err := Self.Friends()
	if err != nil {
		fmt.Println("record", err)
	}
	for _, r := range friends {
		s := fmt.Sprintf("Name: %s, Remark: %s, Nickname: %s \n", r.NickName, r.RemarkName, r.NickName)
		str = append(str, s)
	}
	return str
}

func GetAllGroups() string {
	if Self == nil {
		return "先登陆咯"
	}
	groups, err := Self.Groups()
	fmt.Println(groups, err)
	result := fmt.Sprintf("列表:", groups)
	return result
}
