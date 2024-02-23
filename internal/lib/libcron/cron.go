package libcron

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCron)

type Cron struct {
	scheduler      gocron.Scheduler
	sentryListener *sentryListener
}

func NewCron() (*Cron, error) {
	sl := newSentryListener()
	s, err := gocron.NewScheduler(
		gocron.WithLocation(time.UTC),
		gocron.WithLogger(newCronLogger()),
		gocron.WithGlobalJobOptions(
			gocron.WithEventListeners(sl.EventListeners()...),
		),
	)

	if err != nil {
		return nil, err
	}

	return &Cron{
		s,
		sl,
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
	c.sentryListener.NewDurationJob(name, duration)
	_, err := c.scheduler.NewJob(
		gocron.DurationJob(duration),
		gocron.NewTask(jobFunc, params...),
		gocron.WithName(name),
	)
	return err
}
