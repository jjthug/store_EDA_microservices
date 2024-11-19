package waiter

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type WaitFunc func(ctx context.Context) error

type CleanupFunc func()

type Waiter interface {
	Add(fns ...WaitFunc)
	Cleanup(fns ...CleanupFunc)
	Wait() error
	Context() context.Context
	CancelFunc() context.CancelFunc
}

type waiter struct {
	ctx         context.Context
	waitFuncs   []WaitFunc
	cleanupFuns []CleanupFunc
	cancel      context.CancelFunc
}

type waiterCfg struct {
	parentCtx    context.Context
	catchSignals bool
}

func New(options ...WaiterOption) Waiter {
	cfg := &waiterCfg{
		parentCtx:    context.Background(),
		catchSignals: false,
	}

	for _, option := range options {
		option(cfg)
	}

	w := &waiter{
		waitFuncs:   []WaitFunc{},
		cleanupFuns: []CleanupFunc{},
	}
	w.ctx, w.cancel = context.WithCancel(cfg.parentCtx)
	if cfg.catchSignals {
		w.ctx, w.cancel = signal.NotifyContext(w.ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}
	return w
}

func (w *waiter) Add(funcs ...WaitFunc) {
	w.waitFuncs = append(w.waitFuncs, funcs...)
}

func (w *waiter) Cleanup(fns ...CleanupFunc) {
	w.cleanupFuns = append(w.cleanupFuns, fns...)
}

func (w *waiter) Wait() (err error) {
	g, ctx := errgroup.WithContext(w.ctx)
	g.Go(func() error {
		<-ctx.Done()
		w.cancel()
		return nil
	})

	for _, fn := range w.waitFuncs {
		waitFn := fn
		g.Go(func() error { return waitFn(ctx) })
	}

	for _, fn := range w.cleanupFuns {
		cleanupFn := fn
		defer cleanupFn()
	}

	return g.Wait()
}

func (w *waiter) Context() context.Context {
	return w.ctx
}

func (w *waiter) CancelFunc() context.CancelFunc {
	return w.cancel
}
