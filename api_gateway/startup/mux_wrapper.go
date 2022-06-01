package startup

import (
	"container/list"
	"context"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
			runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler,
				writer http.ResponseWriter, request *http.Request, err error) {
				st, _ := status.FromError(err)
				newError := runtime.HTTPStatusError{
					HTTPStatus: runtime.HTTPStatusFromCode(st.Code()),
					Err:        errors.New(st.Message()),
				}
				runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &newError)
			}),
			runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
				header := request.Header.Get("sub")
				md := metadata.Pairs("sub", header)
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
