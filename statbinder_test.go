package stat

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/golang/mock/gomock"
	"github.com/rs/xstats"
	"github.com/stretchr/testify/assert"
)

type dummyMessageconsumer struct {
	testFunc func(ctx context.Context)
}

func (t *dummyMessageconsumer) ConsumeMessage(ctx context.Context, message *sqs.Message) error {
	return nil
}

func TestStatBinder_ProcessMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStater := NewMockStat(ctrl)

	consumerMessageFunc := func(ctx context.Context) {
		stater := xstats.FromContext(ctx)
		assert.Equal(t, mockStater, stater)
	}

	statBinder := NewStatBinder(mockStater)
	messageConsumer := statBinder(&dummyMessageconsumer{
		testFunc: consumerMessageFunc,
	})

	messageConsumer.ConsumeMessage(context.Background(), &sqs.Message{})
}
