package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/mapper/internal/biz"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/zhihu/norm"
	"github.com/zhihu/norm/dialectors"
)

type nebulaMapperRepo struct {
	db *norm.DB
}

func (r *nebulaMapperRepo) InsertVertex(ctx context.Context, vertex []*biz.Vertex) error {
	// TODO implement me
	panic("implement me")
}

func (r *nebulaMapperRepo) InsertEdge(ctx context.Context, edge []*biz.Edge) error {
	// TODO implement me
	panic("implement me")
}

func (r *nebulaMapperRepo) FetchEqualVertex(ctx context.Context, vertices biz.Vertex) ([]*biz.Vertex, error) {
	// TODO implement me
	panic("implement me")
}

// NewNebula .
func NewNebula(c *conf.Mapper_Data) (*norm.DB, func()) {
	if c == nil || c.GetNebula() == nil {
		return nil, func() {}
	}

	dialector, err := dialectors.NewNebulaDialector(dialectors.DialectorConfig{
		Addresses: c.GetNebula().GetAddress(),
		Timeout:   time.Second * 5, //nolint:gomnd //TODO
		Space:     c.GetNebula().GetSpace(),
		Username:  c.GetNebula().GetUsername(),
		Password:  c.GetNebula().GetPassword(),
	})
	if err != nil {
		logger.Errorf("Failed to initialize Nebula Dialector, %s", err.Error())
		return nil, nil
	}

	db, err := norm.Open(dialector, norm.Config{}, norm.WithLogger(NebulaLoggerWrapper{}))
	if err != nil {
		logger.Errorf("Failed to initialize Nebula DB, %s", err.Error())
		return nil, nil
	}
	return db, func() {
		logger.Info("closing the data resources")
		db.Close()
		dialector.Close()
	}
}

type NebulaLoggerWrapper struct{}

func (l NebulaLoggerWrapper) Info(msg string) {
	logger.Info(msg)
}

func (l NebulaLoggerWrapper) Warn(msg string) {
	logger.Warn(msg)
}

func (l NebulaLoggerWrapper) Error(msg string) {
	logger.Error(msg)
}

func (l NebulaLoggerWrapper) Fatal(msg string) {
	logger.Fatal(msg)
}
