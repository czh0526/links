package ui

import (
	"fmt"
	"github.com/czh0526/links/account"
	"github.com/czh0526/links/global"
	"github.com/czh0526/links/links"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"sort"
	"strings"
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
	friends, err := fetchOrderedFriends(app.my)
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
	input.SetBorder(true)

	chatPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(msgBox, 0, 2, false).
		AddItem(input, 0, 1, true)

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(friendList, 35, 1, false).
		AddItem(chatPanel, 0, 1, false)
	root.SetRoot(flex, true)

	return &AppUI{
		my:         my,
		root:       root,
		friendList: friendList,
		doneCh:     make(chan struct{}),
	}
}

func fetchOrderedFriends(my *account.Account) ([]*links.Friend, error) {
	friends, err := global.GetMyFriends(my.Id)
	if err != nil {
		return nil, err
	}

	friendList := make([]*links.Friend, 0, len(friends))
	for _, friend := range friends {
		friendList = append(friendList, friend)
	}

	// 按照 nickname 排序
	sort.Slice(friendList, func(i, j int) bool {
		return strings.Compare(friendList[i].Nickname, friendList[j].Nickname) > 0
	})

	return friendList, nil
}
