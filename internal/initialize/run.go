package initialize

import (
	"fmt"

	"github.com/nguyenhoang711/go-coin-tracking/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()

	r := InitRouter()

	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	if global.Config.Server.Mode != "release" {
		fmt.Println(serverAddr)
	}

	r.Run(serverAddr)
}
