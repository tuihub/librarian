package conf

import (
	"fmt"

	"github.com/samber/lo"
)

type ConfigDigest struct {
	Name    string
	Enabled *bool
	Driver  *string
	Listen  *string
}

func (d *ConfigDigest) Status() string {
	if d.Driver != nil {
		return fmt.Sprintf("Enable - Driver %s", *d.Driver)
	} else if d.Listen != nil {
		return fmt.Sprintf("Enable - Listen on %s", *d.Listen)
	} else if d.Enabled != nil {
		return "Enable"
	} else {
		return "Disable"
	}
}

func (d *ConfigDigest) String() string {
	return fmt.Sprintf("[%s\t] %s", d.Name, d.Status())
}

func GenConfigDigest(c *Config) []*ConfigDigest {
	var digests []*ConfigDigest

	digests = append(digests, &ConfigDigest{
		Name:    "Server gRPC",
		Enabled: lo.ToPtr(c.Server != nil && c.Server.GetGrpc() != nil),
		Driver:  nil,
		Listen:  lo.ToPtr(c.Server.GetGrpc().GetAddr()),
	})
	digests = append(digests, &ConfigDigest{
		Name:    "Server gRPC-Web",
		Enabled: lo.ToPtr(c.Server != nil && c.Server.GetGrpcWeb() != nil),
		Driver:  nil,
		Listen:  lo.ToPtr(c.Server.GetGrpcWeb().GetAddr()),
	})
	digests = append(digests, &ConfigDigest{
		Name:    "DB",
		Enabled: lo.ToPtr(c.Database != nil && len(c.Database.Driver) != 0),
		Driver:  lo.ToPtr(string(c.Database.Driver)),
		Listen:  nil,
	})
	digests = append(digests, &ConfigDigest{
		Name:    "MQ",
		Enabled: lo.ToPtr(c.MQ != nil && len(c.MQ.Driver) != 0),
		Driver:  lo.ToPtr(string(c.MQ.Driver)),
		Listen:  nil,
	})
	digests = append(digests, &ConfigDigest{
		Name:    "Cache",
		Enabled: lo.ToPtr(c.Cache != nil && len(c.Cache.Driver) != 0),
		Driver:  lo.ToPtr(string(c.Cache.Driver)),
		Listen:  nil,
	})
	digests = append(digests, &ConfigDigest{
		Name:    "Storage",
		Enabled: lo.ToPtr(c.Storage != nil && len(c.Storage.Driver) != 0),
		Driver:  lo.ToPtr(string(c.Storage.Driver)),
		Listen:  nil,
	})
	digests = append(digests, &ConfigDigest{
		Name:    "Consul",
		Enabled: lo.ToPtr(c.Consul != nil && len(c.Consul.GetAddr()) != 0),
		Driver:  nil,
		Listen:  nil,
	})
	digests = append(digests, &ConfigDigest{
		Name:    "OTLP",
		Enabled: lo.ToPtr(c.Otlp != nil && len(c.Otlp.GetProtocol()) != 0),
		Driver:  nil,
		Listen:  nil,
	})

	return digests
}
