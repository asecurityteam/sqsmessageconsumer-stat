package stat

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/rs/xstats"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestStatMessageConsumer_ConsumeMessageSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMessageConsumer := NewMockSQSMessageConsumer(ctrl)
	mockStater := NewMockStat(ctrl)

	statMessageConsumer := StatMessageConsumer{
		ConsumedCounter:        consumedCounter,
		ConsumerSuccessCounter: consumerSuccessCounter,
		ConsumedSize:           consumedSize,
		ConsumerLag:            consumerLag,
		ConsumerTimingSuccess:  consumerTimingSuccess,
		wrapped:                mockMessageConsumer,
	}

	var incomingDataRecord string = "data"

	// random unix time
	currentTime := "1602014628"
	sqsMessage := sqs.Message{
		Body: &incomingDataRecord,
		Attributes: map[string]*string{
			"SentTimestamp": &currentTime,
		},
	}
	gomock.InOrder(
		mockStater.EXPECT().Timing(consumerLag, gomock.Any()),
		mockStater.EXPECT().Count(consumedCounter, gomock.Any()),
		mockStater.EXPECT().Count(consumedSize, gomock.Any()),
		mockStater.EXPECT().Count(consumerSuccessCounter, gomock.Any()),
		mockStater.EXPECT().Timing(consumerTimingSuccess, gomock.Any()),
	)

	mockMessageConsumer.EXPECT().ConsumeMessage(gomock.Any(), gomock.Any()).Return(nil)
	e := statMessageConsumer.ConsumeMessage(xstats.NewContext(context.Background(), mockStater), &sqsMessage)
	assert.Nil(t, e)

}

func TestStatMessageConsumer_ConsumeMessageFailure(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMessageConsumer := NewMockSQSMessageConsumer(ctrl)
	mockStater := NewMockStat(ctrl)

	statMessageConsumer := StatMessageConsumer{
		ConsumedCounter:       consumedCounter,
		ConsumerErrorCounter:  consumerErrorCounter,
		ConsumedSize:          consumedSize,
		ConsumerLag:           consumerLag,
		ConsumerTimingFailure: consumerTimingFailure,
		wrapped:               mockMessageConsumer,
	}

	var incomingDataRecord string = "data"

	// random unix time
	currentTime := "1602014628"
	sqsMessage := sqs.Message{
		Body: &incomingDataRecord,
		Attributes: map[string]*string{
			"SentTimestamp": &currentTime,
		},
	}
	gomock.InOrder(
		mockStater.EXPECT().Timing(consumerLag, gomock.Any()),
		mockStater.EXPECT().Count(consumedCounter, gomock.Any()),
		mockStater.EXPECT().Count(consumedSize, gomock.Any()),
		mockStater.EXPECT().Count(consumerErrorCounter, gomock.Any()),
		mockStater.EXPECT().Timing(consumerTimingFailure, gomock.Any()),
	)

	mockMessageConsumer.EXPECT().ConsumeMessage(gomock.Any(), gomock.Any()).Return(errors.New("consume error"))
	e := statMessageConsumer.ConsumeMessage(xstats.NewContext(context.Background(), mockStater), &sqsMessage)
	assert.NotNil(t, e)

}

func TestNewStatMessageConsumerConfig(t *testing.T) {
	component := NewComponent()
	config := component.Settings()
	statConsumer, err := component.New(context.Background(), config)
	assert.NotNil(t, statConsumer)
	assert.Nil(t, err)

}

func TestStatMessageConsumerConfig_Name(t *testing.T) {
	component := NewComponent()
	config := component.Settings()
	assert.Equal(t, config.Name(), "sqsconsumerMetrics")
}
