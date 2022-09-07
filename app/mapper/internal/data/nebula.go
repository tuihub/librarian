package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/mapper/internal/biz"
	"github.com/tuihub/librarian/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/zhihu/norm"
	"github.com/zhihu/norm/dialectors"
)

type nebulaMapperRepo struct {
	data *Data
}

func (r *nebulaMapperRepo) Save(ctx context.Context, g *biz.Mapper) (*biz.Mapper, error) {
	return g, nil
}

// NewNebula .
func NewNebula(c *conf.Mapper_Data) (*norm.DB, func()) {
	if c == nil || c.GetNebula() == nil {
		return nil, func() {}
	}

	dialector, err := dialectors.NewNebulaDialector(dialectors.DialectorConfig{
		Addresses: c.GetNebula().GetAddress(),
		Timeout:   time.Second * 5,
		Space:     c.GetNebula().GetSpace(),
		Username:  c.GetNebula().GetUsername(),
		Password:  c.GetNebula().GetPassword(),
	})
	if err != nil {
		log.Errorf("Failed to initialize Nebula Dialector, %s", err.Error())
		return nil, nil
	}

	db, err := norm.Open(dialector, norm.Config{}, norm.WithLogger(NebulaLoggerWrapper{}))
	if err != nil {
		log.Errorf("Failed to initialize Nebula DB, %s", err.Error())
		return nil, nil
	}
	return db, func() {
		log.Info("closing the data resources")
		db.Close()
		dialector.Close()
	}
}

type NebulaLoggerWrapper struct{}

func (l NebulaLoggerWrapper) Info(msg string) {
	log.Info(msg)
}

func (l NebulaLoggerWrapper) Warn(msg string) {
	log.Warn(msg)
}

func (l NebulaLoggerWrapper) Error(msg string) {
	log.Error(msg)
}

func (l NebulaLoggerWrapper) Fatal(msg string) {
	log.Fatal(msg)
}
