package actor

import (
	"fmt"

	"github.com/charmbracelet/log"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/shaonibuke/go-actor/actor/base"
	"github.com/shaonibuke/go-actor/actor/mail"
)

// toServerID 为 "" 时，表示发送给所有的Actor
type CallFn func(msg *mail.Mail)

type Actor struct {
	// ActorID Actor的ID
	ActorID string
	// ActorType Actor的类型
	ActorType string
	// MailBox Actor的邮箱
	mailBox chan *mail.Mail
	// StopChan Actor的停止信号
	stopChan chan bool
	// router Actor的路由
	router map[string]interface{}
	// actorManager Actor所属的ActorManager
	actorManager *ActorManager
	// callBacks Actor的回调函数
	callBacks cmap.ConcurrentMap
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
			case msg := <-a.mailBox:
				a.nowProcessMail = msg
				a.processMessage(msg)
				a.checkIsReply()
			case <-a.stopChan:
				return
			}
		}
	}()
}

// Stop 停止Actor
func (a *Actor) Stop() {
	a.stopChan <- true
}

// PushMailBox 压入邮箱
func (a *Actor) PushMailBox(msg *mail.Mail) {
	select {
	case a.mailBox <- msg:
	default:
		log.Errorf("Actor.PushMailBox: a.MailBox is full")
	}
}

// SendMessage 发送消息
func (a *Actor) SendMessage(toServiceType, toServerID, msgName string, msg interface{}) {
	toAc := a.actorManager.GetActor(toServiceType, toServerID)
	if toAc == nil {
		log.Errorf("Actor.SendMessage: toAc is nil")
		return
	}
	toAc.PushMailBox(&mail.Mail{
		Msg:             msg,
		MsgName:         msgName,
		MsgType:         mail.MsgTypeTo,
		FormID:          a.ActorID,
		FormServiceType: a.ActorType,
		ToID:            toAc.ActorID,
		ToServiceType:   toServiceType,
	})

}

// CallMessage 调用消息
func (a *Actor) CallMessage(toServiceType, toServerID, msgName string, msg interface{}, callback CallFn) {
	toAc := a.actorManager.GetActor(toServiceType, toServerID)
	if toAc == nil {
		log.Errorf("Actor.CallMessage: toAc is nil")
		return
	}
	replyID := base.GetID()

	toAc.PushMailBox(&mail.Mail{
		Msg:             msg,
		MsgName:         msgName,
		MsgType:         mail.MsgTypeTo,
		FormID:          a.ActorID,
		FormServiceType: a.ActorType,
		ToID:            toAc.ActorID,
		ToServiceType:   toServiceType,
		ReplyID:         replyID,
	})
	a.callBacks.Set(replyID, callback)
}

// ReplyMessage 回复消息
func (a *Actor) ReplyMessage(toServiceType, toServerID, replyID string, msg interface{}) {
	if replyID == "" {
		log.Errorf("Actor.ReplyMessage: replyID is nil")
		return
	}
	if a.nowReplyID != replyID {
		log.Errorf("Actor.ReplyMessage: a.NowReplyID != replyID")
		return
	}
	a.nowReplyID = ""
	toAc := a.actorManager.GetActor(toServiceType, toServerID)
	if toAc == nil {
		log.Errorf("Actor.ReplyMessage: toAc is nil")
		return
	}

	toAc.PushMailBox(&mail.Mail{
		Msg:             msg,
		MsgType:         mail.MsgTypeReply,
		FormID:          a.ActorID,
		FormServiceType: a.ActorType,
		ToID:            toAc.ActorID,
		ToServiceType:   toServiceType,
		ReplyID:         replyID,
	})

}

// RegisterRouter 注册Actor的路由
func (a *Actor) RegisterRouter(msgName string, msgHandler CallFn) {
	a.router[msgName] = msgHandler
}

// processMessage 处理消息
func (a *Actor) processMessage(m *mail.Mail) {

	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Actor.ProcessMessage: %v", err)
		}
	}()

	if m.MsgType == mail.MsgTypeTo {
		if m.ToID != a.ActorID {
			log.Errorf("Actor.ProcessMessage: m.ToID != a.ActorID")
			return
		}
		if m.ReplyID != "" {
			a.nowReplyID = m.ReplyID
		}
		if m.MsgName == "" {
			log.Errorf("Actor.ProcessMessage: m.MsgName is nil")
			return
		}
		if msgHandler, ok := a.router[m.MsgName]; ok {
			msgHandler.(CallFn)(m)
		} else {
			panic(fmt.Sprintf("Actor.ProcessMessage: msgHandler is nil, m.MsgName:%s", m.MsgName))
		}
	} else if m.MsgType == mail.MsgTypeReply {
		if m.ReplyID == "" {
			log.Errorf("Actor.ProcessMessage: m.ReplyID is nil")
			return
		}

		if callback, ok := a.callBacks.Get(m.ReplyID); ok {
			fn := callback.(CallFn)
			fn(m)
			a.callBacks.Remove(m.ReplyID)
		} else {
			panic(fmt.Sprintf("Actor.ProcessMessage: callback is nil, msg:%v", m))
		}
	}
}

// CheckIsReply 检查是否是回复消息
func (a *Actor) checkIsReply() {
	if a.nowReplyID != "" {
		panic(fmt.Sprintf(" not reply :%s %s", a.nowProcessMail.MsgName, a.nowReplyID))
	}
}
