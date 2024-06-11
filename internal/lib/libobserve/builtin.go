package libobserve

import (
	"container/ring"
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewBuiltInObserver)

const recordSize = 12

type BuiltInObserver struct {
	request *ObserverCounter
	mq      map[string]*ObserverCounter
	mqMu    sync.Mutex
	stopCh  chan struct{}
}

func NewBuiltInObserver() (*BuiltInObserver, error) {
	return &BuiltInObserver{
		request: NewObserverCounter(),
		mq:      make(map[string]*ObserverCounter),
		mqMu:    sync.Mutex{},
		stopCh:  make(chan struct{}),
	}, nil
}

func (o *BuiltInObserver) Start(ctx context.Context) error {
	scheduledJob := func() {
		o.request.next()
		o.mqMu.Lock()
		for _, v := range o.mq {
			v.next()
		}
		o.mqMu.Unlock()
		o.generateReport()
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-o.stopCh:
				return
			default:
				time.Sleep(time.Hour)
				scheduledJob()
			}
		}
	}()
	return nil
}

func (o *BuiltInObserver) Stop(_ context.Context) error {
	o.stopCh <- struct{}{}
	return nil
}

func (o *BuiltInObserver) GetRequest() *ObserverCounter {
	return o.request
}

func (o *BuiltInObserver) NewMQ(key string) *ObserverCounter {
	o.mqMu.Lock()
	defer o.mqMu.Unlock()
	if _, ok := o.mq[key]; !ok {
		o.mq[key] = NewObserverCounter()
	}
	return o.mq[key]
}

type ObserverCounter struct {
	counter       *statusCounter
	counterMu     sync.RWMutex
	counterRing   *ring.Ring
	counterRingMu sync.Mutex
}

func NewObserverCounter() *ObserverCounter {
	return &ObserverCounter{
		counter:       newRequestCounter(),
		counterMu:     sync.RWMutex{},
		counterRing:   ring.New(recordSize + 1),
		counterRingMu: sync.Mutex{},
	}
}

func (o *ObserverCounter) next() {
	o.counterMu.Lock()
	defer o.counterMu.Unlock()
	o.counterRingMu.Lock()
	defer o.counterRingMu.Unlock()

	rc := newRequestCounter()
	o.counterRing.Value = nil
	o.counterRing = o.counterRing.Next()
	o.counterRing.Value = rc
	o.counter = rc
}

func (o *ObserverCounter) Success() {
	o.counterMu.RLock()
	defer o.counterMu.RUnlock()
	o.counter.success.Add(1)
}

func (o *ObserverCounter) Failure() {
	o.counterMu.RLock()
	defer o.counterMu.RUnlock()
	o.counter.failure.Add(1)
}

type statusCounter struct {
	success atomic.Int64
	failure atomic.Int64
}

func newRequestCounter() *statusCounter {
	return &statusCounter{
		success: atomic.Int64{},
		failure: atomic.Int64{},
	}
}

func (o *BuiltInObserver) generateReport() {}
