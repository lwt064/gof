package state

import "fmt"

type AccountState interface {
	View()    // 观看
	Comment() // 评论
	Post()    // 发贴
}

type Account struct {
	State       AccountState
	healthValue int
}

func (a *Account) View() {
	a.State.View()
}

func (a *Account) Comment() {
	a.State.Comment()
}

func (a *Account) Post() {
	a.State.Post()
}

func (a *Account) SetHealthValue(score int) {
	a.healthValue = score
	a.ChangeState()
}

func (a *Account) ChangeState() {
	if a.healthValue < 0 {
		a.State = &ClosedState{}
	} else if a.healthValue < 60 {
		a.State = &RestrictState{}
	} else if a.healthValue < 100 {
		a.State = &NormalState{}
	}
}

type NormalState struct{}

func (ns *NormalState) View() {
	fmt.Println("观看正常")
}

func (ns *NormalState) Comment() {
	fmt.Println("评论正常")
}

func (ns *NormalState) Post() {
	fmt.Println("发贴正常")
}

type RestrictState struct{}

func (rs *RestrictState) View() {
	fmt.Println("观看正常")
}

func (rs *RestrictState) Comment() {
	fmt.Println("评论正常")
}

func (rs *RestrictState) Post() {
	fmt.Println("无法发贴")
}

type ClosedState struct{}

func (cs *ClosedState) View() {
	fmt.Println("观看正常")
}

func (cs *ClosedState) Comment() {
	fmt.Println("无法评论")
}

func (cs *ClosedState) Post() {
	fmt.Println("无法发帖")
}
