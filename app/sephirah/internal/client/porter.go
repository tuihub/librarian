package client

import (
	"context"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libtime"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type Porter struct {
	porter.LibrarianPorterServiceClient
	checker libapp.HealthChecker
}

func NewPorter(
	client porter.LibrarianPorterServiceClient,
	consul *conf.Consul,
	porter *conf.Porter,
) (*Porter, error) {
	checker, err := libapp.NewHealthChecker("porter", consul)
	if err != nil {
		return nil, err
	}
	if libapp.IsEmptyHealthChecker(checker) {
		checker, err = newStaticDiscovery(porter)
		if err != nil {
			return nil, err
		}
	}
	return &Porter{
		LibrarianPorterServiceClient: client,
		checker:                      checker,
	}, nil
}

func (p *Porter) GetServiceAddresses(ctx context.Context) ([]string, error) {
	return p.checker.GetAliveInstances()
}

func NewPorterClient(c *conf.Consul, p *conf.Porter, app *libapp.Settings) (porter.LibrarianPorterServiceClient, error) {
	r, err := libapp.NewDiscovery(c)
	if err != nil {
		return nil, err
	}
	if libapp.IsEmptyDiscovery(r) {
		r, err = newStaticDiscovery(p)
		if err != nil {
			return nil, err
		}
	}
	middlewares := []grpc.ClientOption{
		grpc.WithEndpoint("discovery:///porter"),
		grpc.WithDiscovery(r),
		grpc.WithNodeFilter(
			newPorterNameFilter(),
			newPorterAddressFilter(),
			newPorterFastFailFilter(),
		),
		grpc.WithTimeout(libtime.Minute),
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, grpc.WithMiddleware(recovery.Recovery()))
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		middlewares...,
	)
	cli := porter.NewLibrarianPorterServiceClient(conn)
	return cli, err
}

type requestPorterName struct{}
type requestPorterAddress struct{}
type requestPorterFastFail struct{}

const porterNameKey = "PorterName"

func WithPorterName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, requestPorterName{}, name)
}
func WithPorterAddress(ctx context.Context, address string) context.Context {
	return context.WithValue(ctx, requestPorterAddress{}, address)
}
func WithPorterFastFail(ctx context.Context) context.Context {
	return context.WithValue(ctx, requestPorterFastFail{}, true)
}

func newPorterNameFilter() selector.NodeFilter {
	return func(ctx context.Context, nodes []selector.Node) []selector.Node {
		r, ok := ctx.Value(requestPorterName{}).(string)
		if !ok {
			return nodes
		}
		newNodes := make([]selector.Node, 0)
		for _, n := range nodes {
			n.InitialWeight()
			if v, exist := n.Metadata()[porterNameKey]; exist && v == r {
				newNodes = append(newNodes, n)
			}
		}
		return newNodes
	}
}

func newPorterAddressFilter() selector.NodeFilter {
	return func(ctx context.Context, nodes []selector.Node) []selector.Node {
		r, ok := ctx.Value(requestPorterAddress{}).(string)
		if !ok {
			return nodes
		}
		newNodes := make([]selector.Node, 0)
		for _, n := range nodes {
			n.InitialWeight()
			if n.Address() == r {
				newNodes = append(newNodes, n)
			}
		}
		return newNodes
	}
}

func newPorterFastFailFilter() selector.NodeFilter {
	return func(ctx context.Context, nodes []selector.Node) []selector.Node {
		r, ok := ctx.Value(requestPorterFastFail{}).(bool)
		if !ok || !r {
			return nodes
		}
		return make([]selector.Node, 0)
	}
}
