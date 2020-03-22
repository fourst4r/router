package router

import "fmt"

type Context struct {
	Route   *Route
	Client  Client
	Channel Channel
	Author  Author
	Msg     Msg
	Command string
	Args    []string
}

// Say is a shortcut for Context.Client.Say()
func (c *Context) Say(ch Channel, s string) {
	c.Client.Say(ch, s)
}

func (c *Context) Reply(s ...string) {
	c.Client.Say(c.Channel, fmt.Sprintf("%s %s", c.Author.Mention(), fmt.Sprint(s)))
}

type Client interface {
	ID() string
	Say(Channel, string)
}

type Channel interface {
	ID() string
	Name() string
}

type Author interface {
	ID() string
	Name() string
	Mention() string
}

type Msg interface {
	ID() string
	Text() string
}
