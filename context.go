package router

type Context struct {
	Route   *Route
	Client  Client
	Author  Author
	Msg     Msg
	Command string
	Args    []string
}

type Client interface {
	ID() string
}

type Author interface {
	ID() string
	Name() string
}

type Msg interface {
	ID() string
	Text() string
}
