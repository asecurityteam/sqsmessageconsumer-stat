package stat

import (
	"context"

	"github.com/asecurityteam/runsqs/v3"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/rs/xstats"
)

// Binder is a `SQSMessageConsumer` decorator that injects an
// `xstats.XStater` into the `context.Context` of the given `context`.
//
// The `Binder` can then, in turn, be decorated with a `SQSMessageConsumer` to
// make use of the injected `xstats.XStater` to emit key HTTP metrics on
// each v.
type Binder struct {
	stats   xstats.XStater
	wrapped runsqs.SQSMessageConsumer
}

// ConsumeMessage injects an `xstats.XStater` into the context and invokes the
// wrapped `SQSMessageConsumer`.
func (t *Binder) ConsumeMessage(ctx context.Context, message *sqs.Message) runsqs.SQSMessageConsumerError {
	ctx = xstats.NewContext(ctx, xstats.Copy(t.stats))
	return t.wrapped.ConsumeMessage(ctx, message)
}

// DeadLetter injects an `xstats.XStater` into the context and invokes the
// wrapped `SQSMessageConsumer`.
func (t *Binder) DeadLetter(ctx context.Context, message *sqs.Message) {
	ctx = xstats.NewContext(ctx, xstats.Copy(t.stats))
	t.wrapped.DeadLetter(ctx, message)
}

// NewStatBinder returns a function that wraps a `runsqs.SQSMessageConsumer` in a
// `Binder` `runsqs.SQSMessageConsumer`.
func NewStatBinder(stats xstats.XStater) func(runsqs.SQSMessageConsumer) runsqs.SQSMessageConsumer {
	return func(next runsqs.SQSMessageConsumer) runsqs.SQSMessageConsumer {
		return &Binder{stats: stats, wrapped: next}
	}
}
