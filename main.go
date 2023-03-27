package main

import (
	"fmt"
	"main/services/bag"
	"main/services/player"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func wait(serverName string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	s := <-c
	fmt.Printf("server [%s] exit ------- signal:[%v]", serverName, s)
}

func main() {
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
	player.GetPlayer().GetItem("100")
	player.GetPlayer().GetItem("99")
	player.GetPlayer().GetItem("200")
	time.Sleep(time.Second * 2)
	player.GetPlayer().GetItem("199")

	print("over	............")

	wait("")
}