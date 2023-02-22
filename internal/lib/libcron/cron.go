package libcron

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCron)

type Cron struct {
	scheduler *gocron.Scheduler
}

func NewCron() *Cron {
	s := gocron.NewScheduler(time.UTC)

	return &Cron{
		s,
	}
}

func (c *Cron) Start(ctx context.Context) error {
	c.scheduler.StartBlocking()
	return nil
}
func (c *Cron) Stop(ctx context.Context) error {
	c.scheduler.Stop()
	return nil
}

func (c *Cron) ByCronExpr(expr string, jobFunc interface{}, params ...interface{}) error {
	_, err := c.scheduler.Cron(expr).Do(jobFunc, params...)
	return err
}

func (c *Cron) BySeconds(seconds int, jobFunc interface{}, params ...interface{}) error {
	_, err := c.scheduler.Every(seconds).Seconds().Do(jobFunc, params...)
	return err
}
