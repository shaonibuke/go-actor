package player

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/shaonibuke/go-actor/actor"
	"github.com/shaonibuke/go-actor/actor/mail"
	"github.com/shaonibuke/go-actor/services/bag"
)

type Player struct {
	actor *actor.Actor
	Name  string // 玩家名字
	count int
}

func (p *Player) AddItem(item bag.Item) {
	p.actor.SendMessage("bag", "", "addItem", item)
}

func (p *Player) GetItem2(id string) {
	item := bag.Item{ID: id}
	p.actor.CallMessage("bag", "", "getItem", item, func(m *mail.Mail) {
		time.Sleep(time.Second * 1)
		if m.Msg.(bag.Item).ID == "" {
			// 没找到
			log.Debug("not find items", id)
		} else {
			// 找到了
			log.Debug("find items", m.Msg, id)
		}
	})
}

func (p *Player) GetItem(id string) {
	item := bag.Item{ID: id}
	p.actor.CallMessage("bag", "", "getItem", item, p.backGetItem)
}

func (p *Player) GetSyncItem(id string) {
	item := bag.Item{ID: id}
	ret := p.actor.SyncCallMessage("bag", "", "getItem", item)
	if ret.(bag.Item).ID == "" {
		// 没找到
		log.Debug("not find items...", id)
	} else {
		// 找到了
		log.Debugf("find items %v %s", ret.(bag.Item), id)
	}
}
