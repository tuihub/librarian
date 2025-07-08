package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libdiscovery"
	"github.com/tuihub/librarian/internal/lib/libtime"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/fullstorydev/grpchan/inprocgrpc"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/samber/lo"
	grpcstd "google.golang.org/grpc"
)

type Porter struct {
	porter.LibrarianPorterServiceClient
	checker    libdiscovery.HealthChecker
	inprocKeys []string
}

func NewPorter(
	client porter.LibrarianPorterServiceClient,
	consul *conf.Consul,
	porter *conf.Porter,
	inprocPorters *InprocPorter,
) (*Porter, error) {
	checker, err := libdiscovery.NewHealthChecker("porter", consul)
	if err != nil {
		return nil, fmt.Errorf("failed to create health checker: %w", err)
	}
	if libdiscovery.IsEmptyHealthChecker(checker) && porter != nil && len(porter.Addresses) > 0 {
		checker, err = libdiscovery.NewStaticDiscovery(porter.Addresses, "porter", "")
		if err != nil {
			return nil, fmt.Errorf("failed to create static discovery: %w", err)
		}
	}
	return &Porter{
		LibrarianPorterServiceClient: client,
		checker:                      checker,
		inprocKeys:                   lo.Keys(inprocPorters.Instances),
	}, nil
}

func (p *Porter) GetServiceAddresses(ctx context.Context) ([]string, error) {
	res, err := p.checker.GetAliveInstances()
	if err != nil {
		return nil, fmt.Errorf("failed to get alive instances: %w", err)
	}
	return append(res, p.inprocKeys...), nil
}

func NewPorterClient(
	c *conf.Consul,
	p *conf.Porter,
	app *libapp.Settings,
	inprocPorters *InprocPorter,
) (porter.LibrarianPorterServiceClient, error) {
	// inproc
	inprocConn := map[string]*inprocgrpc.Channel{}
	for k, inst := range inprocPorters.Instances {
		channel := inprocgrpc.Channel{}
		porter.RegisterLibrarianPorterServiceServer(&channel, inst)
		inprocConn[k] = &channel
	}

	// network
	r, err := libdiscovery.NewDiscovery(c)
	if err != nil {
		return nil, fmt.Errorf("failed to create discovery: %w", err)
	}
	if libdiscovery.IsEmptyDiscovery(r) && p != nil && len(p.Addresses) > 0 {
		r, err = libdiscovery.NewStaticDiscovery(p.Addresses, "porter", "")
		if err != nil {
			return nil, fmt.Errorf("failed to create static discovery: %w", err)
		}
	}
	middlewares := []grpc.ClientOption{
		grpc.WithEndpoint("discovery:///porter"),
		grpc.WithDiscovery(r),
		grpc.WithNodeFilter(
			newPorterAddressFilter(),
			newPorterFastFailFilter(),
		),
		grpc.WithTimeout(libtime.Minute),
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, grpc.WithMiddleware(recovery.Recovery()))
	}
	networkConn, err := grpc.DialInsecure(
		context.Background(),
		middlewares...,
	)

	cli := porter.NewLibrarianPorterServiceClient(&hybridClientConn{
		inproc:  inprocConn,
		network: networkConn,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}
	return cli, nil
}

type requestPorterAddress struct{}
type requestPorterFastFail struct{}

func WithPorterAddress(ctx context.Context, addresses []string) context.Context {
	return context.WithValue(ctx, requestPorterAddress{}, addresses)
}
func WithPorterFastFail(ctx context.Context) context.Context {
	return context.WithValue(ctx, requestPorterFastFail{}, true)
}

func newPorterAddressFilter() selector.NodeFilter {
	return func(ctx context.Context, nodes []selector.Node) []selector.Node {
		r, ok := ctx.Value(requestPorterAddress{}).([]string)
		if !ok {
			return nodes
		}
		newNodes := make([]selector.Node, 0)
		for _, n := range nodes {
			n.InitialWeight()
			for _, a := range r {
				if n.Address() == a {
					newNodes = append(newNodes, n)
				}
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

type hybridClientConn struct {
	inproc  map[string]*inprocgrpc.Channel
	network *grpcstd.ClientConn
}

func (h *hybridClientConn) Invoke(
	ctx context.Context,
	method string,
	args interface{},
	reply interface{},
	opts ...grpcstd.CallOption,
) error {
	r, ok := ctx.Value(requestPorterAddress{}).([]string)
	if ok {
		for _, k := range r {
			if channel, exists := h.inproc[k]; exists {
				return channel.Invoke(ctx, method, args, reply, opts...)
			}
		}
	}

	if h.network != nil {
		return h.network.Invoke(ctx, method, args, reply, opts...)
	}
	return errors.New("no available connection")
}

func (h *hybridClientConn) NewStream(
	ctx context.Context,
	desc *grpcstd.StreamDesc,
	method string,
	opts ...grpcstd.CallOption,
) (grpcstd.ClientStream, error) {
	r, ok := ctx.Value(requestPorterAddress{}).([]string)
	if ok {
		for _, k := range r {
			if channel, exists := h.inproc[k]; exists {
				return channel.NewStream(ctx, desc, method, opts...)
			}
		}
	}

	if h.network != nil {
		return h.network.NewStream(ctx, desc, method, opts...)
	}
	return nil, errors.New("no available connection")
}
