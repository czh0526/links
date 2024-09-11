package main

import (
	"flag"
	"fmt"
	"github.com/czh0526/links/account"
	"github.com/czh0526/links/global"
	"github.com/czh0526/links/ui"
	"github.com/libp2p/go-libp2p"
	"log"
	"os"
)

var (
	nickFlag = flag.String("nick", "", "nickname to login")
	passFlag = flag.String("pass", "", "password to login")
)

func loadMyLinks(my *account.Account) error {

	_, err := global.GetMyFriends(my.Id)
	if err != nil {
		return err
	}

	_, err = global.GetMyGroups(my.Id)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	//ctx := context.Background()
	flag.Parse()

	nickname := *nickFlag
	passphrase := *passFlag

	// 查找账户
	myAccount, err := global.GetMyAccount(nickname)
	if err != nil {
		panic(fmt.Sprintf("查找账户出错：err = %v", err))
	}

	// 验证密码
	privateKey, err := global.GetPrivateKey(passphrase)
	if err != nil {
		panic(fmt.Sprintf("验证密码出错: err = %v", err))
	}

	// 加载联系人
	err = loadMyLinks(myAccount)
	if err != nil {
		panic(fmt.Sprintf("加载我的联系人出错: err = %v", err))
	}

	// 创建主机
	h, err := libp2p.New(
		libp2p.Identity(privateKey),
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(fmt.Sprintf("创建主机失败: err = %v", err))
	}

	// 初始化日志文件
	logFile, err := initLog(myAccount.Nickname, shortID(h))
	if err != nil {
	}
	defer logFile.Close()

	// 启动 UI 界面
	appUi := ui.NewAppUI(myAccount)
	if err = appUi.Run(); err != nil {
		panic(fmt.Sprintf("启动界面失败: err = %v", err))

	}
}

func initLog(nickname string, id string) (*os.File, error) {
	file, err := os.OpenFile(
		fmt.Sprintf("logs/%s_%s.log", nickname, id),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)

	return file, nil
}
