package main

import (
	"net/http"
	_ "net/http/pprof"

	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/filecloud/config"
	"github.com/huoxue1/filecloud/model"
	"github.com/huoxue1/filecloud/router" //nolint:gci
	"github.com/huoxue1/filecloud/utils"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

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
	http.ListenAndServe("0.0.0.0:8081", engine)
	//if err := engine.Run(fmt.Sprintf("%s:%d", config.GetConfig().Address, config.GetConfig().Port)); err != nil {
	//	log.Panicln(err.Error())
	//}
}
