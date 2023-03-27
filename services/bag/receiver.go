package bag

import (
	"fmt"
	"main/actor/mail"
)

func (b *Bag) AddItem(m *mail.Mail) {
	item := m.Msg.(Item)
	b.items = append(b.items, item)
	b.actor.CallMessage(m.FormServiceType, m.FormID, "GetPlayerName", nil, func(msg *mail.Mail) {
		name := msg.Msg.(string)
		fmt.Printf("name: :%s\n", name)
	})
}

func (b *Bag) GetItem(m *mail.Mail) {
	item := m.Msg.(Item)
	for _, v := range b.items {
		if v.ID == item.ID {
			b.actor.ReplyMessage(m.FormServiceType, m.FormID, m.ReplyID, v)
			return
		}
	}
	b.actor.ReplyMessage(m.FormServiceType, m.FormID, m.ReplyID, Item{})
}
