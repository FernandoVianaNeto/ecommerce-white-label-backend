package main

import (
	"ecommerce-white-label-backend/cmd/cli"
	configs "ecommerce-white-label-backend/cmd/config"
)

func main() {
	configs.InitializeConfigs()

	cli.Execute()
}
