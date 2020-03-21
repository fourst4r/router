package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fourst4r/router"
)

func main() {
	r := router.New()
	r.On("ping", func(ctx *router.Context) {
		ctx.Client.(*twitchClient).Say("gofur", "pong!")
	})

	tc := new(twitchClient)
	go tc.Start(r)
	defer tc.Stop()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	fmt.Println("shutting down...")
}
