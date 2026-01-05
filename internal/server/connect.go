package server

import (
	"context"
	"fmt"
	"net"
	stdhttp "net/http"
	"strconv"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1/v1connect"
	sentinel "github.com/tuihub/protos/pkg/librarian/sentinel/v1/v1connect"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1/v1connect"

	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	ReadHeaderTimeout = 30 * time.Second
)

type ConnectServer struct {
	srv *http.Server
}

func NewConnectServer(
	c *conf.Server,
	auth *libauth.Auth,
	sephirahservice sephirah.LibrarianSephirahServiceHandler,
	sentinelservice sentinel.LibrarianSephirahSentinelServiceHandler,
	porterservice porter.LibrarianSephirahPorterServiceHandler,
	app *libapp.Settings,
) (*ConnectServer, error) {
	mux := stdhttp.NewServeMux()
	checker := grpchealth.NewStaticChecker(
		"librarian.sephirah.v1.LibrarianSephirahService",
		"librarian.porter.v1.LibrarianSephirahPorterService",
		"librarian.sentinel.v1.LibrarianSentinelService",
	)
	mux.Handle(grpchealth.NewHandler(checker))
	reflector := grpcreflect.NewStaticReflector(
		"librarian.sephirah.v1.LibrarianSephirahService",
		"librarian.porter.v1.LibrarianSephirahPorterService",
		"librarian.sentinel.v1.LibrarianSentinelService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	mux.Handle(sephirah.NewLibrarianSephirahServiceHandler(sephirahservice))
	mux.Handle(sentinel.NewLibrarianSephirahSentinelServiceHandler(sentinelservice))
	mux.Handle(porter.NewLibrarianSephirahPorterServiceHandler(porterservice))

	var middlewares = []middleware.Middleware{
		logging.Server(libapp.GetLogger()),
		recovery.Recovery(),
	}
	middlewares = append(middlewares, NewTokenMatcher(auth)...)

	srv := http.NewServer(
		http.Address(net.JoinHostPort(c.Main.Host, strconv.Itoa(int(c.Main.Port)))),
		http.Filter(func(h stdhttp.Handler) stdhttp.Handler {
			return h2c.NewHandler(h, &http2.Server{})
		}),
		http.Timeout(ReadHeaderTimeout),
	)
	srv.HandlePrefix("/", wrapMiddleware(mux, middlewares))

	return &ConnectServer{
		srv: srv,
	}, nil
}

func (s *ConnectServer) Start(ctx context.Context) error {
	err := s.srv.Start(ctx)
	if err == nil && ctx.Err() == nil {
		return fmt.Errorf("ConnectServer stopped unexpectedly")
	}
	return err
}

func (s *ConnectServer) Stop(ctx context.Context) error {
	return s.srv.Stop(ctx)
}

func wrapMiddleware(h stdhttp.Handler, ms []middleware.Middleware) stdhttp.Handler {
	chain := middleware.Chain(ms...)
	return stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		ctx := r.Context()
		if tr, ok := transport.FromServerContext(ctx); ok {
			if tr.Operation() == "/" && r.URL.Path != "/" {
				ctx = transport.NewServerContext(ctx, &Transport{
					KindVal:        tr.Kind(),
					EndpointVal:    tr.Endpoint(),
					OperationVal:   r.URL.Path,
					ReqHeaderVal:   tr.RequestHeader(),
					ReplyHeaderVal: tr.ReplyHeader(),
				})
			}
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
			return nil, nil
		}
		if _, err := chain(next)(ctx, r); err != nil {
			http.DefaultErrorEncoder(w, r, err)
		}
	})
}

type Transport struct {
	KindVal        transport.Kind
	EndpointVal    string
	OperationVal   string
	ReqHeaderVal   transport.Header
	ReplyHeaderVal transport.Header
}

func (t *Transport) Kind() transport.Kind            { return t.KindVal }
func (t *Transport) Endpoint() string                { return t.EndpointVal }
func (t *Transport) Operation() string               { return t.OperationVal }
func (t *Transport) RequestHeader() transport.Header { return t.ReqHeaderVal }
func (t *Transport) ReplyHeader() transport.Header   { return t.ReplyHeaderVal }
