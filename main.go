package main

import (
	"github.com/Bertin203/gopportunities/config"
	"github.com/Bertin203/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	// initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize() // initialize Router
}
