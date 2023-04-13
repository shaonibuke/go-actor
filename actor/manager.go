package actor

import (
	"github.com/charmbracelet/log"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/shaonibuke/go-actor/actor/base"
	"github.com/shaonibuke/go-actor/actor/mail"
)

const (
	MAX_MAIL_COUNT = 1024 * 1024
)

// ActorManager 管理Actor
type ActorManager struct {
	// Actors Actor的集合
	Actors map[string]map[string]*Actor // ActorType -> ActorID -> Actor
}

func (am *ActorManager) NewActor(serviceType string) *Actor {
	ac := &Actor{
		ActorID:      base.GetID(),
		ActorType:    serviceType,
		StopChan:     make(chan bool),
		MailBox:      make(chan *mail.Mail, MAX_MAIL_COUNT),
		router:       make(map[string]interface{}),
		actorManager: am,
		callBacks:    cmap.New(),
	}
	ac.Run()

	if _, ok := am.Actors[serviceType]; !ok {
		am.Actors[serviceType] = make(map[string]*Actor)
	}
	am.Actors[serviceType][ac.ActorID] = ac

	return ac
}

// Stop 停止ActorManager 停止需要仔细处理
func (am *ActorManager) Stop() {
	for _, ac := range am.Actors {
		for _, a := range ac {
			a.Stop()
		}
	}
}

// SendMessage 发送消息
func (am *ActorManager) SendMessage(toServiceType, toServerID, msgName string, msg interface{}) {

	toAc := am.GetActor(toServiceType, toServerID)
	if toAc == nil {
		log.Errorf("ActorManager.SendMessage toAc is nil %s %s %s", toServiceType, toServerID, msgName)
		return
	}
	toAc.SendMessage(toServiceType, toServerID, msgName, msg)
}

// GetActor 获取Actor
func (am *ActorManager) GetActor(serviceType, serverID string) *Actor {
	if _, ok := am.Actors[serviceType]; !ok {
		log.Errorf("ActorManager.GetActor: serviceType not found %s\n", serviceType)
		return nil
	}
	if serverID == "" {
		return am.GetOneActorByType(serviceType)
	}
	if _, ok := am.Actors[serviceType][serverID]; !ok {
		log.Errorf("ActorManager.GetActor: serverID not found %s %s\n", serviceType, serverID)
		return nil
	}

	return am.Actors[serviceType][serverID]
}

// GetActorByType 获取Actor
func (am *ActorManager) GetActorByType(serviceType string) map[string]*Actor {
	if _, ok := am.Actors[serviceType]; !ok {
		log.Errorf("ActorManager.GetActor: serviceType not found\n")
		return nil
	}
	return am.Actors[serviceType]
}

func (am *ActorManager) GetOneActorByType(serviceType string) *Actor {
	if _, ok := am.Actors[serviceType]; !ok {
		log.Errorf("ActorManager.GetActor: serviceType not found\n")
		return nil
	}
	for _, a := range am.Actors[serviceType] {
		return a
	}
	return nil
}

// GetActorCount 获取Actor数量
func (am *ActorManager) GetActorCount() int {
	count := 0
	for _, ac := range am.Actors {
		count += len(ac)
	}
	return count
}

// newActorManager 创建一个ActorManager
func newActorManager() *ActorManager {
	// ...
	return &ActorManager{
		Actors: make(map[string]map[string]*Actor),
	}
}
