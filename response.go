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
func SetError(rCtx *fasthttp.RequestCtx, err error) {

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
	rCtx.SetBody(b)
}

func SetResult(rCtx *fasthttp.RequestCtx, result interface{}) {

	id := rCtx.UserValue("id")
	if _, ok := id.(string); !ok {
		fmt.Println("setResult id not found")
		return
	}
	b1, err1 := cbor.Marshal(result)
	if err1 != nil {
		fmt.Println("cbor.Marshal err = ", err1.Error())
		return
	}

	args := &Base{}
	args.Err = ""
	args.Data = b1

	b2, err2 := proto.Marshal(args)
	if err2 != nil {
		fmt.Println("SetResult args.MarshalBinary = ", err2.Error())
		return
	}
	rCtx.SetBody(b2)
}
