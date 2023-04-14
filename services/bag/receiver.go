package bag

import (
	"github.com/charmbracelet/log"
	"github.com/shaonibuke/go-actor/actor/mail"
)

func (b *Bag) addItem(m *mail.Mail) {
	item := m.Msg.(Item)
	b.items = append(b.items, item)
	b.actor.CallMessage(m.FormServiceType, m.FormID, "getPlayerName", nil, func(msg *mail.Mail) {
		name := msg.Msg.(string)
		log.Debugf("name: :%s", name)
	})
}

func (b *Bag) getItem(m *mail.Mail) {
	item := m.Msg.(Item)
	for _, v := range b.items {
		if v.ID == item.ID {
			b.actor.ReplyMessage(m.FormServiceType, m.FormID, m.ReplyID, v)
			return
		}
	}
	b.actor.ReplyMessage(m.FormServiceType, m.FormID, m.ReplyID, Item{})
}
