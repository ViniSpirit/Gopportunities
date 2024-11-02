package main

import (
	"github.com/ViniSpirit/Gopportunities/config"
	"github.com/ViniSpirit/Gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	//Initialize configs
	logger = config.GetLogger("main")

	err := config.InitConfig()

	if err != nil {
		logger.Errorf("Config inicialization erro: %v", err)
		return
	}

	router.Inicialize()
}
