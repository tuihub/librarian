package libcron

import (
	"math"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

type sentryListener struct {
	idMap     map[uuid.UUID]*sentry.EventID
	configMap map[string]*sentry.MonitorConfig
}

func newSentryListener() *sentryListener {
	return &sentryListener{
		idMap:     make(map[uuid.UUID]*sentry.EventID),
		configMap: make(map[string]*sentry.MonitorConfig),
	}
}

func (s sentryListener) NewDurationJob(name string, duration time.Duration) {
	s.configMap[name] = &sentry.MonitorConfig{ //nolint:exhaustruct // no need
		Schedule:      sentry.IntervalSchedule(int64(math.Ceil(duration.Minutes())), sentry.MonitorScheduleUnitMinute),
		CheckInMargin: 1,
	}
}

func (s sentryListener) EventListeners() []gocron.EventListener {
	return []gocron.EventListener{
		gocron.BeforeJobRuns(func(jobID uuid.UUID, jobName string) {
			s.idMap[jobID] = sentry.CaptureCheckIn(
				&sentry.CheckIn{ //nolint:exhaustruct // no need
					MonitorSlug: jobName,
					Status:      sentry.CheckInStatusInProgress,
				},
				s.configMap[jobName],
			)
		}),
		gocron.AfterJobRuns(func(jobID uuid.UUID, jobName string) {
			sentry.CaptureCheckIn(
				&sentry.CheckIn{ //nolint:exhaustruct // no need
					ID:          *s.idMap[jobID],
					MonitorSlug: jobName,
					Status:      sentry.CheckInStatusOK,
				},
				s.configMap[jobName],
			)
		}),
		gocron.AfterJobRunsWithError(func(jobID uuid.UUID, jobName string, err error) {
			sentry.CaptureCheckIn(
				&sentry.CheckIn{ //nolint:exhaustruct // no need
					ID:          *s.idMap[jobID],
					MonitorSlug: jobName,
					Status:      sentry.CheckInStatusError,
				},
				s.configMap[jobName],
			)
		}),
	}
}
