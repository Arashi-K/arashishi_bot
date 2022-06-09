package main

import (
	"arashishi_bot/handler"
)

const (
	selectVersionAction     = "select-version"
	confirmDeploymentAction = "confirm-deployment"
)

func main() {
	handler.RunHandlers()
}
