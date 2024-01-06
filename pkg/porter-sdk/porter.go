package portersdk

import (
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	"os"
	"time"
)

type Porter struct {
	server     *grpc.Server
	controller controller
	logger     log.Logger
	app        *kratos.App
}

type PorterConfig struct {
	Name           string
	Version        string
	GlobalName     string
	FeatureSummary *pb.PorterFeatureSummary
	Server         ServerConfig
}

type ServerConfig struct {
	Network string
	Addr    string
	Timeout *time.Duration
}

type Option func(*Porter)

func WithLogger(logger log.Logger) Option {
	return func(p *Porter) {
		p.logger = logger
	}
}

func (p *Porter) Run() error {
	return p.app.Run()
}

func (p *Porter) Stop() error {
	return p.app.Stop()
}

func New(config *PorterConfig, handler Handler, options ...Option) (*Porter, error) {
	if config == nil {
		return nil, fmt.Errorf("config is nil")
	}
	if handler == nil {
		return nil, fmt.Errorf("handler is nil")
	}
	p := new(Porter)
	for _, o := range options {
		o(p)
	}
	client, err := newSephirahClient()
	if err != nil {
		return nil, err
	}
	c := controller{
		handler: handler,
		config:  config,
		logger:  p.logger,
		client:  client,
	}
	p.controller = c
	p.server = newServer(
		&config.Server,
		newService(c),
		p.logger,
	)
	id, _ := os.Hostname()
	name := "porter"
	app := kratos.New(
		kratos.ID(id+name),
		kratos.Name(name),
		kratos.Version(p.controller.config.Version),
		kratos.Metadata(map[string]string{
			"PorterName": p.controller.config.GlobalName,
		}),
		kratos.Server(p.server),
	)
	p.app = app
	return p, nil
}
