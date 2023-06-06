package rysrv

import (
	"errors"

	"github.com/fasthttp/router"
	"github.com/fxamacker/cbor/v2"
	"github.com/valyala/fasthttp"
)

// NewRepository returns empty repository.
//
// It's safe to use Repository default value.
func NewRepository() *Repository {

	repo := &Repository{}
	repo.handlers = router.New()
	return repo
}

// Repository is a JSON-RPC 2.0 methods repository.
type Repository struct {

	//handlersMu sync.RWMutex
	handlers *router.Router
}

// RequestHandler is suitable for using with fasthttp.
func (r *Repository) RequestHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if !ctx.IsPost() {
			ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			return
		}

		if ctx.PostArgs().Len() != 4 {
			SetError(ctx, renderedParseError)
			return
		}

		if id := ctx.PostArgs().Peek("id"); id != nil {
			ctx.SetUserValue("id", string(id))
		}

		r.handlers.Handler(ctx)
	}
}

// Register registers new method handler.
func (r *Repository) Register(method string, handler func(ctx *fasthttp.RequestCtx)) {

	r.handlers.POST(method, handler)
}

func Unmarshal(ctx *fasthttp.RequestCtx, v interface{}) error {

	if ctx.PostArgs().Has("params") {

		data := ctx.PostArgs().Peek("params")
		return cbor.Unmarshal(data, v)
	}
	return errors.New("params error")
}
