package main

import (
	"github.com/marvelalexius/gymbroapp/config"
	"github.com/marvelalexius/gymbroapp/internal/app"
)

func main() {
	cfg := config.GetInstance()
	app.Run(cfg)
}
