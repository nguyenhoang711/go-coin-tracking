package global

import (
	database "github.com/nguyenhoang711/go-coin-tracking/database/sqlc"
	"github.com/nguyenhoang711/go-coin-tracking/pkg/logger"
	"github.com/nguyenhoang711/go-coin-tracking/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Db            *database.Store
	Rdb           *redis.Client
	KafkaProducer *kafka.Writer
	KafkaConsumer *kafka.Reader
)
