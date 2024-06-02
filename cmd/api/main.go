package main

import (
	"github.com/HLerman/test/cmd/api/setup"
)

func main() {
	routeur := setup.SetupRouter()
	routeur.Run()
}
