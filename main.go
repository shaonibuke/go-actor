package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/shaonibuke/go-actor/actor"
	"github.com/shaonibuke/go-actor/services/bag"
	"github.com/shaonibuke/go-actor/services/player"
)

func wait(serverName string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	s := <-c
	log.Debugf("server [%s] exit ------- signal:[%v]", serverName, s)
}

func main() {
	log.Debug("start	............")
	player.GetPlayer().GetItem("1")
	player.GetPlayer().GetItem("2")
	//player.GetPlayer().AddItem(bag.Item{ID: "1", Price: 1})
	go func() {
		for i := 0; i < 100; i++ {
			player.GetPlayer().AddItem(bag.Item{ID: fmt.Sprintf("%d", i), Price: i})
		}
	}()
	go func() {
		for i := 100; i < 200; i++ {
			player.GetPlayer().AddItem(bag.Item{ID: fmt.Sprintf("%d", i), Price: i})
		}
	}()

	player.GetPlayer().GetItem("1")
	time.Sleep(time.Second * 2)
	player.GetPlayer().GetItem("199")
	player.GetPlayer().GetSyncItem("2")
	time.Sleep(time.Second * 6)
	player.GetPlayer().GetSyncItem("1")

	actor.GetActorManager().ReportActorAvgCost()

	log.Debug("over	............")

	wait("")
}
