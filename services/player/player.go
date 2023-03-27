package player

import (
	"fmt"
	"main/actor"
	"main/actor/mail"
	"main/services/bag"
	"time"

	"github.com/gookit/goutil/dump"
)

type Player struct {
	actor *actor.Actor
	Name  string // 玩家名字
	count int
}

func (p *Player) GetPlayerName(m *mail.Mail) {
	//p.actor.SendMessage("bag", "", "GetPlayerName", p.Name)

	p.count += 1
	fmt.Printf("GetPlayerName %s %d\n", m.ReplyID, p.count)
	p.actor.ReplyMessage(m.FormServiceType, m.FormID, m.ReplyID, p.Name)
}

func (p *Player) AddItem(item bag.Item) {
	p.actor.SendMessage("bag", "", "AddItem", item)
}

func (p *Player) GetItem2(id string) {
	item := bag.Item{ID: id}
	p.actor.CallMessage("bag", "", "GetItem", item, func(m *mail.Mail) {
		time.Sleep(time.Second * 1)
		if m.Msg.(bag.Item).ID == "" {
			// 没找到
			dump.P("not find items", id)
		} else {
			// 找到了
			dump.P("find items", m.Msg, id)
		}
	})
}

func (p *Player) GetItem(id string) {
	item := bag.Item{ID: id}
	p.actor.CallMessage("bag", "", "GetItem", item, p.backGetItem)
}
