package initialize

import (
	"github.com/nguyenhoang711/go-coin-tracking/global"
	"github.com/nguyenhoang711/go-coin-tracking/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
