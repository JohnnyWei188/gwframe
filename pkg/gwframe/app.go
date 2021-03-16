package gwframe

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/JohnnyWei188/gwframe/pkg/transport"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// App is an application components lifecycle manager
type App struct {
	opts    options
	ctx     context.Context
	cancel  func()
	servers []transport.Server
}

// New create an application lifecycle manager
func New(opts ...Option) *App {
	options := options{
		ctx:  context.Background(),
		sigs: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
	if id, err := uuid.NewUUID(); err == nil {
		options.id = id.String()
	}
	for _, o := range opts {
		o(&options)
	}
	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		opts:   options,
		ctx:    ctx,
		cancel: cancel,
	}
}

// Run app
func (a *App) Run() error {
	g, ctx := errgroup.WithContext(a.ctx)
	for _, srv := range a.opts.servers {
		srv := srv
		g.Go(func() error {
			<-ctx.Done() // wait for stop signal
			return srv.Stop()
		})
		g.Go(func() error {
			return srv.Start()
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.opts.sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.Stop()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop gracefully stops the app
func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
