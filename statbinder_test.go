package stat

import (
	"context"
	"testing"

	runsqs "github.com/asecurityteam/runsqs/v4"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/rs/xstats"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type dummyMessageconsumer struct {
	testFunc func(ctx context.Context)
}

func (t *dummyMessageconsumer) ConsumeMessage(ctx context.Context, message *types.Message) runsqs.SQSMessageConsumerError {
	return nil
}

func (t *dummyMessageconsumer) DeadLetter(ctx context.Context, message *types.Message) {}

func TestStatBinder_ProcessMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStater := NewMockXStater(ctrl)

	consumerMessageFunc := func(ctx context.Context) {
		stater := xstats.FromContext(ctx)
		assert.Equal(t, mockStater, stater)
	}

	statBinder := NewStatBinder(mockStater)
	messageConsumer := statBinder(&dummyMessageconsumer{
		testFunc: consumerMessageFunc,
	})

	messageConsumer.ConsumeMessage(context.Background(), &types.Message{}) // nolint
}
