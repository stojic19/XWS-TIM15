package startup

import (
	"container/list"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type MiddlewareType func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))

type MuxWrapper struct {
	runtime.ServeMux
	middlewares list.List
}

func NewMuxWrapper() *MuxWrapper {
	return &MuxWrapper{
		ServeMux: *runtime.NewServeMux(
			runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
				header := request.Header.Get("Server")
				md := metadata.Pairs("evo", header)
				ctx = context.WithValue(ctx, "evo", header)
				evo := ctx.Value("evo")
				fmt.Println(evo)
				return md
			})),
		middlewares: list.List{},
	}
}

func (mux *MuxWrapper) AppendMiddleware(middleware func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))) {
	mux.middlewares.PushBack(MiddlewareType(middleware))
}

func (mux *MuxWrapper) PrependMiddleware(middleware func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))) {
	mux.middlewares.PushFront(MiddlewareType(middleware))
}

func (mux *MuxWrapper) nextMiddleware(el *list.Element) func(w http.ResponseWriter, req *http.Request) {
	if el != nil {
		return func(w http.ResponseWriter, req *http.Request) {
			el.Value.(MiddlewareType)(w, req, mux.nextMiddleware(el.Next()))
		}
	}
	return mux.ServeMux.ServeHTTP
}

func (mux *MuxWrapper) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mux.nextMiddleware(mux.middlewares.Front())(w, req)
}
