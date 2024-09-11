package ui

import (
	"fmt"
	"github.com/czh0526/links/account"
	"github.com/czh0526/links/global"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"time"
)

type AppUI struct {
	my         *account.Account
	root       *tview.Application
	friendList *tview.TextView
	doneCh     chan struct{}
}

func (app *AppUI) Run() error {
	go app.handleEvents()
	defer app.end()

	return app.root.Run()
}

func (app *AppUI) end() {
	app.doneCh <- struct{}{}
}

func (app *AppUI) handleEvents() {
	friendsRefreshTicker := time.NewTicker(time.Second)
	for {
		select {
		case <-friendsRefreshTicker.C:
			{
				app.refreshLinks()
			}
		case <-app.doneCh:
			{
				log.Println("【ui】: ui done")
				return
			}
		}
	}
}

func (app *AppUI) refreshLinks() {
	friends, err := global.GetMyFriends(app.my.Id)
	if err != nil {
		log.Printf("【ui】: 刷新用户列表出错：err = %v", err)
		return
	}

	app.friendList.Clear()
	for _, friend := range friends {
		_, _ = fmt.Fprintln(app.friendList, friend.Nickname)
	}
	app.root.Draw()
}

func NewAppUI(my *account.Account) *AppUI {
	root := tview.NewApplication()

	msgBox := tview.NewTextView()
	msgBox.SetDynamicColors(true)
	msgBox.SetBorder(true)
	msgBox.SetTitle(fmt.Sprintf("wait for setting"))

	friendList := tview.NewTextView()
	friendList.SetBorder(true)
	friendList.SetTitle("Peers")
	friendList.SetChangedFunc(func() { root.Draw() })

	input := tview.NewInputField().
		SetLabel(my.Nickname + " > ").
		SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.ColorBlack)

	chatPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(msgBox, 0, 1, false).
		AddItem(input, 0, 1, true)

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(friendList, 35, 1, false).
		AddItem(chatPanel, 1, 1, false)
	root.SetRoot(flex, true)

	return &AppUI{
		my:         my,
		root:       root,
		friendList: friendList,
		doneCh:     make(chan struct{}),
	}
}
