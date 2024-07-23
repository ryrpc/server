package rysrv

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/golang/protobuf/proto"
	"github.com/valyala/fasthttp"
)

// SetError writes JSON-RPC response with error.
//
// It overwrites previous calls of SetResult and SetError.

func SetError(fctx *fasthttp.RequestCtx, err error) {

	/*
		id := rCtx.UserValue("id")
		if _, ok := id.(string); !ok {
			fmt.Println("SetError id not found")
			return
		}
	*/

	args := &Base{}
	args.Err = err.Error()

	b, err := proto.Marshal(args)
	if err != nil {
		fmt.Println("SetError args.MarshalBinary = ", err.Error())
		return
	}
	fctx.SetBody(b)
}

func SetResult(fctx *fasthttp.RequestCtx, result interface{}) {

	args := &Base{}
	args.Err = ""

	if vv, ok := result.(string); ok {
		args.Data = []byte(vv)
	} else if vv, ok := result.([]byte); ok {
		args.Data = vv
	} else {
		b1, err1 := cbor.Marshal(result)
		if err1 != nil {
			fmt.Println("cbor.Marshal err = ", err1.Error())
			return
		}
		args.Data = b1
	}

	b2, err2 := proto.Marshal(args)
	if err2 != nil {
		fmt.Println("SetResult args.MarshalBinary = ", err2.Error())
		return
	}
	fctx.Response.Header.Set("Connection", "keep-alive")
	fctx.SetBody(b2)
}
