package main

import (
	"Project/configs"
	"Project/routes"
	"fmt"
	_ "image/png" // Import image packages to support PNG, JPEG, etc.
)

func main() {

	configs.InitiEnvConfigs()
	routes.InitializeRoutes()
	routes.Router.Run(":" + configs.EnvConfigs.LocalServerPort)

	fmt.Printf("hahahha")
	fmt.Printf("hahahha")
	fmt.Printf("hahahha")
}
