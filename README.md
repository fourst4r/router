# router
A fork of [dgrouter](https://github.com/Necroforger/dgrouter/)

## example
```go 
router.On("ping", func(ctx *router.Context) { ctx.Reply("pong")}).Desc("responds with pong")

router.On("avatar", func(ctx *router.Context) {
	ctx.Reply(ctx.Msg.Author.AvatarURL("2048"))
}).Desc("returns the user's avatar")

router.Default = router.On("help", func(ctx *router.Context) {
	var text = ""
	for _, v := range router.Routes {
		text += v.Name + " : \t" + v.Description + "\n"
	}
	ctx.Reply("```" + text + "```")
}).Desc("prints this help menu")
```
