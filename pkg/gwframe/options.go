package gwframe

import (
	"context"
	"os"

	"github.com/JohnnyWei188/gwframe/pkg/transport"
)

// Option is an application option.
type Option func(o *options)

// options is an application options.
type options struct {
	id        string
	endpoints []string

	ctx     context.Context
	sigs    []os.Signal
	servers []transport.Server
}

// ID with service id.
func ID(id string) Option {
	return func(o *options) { o.id = id }
}

// Context with service context.
func Context(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// Signal with exit signals.
func Signal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

// WithServer with app server
func WithServer(s ...transport.Server) Option {
	return func(o *options) { o.servers = s }
}
