package stat

import (
	"context"
	"strconv"
	"time"

	"github.com/asecurityteam/runsqs/v4"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/rs/xstats"
)

const (
	consumedCounter           = "sqs.consumed"
	consumerSuccessCounter    = "sqs.consumer_success"
	consumerErrorCounter      = "sqs.consumer_error"
	consumedSize              = "sqs.consumed_size"
	consumerLag               = "sqs.consumer_lag.timing"
	consumerTimingSuccess     = "sqs.consumer.timing.success"
	consumerTimingFailure     = "sqs.consumer.timing.failure"
	consumerDeadLetterCounter = "sqs.deadletter"
)

// MessageConsumerConfig is the config for creating a MessageConsumer
type MessageConsumerConfig struct {
	ConsumedCounter           string `description:"Name of overall sqs record consumption metric."`
	ConsumerSuccessCounter    string `description:"Name of overall successful sqs message consumption metric."`
	ConsumerErrorCounter      string `description:"Name of overall failed sqs message consumption metric."`
	ConsumedSize              string `description:"Name of consumed sqs message size metric."`
	ConsumerLag               string `description:"Name of lag time between sqs production and consumption metric."`
	ConsumerTimingSuccess     string `description:"Name of time to process successful sqs message metric."`
	ConsumerTimingFailure     string `description:"Name of time to process failed sqs message metric."`
	ConsumerDeadLetterCounter string `description:"Name of overall dead lettered sqs messages metric."`
}

// Name of the config root.
func (*MessageConsumerConfig) Name() string {
	return "sqsconsumerMetrics"
}

// MessageConsumerComponent implements the settings.Component interface.
type MessageConsumerComponent struct{}

// NewComponent populates default values.
func NewComponent() *MessageConsumerComponent {
	return &MessageConsumerComponent{}
}

// Settings generates a config populated with defaults.
func (*MessageConsumerComponent) Settings() *MessageConsumerConfig {
	return &MessageConsumerConfig{
		ConsumedCounter:           consumedCounter,
		ConsumerSuccessCounter:    consumerSuccessCounter,
		ConsumerErrorCounter:      consumerErrorCounter,
		ConsumedSize:              consumedSize,
		ConsumerLag:               consumerLag,
		ConsumerTimingSuccess:     consumerTimingSuccess,
		ConsumerTimingFailure:     consumerTimingFailure,
		ConsumerDeadLetterCounter: consumerDeadLetterCounter,
	}
}

func (c *MessageConsumerComponent) New(_ context.Context, conf *MessageConsumerConfig) (func(runsqs.SQSMessageConsumer) runsqs.SQSMessageConsumer, error) { // nolint

	return func(consumer runsqs.SQSMessageConsumer) runsqs.SQSMessageConsumer {
		return &MessageConsumer{
			ConsumedCounter:           conf.ConsumedCounter,
			ConsumerSuccessCounter:    conf.ConsumerSuccessCounter,
			ConsumerErrorCounter:      conf.ConsumerErrorCounter,
			ConsumedSize:              conf.ConsumedSize,
			ConsumerLag:               conf.ConsumerLag,
			ConsumerTimingSuccess:     conf.ConsumerTimingSuccess,
			ConsumerTimingFailure:     conf.ConsumerTimingFailure,
			ConsumerDeadLetterCounter: conf.ConsumerDeadLetterCounter,
			wrapped:                   consumer,
		}
	}, nil
}

// MessageConsumer a is wrapper around runsqs.SQSMessageConsumer to capture and emit SQS related stats
type MessageConsumer struct {
	ConsumedCounter           string
	ConsumerSuccessCounter    string
	ConsumerErrorCounter      string
	ConsumedSize              string
	ConsumerLag               string
	ConsumerTimingSuccess     string
	ConsumerTimingFailure     string
	ConsumerDeadLetterCounter string
	wrapped                   runsqs.SQSMessageConsumer
}

// ConsumeMessage pulls an `xstats.XStater` from the context, performs stats around message consumption and invokes the
// wrapped `SQSMessageConsumer.ConsumeMessage`.
func (t MessageConsumer) ConsumeMessage(ctx context.Context, message *types.Message) runsqs.SQSMessageConsumerError {
	stat := xstats.FromContext(ctx)
	// consumerLag - Time.Duration between the time immediately before the message is processed and its Record.ApproximateArrivalTimestamp,
	//which is the timestamp of when the record was inserted into the SQS queue.
	messageArrivalTimeStampInUnix := message.Attributes["SentTimestamp"]

	unixTimeStamp, err := strconv.ParseInt(messageArrivalTimeStampInUnix, 10, 64)
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
	consumeErr := t.wrapped.ConsumeMessage(ctx, message)
	var end = time.Now().Sub(start) // nolint
	if consumeErr == nil {
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
	return consumeErr
}

// DeadLetter pulls an `xstats.XStater` from the context, performs stats around message dead lettering and invokes the
// wrapped `SQSMessageConsumer.DeadLetter`.
func (t MessageConsumer) DeadLetter(ctx context.Context, message *types.Message) {
	stat := xstats.FromContext(ctx)

	// for DeadLetter stating, we just count how many messages are DeadLettered
	stat.Count(t.ConsumerDeadLetterCounter, 1)

	t.wrapped.DeadLetter(ctx, message)
}

// NewStatMessageConsumer returns a function that wraps a `runsqs.SQSMessageConsumer` in a
// `MessageConsumer` `runsqs.SQSMessageConsumer`.
func NewStatMessageConsumer() func(runsqs.SQSMessageConsumer) runsqs.SQSMessageConsumer {
	return func(consumer runsqs.SQSMessageConsumer) runsqs.SQSMessageConsumer {
		return &MessageConsumer{wrapped: consumer}
	}
}
