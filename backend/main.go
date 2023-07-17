package main

import (
	"Minifyr/model"
	"Minifyr/server"
)

func main() {
	model.Setup()
	server.SetupServerListener()
}
