package consumer

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenhoang711/go-coin-tracking/global"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, groupID string, topics []string) *Consumer {
	consumer := &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     brokers,
			GroupID:     groupID,
			GroupTopics: topics,
		}),
	}
	return consumer
}

func (c *Consumer) Start(ctx *gin.Context, handler func(topic string, message []byte) error) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			global.Logger.Error("error reading messages ", zap.Error(err))
			return err
		}
		global.Logger.Info("message received ",
			zap.String("topic", msg.Topic),
			zap.String("msgKey", string(msg.Key)),
			zap.String("value", string(msg.Value)),
		)
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
