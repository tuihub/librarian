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
		Enabled: lo.ToPtr(c.GetServer() != nil && c.GetServer().GetGrpc() != nil),
		Driver:  nil,
		Listen:  lo.ToPtr(c.GetServer().GetGrpc().GetAddr()),
	})
	digests = append(digests, &ConfigDigest{
		Name:    "Server gRPC-Web",
		Enabled: lo.ToPtr(c.GetServer() != nil && c.GetServer().GetGrpcWeb() != nil),
		Driver:  nil,
		Listen:  lo.ToPtr(c.GetServer().GetGrpcWeb().GetAddr()),
	})
	digests = append(digests, &ConfigDigest{
		Name:    "DB",
		Enabled: lo.ToPtr(c.GetDatabase() != nil && len(c.GetDatabase().GetDriver()) != 0),
		Driver:  lo.ToPtr(c.GetDatabase().GetDriver()),
		Listen:  nil,
	})
	digests = append(digests, &ConfigDigest{
		Name:    "MQ",
		Enabled: lo.ToPtr(c.GetMq() != nil && len(c.GetMq().GetDriver()) != 0),
		Driver:  lo.ToPtr(c.GetMq().GetDriver()),
		Listen:  nil,
	})
	digests = append(digests, &ConfigDigest{
		Name:    "Cache",
		Enabled: lo.ToPtr(c.GetCache() != nil && len(c.GetCache().GetDriver()) != 0),
		Driver:  lo.ToPtr(c.GetCache().GetDriver()),
		Listen:  nil,
	})
	digests = append(digests, &ConfigDigest{
		Name:    "S3",
		Enabled: lo.ToPtr(c.GetS3() != nil && len(c.GetS3().GetDriver()) != 0),
		Driver:  lo.ToPtr(c.GetS3().GetDriver()),
		Listen:  nil,
	})
	digests = append(digests, &ConfigDigest{
		Name:    "Consul",
		Enabled: lo.ToPtr(c.GetConsul() != nil && len(c.GetConsul().GetAddr()) != 0),
		Driver:  nil,
		Listen:  nil,
	})
	digests = append(digests, &ConfigDigest{
		Name:    "OTLP",
		Enabled: lo.ToPtr(c.GetOtlp() != nil && len(c.GetOtlp().GetProtocol()) != 0),
		Driver:  nil,
		Listen:  nil,
	})

	return digests
}
