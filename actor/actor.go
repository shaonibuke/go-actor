package actor

import (
	"fmt"
	"main/actor/base"
	"main/actor/mail"
)

// toServerID 为 "" 时，表示发送给所有的Actor
type CallFn func(msg *mail.Mail)

type Actor struct {
	// ActorID Actor的ID
	ActorID string
	// ActorType Actor的类型
	ActorType string
	// MailBox Actor的邮箱
	MailBox chan *mail.Mail
	// StopChan Actor的停止信号
	StopChan chan bool
	// router Actor的路由
	router map[string]interface{}
	// actorManager Actor所属的ActorManager
	actorManager *ActorManager
	// callBacks Actor的回调函数
	callBacks map[string]CallFn
	// nowProcessMail 正在处理的Mail
	nowProcessMail *mail.Mail
	// nowReplyID 正在回复的ID ""表示已经不需要回复了
	nowReplyID string
}

// Run 运行Actor
func (a *Actor) Run() {
	go func() {
		for {
			select {
			case msg := <-a.MailBox:
				// dump.P(msg)
				a.nowProcessMail = msg
				a.processMessage(msg)
				a.checkIsReply()
			case <-a.StopChan:
				return
			}
		}
	}()
}

// Stop 停止Actor
func (a *Actor) Stop() {
	a.StopChan <- true
}

// SendMessage 发送消息
func (a *Actor) SendMessage(toServiceType, toServerID, msgName string, msg interface{}) {
	toAc := a.actorManager.GetActor(toServiceType, toServerID)
	if toAc == nil {
		fmt.Printf("Actor.SendMessage: toAc is nil\n")
		return
	}
	select {
	case toAc.MailBox <- &mail.Mail{
		Msg:             msg,
		MsgName:         msgName,
		MsgType:         mail.MsgTypeTo,
		FormID:          a.ActorID,
		FormServiceType: a.ActorType,
		ToID:            toAc.ActorID,
		ToServiceType:   toServiceType,
	}:
	default:
		fmt.Printf("Actor.SendMessage: toAc.MailBox is full\n")
	}
}

// CallMessage 调用消息
func (a *Actor) CallMessage(toServiceType, toServerID, msgName string, msg interface{}, callback CallFn) {
	toAc := a.actorManager.GetActor(toServiceType, toServerID)
	if toAc == nil {
		fmt.Printf("Actor.CallMessage: toAc is nil\n")
		return
	}
	replyID := base.GetID()
	select {
	case toAc.MailBox <- &mail.Mail{
		Msg:             msg,
		MsgName:         msgName,
		MsgType:         mail.MsgTypeTo,
		FormID:          a.ActorID,
		FormServiceType: a.ActorType,
		ToID:            toAc.ActorID,
		ToServiceType:   toServiceType,
		ReplyID:         replyID,
	}:
	default:
		fmt.Printf("Actor.CallMessage: toAc.MailBox is full\n")
	}
	a.callBacks[replyID] = callback
}

// ReplyMessage 回复消息
func (a *Actor) ReplyMessage(toServiceType, toServerID, replyID string, msg interface{}) {
	if replyID == "" {
		fmt.Print("Actor.ReplyMessage: replyID is nil\n")
		return
	}
	if a.nowReplyID != replyID {
		fmt.Print("Actor.ReplyMessage: a.NowReplyID != replyID\n")
		return
	}
	a.nowReplyID = ""
	//fmt.Printf("ToServiceType:%s, toServerID:%s, replyID:%s, msg:%v\n", ToServiceType, toServerID, replyID, msg)
	toAc := a.actorManager.GetActor(toServiceType, toServerID)
	if toAc == nil {
		fmt.Print("Actor.ReplyMessage: toAc is nil\n")
		return
	}
	select {
	case toAc.MailBox <- &mail.Mail{
		Msg:             msg,
		MsgType:         mail.MsgTypeReply,
		FormID:          a.ActorID,
		FormServiceType: a.ActorType,
		ToID:            toAc.ActorID,
		ToServiceType:   toServiceType,
		ReplyID:         replyID,
	}:
	default:
		fmt.Print("Actor.ReplyMessage: toAc.MailBox is full\n")
	}
}

// RegisterRouter 注册Actor的路由
func (a *Actor) RegisterRouter(msgName string, msgHandler CallFn) {
	a.router[msgName] = msgHandler
}

// processMessage 处理消息
func (a *Actor) processMessage(m *mail.Mail) {

	if m.MsgType == mail.MsgTypeTo {
		if m.ToID != a.ActorID {
			fmt.Print("Actor.ProcessMessage: m.ToID != a.ActorID\n")
			return
		}
		if m.ReplyID != "" {
			a.nowReplyID = m.ReplyID
		}
		if m.MsgName == "" {
			fmt.Print("Actor.ProcessMessage: m.MsgName is nil")
			return
		}
		if msgHandler, ok := a.router[m.MsgName]; ok {
			msgHandler.(CallFn)(m)
		} else {
			panic(fmt.Sprintf("Actor.ProcessMessage: msgHandler is nil, m.MsgName:%s", m.MsgName))
		}
	} else if m.MsgType == mail.MsgTypeReply {
		if m.ReplyID == "" {
			fmt.Print("Actor.ProcessMessage: m.ReplyID is nil\n")
			return
		}
		if callback, ok := a.callBacks[m.ReplyID]; ok {
			callback(m)
			delete(a.callBacks, m.ReplyID)
		} else {
			fmt.Printf("Actor.ProcessMessage: callback is nil, msg:%v\n", m)
			panic("Actor.ProcessMessage: callback is nil")
		}
	}
}

// CheckIsReply 检查是否是回复消息
func (a *Actor) checkIsReply() {
	if a.nowReplyID == "" {
		return
	}
	if a.nowReplyID != a.nowProcessMail.ReplyID {
		return
	}
	if a.nowReplyID != "" {
		panic(fmt.Sprintf(" not reply :%s %s", a.nowProcessMail.MsgName, a.nowReplyID))
		//a.nowReplyID = ""
		//a.ReplyMessage(a.nowProcessMail.ToServiceType, a.nowProcessMail.FormID, a.nowProcessMail.ReplyID, nil)
	}
}
