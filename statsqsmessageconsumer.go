package stat

import (
	"context"
	"strconv"
	"time"

	"github.com/asecurityteam/runsqs"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/rs/xstats"
)

const (
	consumedCounter        = "sqs.consumed"
	consumerSuccessCounter = "sqs.consumer_success"
	consumerErrorCounter   = "sqs.consumer_error"
	consumedSize           = "sqs.consumed_size"
	consumerLag            = "sqs.consumer_lag.timing"
	consumerTimingSuccess  = "sqs.consumer.timing.success"
	consumerTimingFailure  = "sqs.consumer.timing.failure"
)

// StatMessageConsumerConfig is the config for creating a StatMessageConsumer
type StatMessageConsumerConfig struct {
	ConsumedCounter        string `description:"Name of overall sqs record consumption metric."`
	ConsumerSuccessCounter string `description:"Name of overall successful sqs message consumption metric."`
	ConsumerErrorCounter   string `description:"Name of overall failed sqs message consumption metric."`
	ConsumedSize           string `description:"Name of consumed sqs message size metric."`
	ConsumerLag            string `description:"Name of lag time between sqs production and consumption metric."`
	ConsumerTimingSuccess  string `description:"Name of time to process successful sqs message metric."`
	ConsumerTimingFailure  string `description:"Name of time to process failed sqs message metric."`
}

// Name of the config root.
func (*StatMessageConsumerConfig) Name() string {
	return "sqsconsumerMetrics"
}

// StatMessageConsumerComponent implements the settings.Component interface.
type StatMessageConsumerComponent struct{}

// NewComponent populates default values.
func NewComponent() *StatMessageConsumerComponent {
	return &StatMessageConsumerComponent{}
}

// Settings generates a config populated with defaults.
func (*StatMessageConsumerComponent) Settings() *StatMessageConsumerConfig {
	return &StatMessageConsumerConfig{
		ConsumedCounter:        consumedCounter,
		ConsumerSuccessCounter: consumerSuccessCounter,
		ConsumerErrorCounter:   consumerErrorCounter,
		ConsumedSize:           consumedSize,
		ConsumerLag:            consumerLag,
		ConsumerTimingSuccess:  consumerTimingSuccess,
		ConsumerTimingFailure:  consumerTimingFailure,
	}
}

func (c *StatMessageConsumerComponent) New(_ context.Context, conf *StatMessageConsumerConfig) (func(runsqs.SQSMessageConsumer) runsqs.SQSMessageConsumer, error) { // nolint

	return func(consumer runsqs.SQSMessageConsumer) runsqs.SQSMessageConsumer {
		return &StatMessageConsumer{
			ConsumedCounter:        conf.ConsumedCounter,
			ConsumerSuccessCounter: conf.ConsumerSuccessCounter,
			ConsumerErrorCounter:   conf.ConsumerErrorCounter,
			ConsumedSize:           conf.ConsumedSize,
			ConsumerLag:            conf.ConsumerLag,
			ConsumerTimingSuccess:  conf.ConsumerTimingSuccess,
			ConsumerTimingFailure:  conf.ConsumerTimingFailure,
			wrapped:                consumer,
		}
	}, nil
}

// StatMessageConsumer a is wrapper around runsqs.SQSMessageConsumer to capture and emit SQS related stats
type StatMessageConsumer struct {
	ConsumedCounter        string
	ConsumerSuccessCounter string
	ConsumerErrorCounter   string
	ConsumedSize           string
	ConsumerLag            string
	ConsumerTimingSuccess  string
	ConsumerTimingFailure  string
	wrapped                runsqs.SQSMessageConsumer
}

// ConsumeMessage pulls an `xstats.XStater` from the context, performs stats around message consumption and invokes the
// wrapped `SQSMessageConsumer.ConsumeMessage`.
func (t StatMessageConsumer) ConsumeMessage(ctx context.Context, message *sqs.Message) runsqs.SQSMessageConsumerError {
	stat := xstats.FromContext(ctx)
	// consumerLag - Time.Duration between the time immediately before the message is processed and its Record.ApproximateArrivalTimestamp,
	//which is the timestamp of when the record was inserted into the SQS queue.
	messageArrivalTimeStampInUnix := message.Attributes["SentTimestamp"]

	unixTimeStamp, err := strconv.ParseInt(*messageArrivalTimeStampInUnix, 10, 64)
	if err != nil {
		panic(err)
	}
	messageArrivalTimeStamp := time.Unix(0, unixTimeStamp*int64(time.Millisecond)).Local()

	lagDuration := time.Since(messageArrivalTimeStamp)
	stat.Timing(t.ConsumerLag, lagDuration)

	// consumedCounter - Incremented for every message received, regardless of success or failure
	stat.Count(t.ConsumedCounter, 1)

	// length of message body
	stat.Count(t.ConsumedSize, float64(len([]byte(*message.Body))))

	var start = time.Now()
	consumeerr := t.wrapped.ConsumeMessage(ctx, message)
	var end = time.Now().Sub(start)
	if consumeerr == nil {
		// consumerSuccessCounter - Incremented for every message processed successfully
		stat.Count(t.ConsumerSuccessCounter, 1)
		// consumerTimingSuccess - Time.Duration for processing of a message which is successfully processed by underlying SQS consumer
		stat.Timing(t.ConsumerTimingSuccess, end)
	} else {
		// consumerFailure - Incremented for every message which is failed to be processed
		stat.Count(t.ConsumerErrorCounter, 1)
		// consumerTimingFailure - Time.Duration for processing of a message which underlying SQS consumer fails to process
		stat.Timing(t.ConsumerTimingFailure, end)
	}
	return consumeerr
}

// NewStatMessageConsumer returns a function that wraps a `runsqs.SQSMessageConsumer` in a
// `StatMessageConsumer` `runsqs.SQSMessageConsumer`.
func NewStatMessageConsumer() func(runsqs.SQSMessageConsumer) runsqs.SQSMessageConsumer {
	return func(consumer runsqs.SQSMessageConsumer) runsqs.SQSMessageConsumer {
		return &StatMessageConsumer{wrapped: consumer}
	}
}
