package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/filecloud/config"
	"github.com/huoxue1/filecloud/model"
	_ "github.com/huoxue1/filecloud/model"
	"github.com/huoxue1/filecloud/router" //nolint:gci
	"github.com/huoxue1/filecloud/utils"
)

func main() {
	config.CheckConfigFile("config.yml")

	err := config.LoadConfig("config.yml")
	if err != nil {
		log.Panicln("初始化配置文件错误" + err.Error())
		return
	}
	utils.CheckDir(config.GetConfig().StoragePath)
	utils.CheckFile(config.GetConfig().DBPath)
	model.Init()
	model.InitRoot()
	engine := router.Router()
	if err := engine.Run(fmt.Sprintf("%s:%d", config.GetConfig().Address, config.GetConfig().Port)); err != nil {
		log.Panicln(err.Error())
	}
}
