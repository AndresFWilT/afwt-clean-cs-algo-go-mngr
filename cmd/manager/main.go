package main

import "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/config"

func main() {
	serverConfigured, port := config.InitializeServer()
	err := serverConfigured.Run(":" + port)
	if err != nil {
		return
	}
}
