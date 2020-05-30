package router

import (
	"io"
)

func MakeResp(writeToFn func(w io.Writer) (n int64, err error)) Resp {
	return genericResp{writeToFn: writeToFn}
}

type genericResp struct {
	writeToFn func(w io.Writer) (n int64, err error)
	content   interface{}
	mention   bool
}

func (r genericResp) Reply(content interface{}) Resp {
	r.content = content
	return r
}

func (r genericResp) Mention(b ...bool) Resp {
	r.mention = true
	if len(b) > 0 {
		r.mention = b[0]
	}
	return r
}

func (r genericResp) WriteTo(w io.Writer) (n int64, err error) {
	return r.writeToFn(w)
}

type Resp interface {
	io.WriterTo
	Reply(interface{}) Resp
	Mention(...bool) Resp
}
