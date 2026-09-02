package httprouter

import (
	"context"
	"net/http"
	"sync"
)

type Handle func(http.ResponseWriter, *http.Request, Params)

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (ps Params) ByName(name string) string { _ = "STUB: not implemented"; return "" }

type paramsKey struct{}

var ParamsKey = paramsKey{}

func ParamsFromContext(ctx context.Context) Params { _ = "STUB: not implemented"; return *new(Params) }

var MatchedRoutePathParam = "$matchedRoutePath"

func (ps Params) MatchedRoutePath() string { _ = "STUB: not implemented"; return "" }

type Router struct {
	trees map[string]*node

	paramsPool sync.Pool
	maxParams  uint16

	SaveMatchedRoutePath bool

	RedirectTrailingSlash bool

	RedirectFixedPath bool

	HandleMethodNotAllowed bool

	HandleOPTIONS bool

	GlobalOPTIONS http.Handler

	globalAllowed string

	NotFound http.Handler

	MethodNotAllowed http.Handler

	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
}

var _ http.Handler = New()

func New() *Router { _ = "STUB: not implemented"; return nil }

func (r *Router) getParams() *Params { _ = "STUB: not implemented"; return nil }

func (r *Router) putParams(ps *Params) { _ = "STUB: not implemented"; return }

func (r *Router) saveMatchedRoutePath(path string, handle Handle) Handle {
	_ = "STUB: not implemented"
	return *new(Handle)
}

func (r *Router) GET(path string, handle Handle) { _ = "STUB: not implemented"; return }

func (r *Router) HEAD(path string, handle Handle) { _ = "STUB: not implemented"; return }

func (r *Router) OPTIONS(path string, handle Handle) { _ = "STUB: not implemented"; return }

func (r *Router) POST(path string, handle Handle) { _ = "STUB: not implemented"; return }

func (r *Router) PUT(path string, handle Handle) { _ = "STUB: not implemented"; return }

func (r *Router) PATCH(path string, handle Handle) { _ = "STUB: not implemented"; return }

func (r *Router) DELETE(path string, handle Handle) { _ = "STUB: not implemented"; return }

func (r *Router) Handle(method, path string, handle Handle) { _ = "STUB: not implemented"; return }

func (r *Router) Handler(method, path string, handler http.Handler) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) HandlerFunc(method, path string, handler http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) ServeFiles(path string, root http.FileSystem) { _ = "STUB: not implemented"; return }

func (r *Router) recv(w http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func (r *Router) Lookup(method, path string) (Handle, Params, bool) {
	_ = "STUB: not implemented"
	return *new(Handle), *new(Params), false
}

func (r *Router) allowed(path, reqMethod string) (allow string) {
	_ = "STUB: not implemented"
	return ""
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}
