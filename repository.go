package rysrv

import (
	"github.com/fasthttp/router"
	"github.com/fxamacker/cbor/v2"
	"github.com/valyala/fasthttp"
)

// Repository is a JSON-RPC 2.0 methods repository.
type Repository struct {

	//handlersMu sync.RWMutex
	obj      map[string]interface{}
	handlers *router.Router
}

// NewRepository returns empty repository.
//
// It's safe to use Repository default value.
func NewRepository() *Repository {

	repo := &Repository{}
	repo.obj = map[string]interface{}{}
	repo.handlers = router.New()
	return repo
}

func (r *Repository) Use(name string, obj interface{}) {
	r.obj[name] = obj
}

// RequestHandler is suitable for using with fasthttp.
func (r *Repository) RequestHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if !ctx.IsPost() {
			ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			return
		}

		/*
			if ctx.PostArgs().Len() != 4 {
				SetError(ctx, renderedParseError)
				return
			}

			if id := ctx.PostArgs().Peek("id"); id != nil {
				ctx.SetUserValue("id", string(id))
			}
		*/
		for k, v := range r.obj {
			ctx.SetUserValue(k, v)
		}
		r.handlers.Handler(ctx)
	}
}

// Register registers new method handler.
func (r *Repository) Register(method string, handler func(ctx *fasthttp.RequestCtx)) {

	r.handlers.POST(method, handler)
}

func Unmarshal(ctx *fasthttp.RequestCtx, v interface{}) error {

	data := ctx.PostBody()
	return cbor.Unmarshal(data, v)
}
