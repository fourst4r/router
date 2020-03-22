package router

import (
	"fmt"
	"io"
)

type Context struct {
	Route   *Route
	Client  Client
	Channel Channel
	Author  Author
	Msg     Msg
	Command string
	Args    []string
	Writer  io.StringWriter
}

// type Context interface {
// 	Route() *Route
// 	Client() Client
// 	Channel() Channel
// 	Author() Author
// 	Msg() Msg
// 	Command() string
// 	Args() []string
// }

// Say is a shortcut for Context.Client.Say()
func (c *Context) Say(ch Channel, s string) {
	if c.Writer != nil {
		c.Writer.WriteString(s)
	} else {
		c.Client.Say(ch, s)
	}
}

// Reply sends s and mentions the user who ran the command.
func (c *Context) Reply(s ...string) {
	text := fmt.Sprintf("%s %s", c.Author.Mention(), fmt.Sprint(s))
	if c.Writer != nil {
		c.Writer.WriteString(text)
	} else {
		c.Client.Say(c.Channel, text)
	}
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
