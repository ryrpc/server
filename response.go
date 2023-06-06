package rysrv

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/valyala/fasthttp"
)

// SetError writes JSON-RPC response with error.
//
// It overwrites previous calls of SetResult and SetError.
func SetError(rCtx *fasthttp.RequestCtx, err error) {

	id := rCtx.UserValue("id")
	if _, ok := id.(string); !ok {
		fmt.Println("SetError id not found")
		return
	}

	args := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(args)

	args.Add("jsonrpc", "2.0")
	args.Add("id", id.(string))
	args.Add("error", err.Error())

	qs := args.QueryString()

	rCtx.SetBody(qs)
}

func SetResult(rCtx *fasthttp.RequestCtx, result interface{}) {

	id := rCtx.UserValue("id")
	if _, ok := id.(string); !ok {
		fmt.Println("setResult id not found")
		return
	}
	b, err := cbor.Marshal(result)
	if err != nil {
		fmt.Println("cbor.Marshal err = ", err)
		return
	}

	args := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(args)

	args.Add("jsonrpc", "2.0")
	args.Add("id", id.(string))

	args.AddBytesV("result", b)

	qs := args.QueryString()

	rCtx.SetBody(qs)
}
