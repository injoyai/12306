package logic

import (
	"context"
	"github.com/injoyai/base/chans"
	"time"
)

type timer struct {
	*chans.Rerun
	interval time.Duration
	handler  func(ctx context.Context)
}

func NewTimer(interval time.Duration, f func(ctx context.Context)) *timer {
	t := &timer{
		interval: interval,
		handler:  f,
	}
	t.Rerun = chans.NewRerun(func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(t.interval):
				t.handler(ctx)
			}
		}
	})
	return t
}
