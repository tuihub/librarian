package libcron

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCron)

type Cron struct {
	scheduler gocron.Scheduler
}

func NewCron() (*Cron, error) {
	s, err := gocron.NewScheduler(
		gocron.WithLocation(time.UTC),
		gocron.WithLogger(newCronLogger()),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create gocron scheduler: %w", err)
	}

	return &Cron{
		s,
	}, nil
}

func (c *Cron) Start(ctx context.Context) error {
	c.scheduler.Start()
	return nil
}
func (c *Cron) Stop(ctx context.Context) error {
	return c.scheduler.StopJobs()
}

func (c *Cron) BySeconds(name string, seconds int, jobFunc interface{}, params ...interface{}) error {
	return c.Duration(name, time.Duration(seconds)*time.Second, jobFunc, params...)
}

func (c *Cron) Duration(name string, duration time.Duration, jobFunc interface{}, params ...interface{}) error {
	_, err := c.scheduler.NewJob(
		gocron.DurationJob(duration),
		gocron.NewTask(jobFunc, params...),
		gocron.WithName(name),
	)
	return err
}
