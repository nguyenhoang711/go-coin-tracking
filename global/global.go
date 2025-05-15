package global

import (
	database "github.com/nguyenhoang711/go-coin-tracking/database/sqlc"
	"github.com/nguyenhoang711/go-coin-tracking/pkg/logger"
	"github.com/nguyenhoang711/go-coin-tracking/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Db     *database.Store
)
