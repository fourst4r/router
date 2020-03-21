package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/fourst4r/router"
	"github.com/gempir/go-twitch-irc"
)

func parseCommand(prefix, text string) (command string, args []string, err error) {
	if !strings.HasPrefix(text, prefix) {
		return "", nil, errors.New("no prefix")
	}

	fields := strings.Fields(text)
	if len(fields) == 0 {
		return "", nil, errors.New("empty message")
	}
	fmt.Println(fields)

	command = strings.TrimPrefix(fields[0], prefix)
	if len(fields) > 1 {
		args = fields[1:]
	} else {
		args = make([]string, 0)
	}

	return
}

type twitchClient struct {
	*twitch.Client
	// broker   Broker
	myBadges map[string]string
	lastResp string
	// banphr   *urlrgxBanphraser
}

func (tc *twitchClient) ID() string { return "TWITCH" }

func (tc *twitchClient) isMod(channel string) bool {
	// lazy but effective ;)
	return strings.Contains(tc.myBadges[channel], "moderator")
}

func (tc *twitchClient) isVIP(channel string) bool {
	return strings.Contains(tc.myBadges[channel], "vip")
}

func (tc *twitchClient) Start(r *router.Route) {
	tc.Client = twitch.NewClient("gofir", "oauth:4v3tfl1l92q4mufp28u62qb4um9kn5")

	tc.OnNewMessage(func(channel string, user twitch.User, msg twitch.Message) {
		cmd, args, err := parseCommand(",", msg.Text)
		if err != nil {
			return
		}
		// args = append([]string{strings.TrimPrefix(cmd, ",")}, args)
		const separator = " "
		fmt.Println("cmd=", cmd)
		rt := r.Find(cmd)
		if rt != nil {
			rt.Handler(&router.Context{
				Route:   rt,
				Client:  tc,
				Author:  &twitchAuthor{&user},
				Msg:     &twitchMessage{&msg},
				Command: cmd,
				Args:    args,
			})
		}
		// if rt, depth := r.FindFull(args...); depth > 0 {
		// 	args = append([]string{strings.Join(args[:depth], string(separator))}, args[depth:]...)
		// 	rt.Handler(&router.Context{
		// 		Route:   rt,
		// 		Client:  tc,
		// 		Author:  &twitchAuthor{&user},
		// 		Msg:     &twitchMessage{&msg},
		// 		Command: cmd,
		// 		Args:    args,
		// 	})
		// }
	})

	tc.Join("gofur")

	err := tc.Connect()
	if err != nil {
		panic(err)
	}
}

func (tc *twitchClient) Stop() {
	tc.Say("gofur", "shutting down MrDestructoid")
	time.Sleep(time.Second / 2) // give some time for the message to send, if it doesn't no biggie
	tc.Disconnect()
}

type twitchAuthor struct{ *twitch.User }

func (ta *twitchAuthor) Name() string { return ta.Username }
func (ta *twitchAuthor) ID() string   { return ta.UserID }

type twitchMessage struct{ *twitch.Message }

func (tm *twitchMessage) Text() string { return tm.Message.Text }
func (tm *twitchMessage) ID() string   { return tm.Tags["id"] }
