package rysrv

import (
	"github.com/bytedance/sonic"
	"github.com/valyala/fastjson"
)

//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=.

func (ctx *RequestCtx) writeString(s string) {
	_, _ = ctx.response.WriteString(s)
}

// SetError writes JSON-RPC response with error.
//
// It overwrites previous calls of SetResult and SetError.
func (ctx *RequestCtx) SetError(err error) {
	if len(ctx.id) == 0 {
		return
	}

	ctx.response.Reset()

	if err != nil {
		buf := ctx.bytebufferpool.Get()

		buf.SetString(`{"jsonrpc":"2.0","error":"`)
		buf.WriteString(err.Error())
		buf.WriteString(`","id":`)
		buf.WriteString(string(ctx.id))
		buf.WriteString(`}`)

		ctx.fasthttpCtx.SetBody(buf.Bytes())
		//ctx.writeString(buf.String())
		ctx.bytebufferpool.Put(buf)
	}
}

// SetResult writes JSON-RPC response with result.
//
// It overwrites previous calls of SetResult and SetError.
//
// result may be *fastjson.Value, []byte, or interface{} (slower).
func (ctx *RequestCtx) SetResult(result interface{}) {
	if len(ctx.id) == 0 {
		return
	}

	ctx.response.Reset()

	switch v := result.(type) {
	case *fastjson.Value:
		buf := ctx.bytebufferpool.Get()

		buf.SetString(`{"jsonrpc":"2.0","result":"`)
		buf.Write(v.MarshalTo(buf.B))
		buf.WriteString(`","id":`)
		buf.WriteString(string(ctx.id))
		buf.WriteString(`}`)
		ctx.fasthttpCtx.SetBody(buf.Bytes())

		ctx.bytebufferpool.Put(buf)
	default:
		buf := ctx.bytebufferpool.Get()

		//_ = sonic.ConfigDefault.NewEncoder(buf).Encode(result)
		output, _ := sonic.Marshal(result)
		buf.SetString(`{"jsonrpc":"2.0","result":"`)
		buf.Write(output)
		buf.WriteString(`","id":`)
		buf.WriteString(string(ctx.id))
		buf.WriteString(`}`)
		ctx.fasthttpCtx.SetBody(buf.Bytes())
		ctx.bytebufferpool.Put(buf)
	}
}
