package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/sudomer/boiler-fiber/internal/server"
	"github.com/sudomer/boiler-fiber/pkg/lib"
)

func main() {

	lib.Log().Info("Service started.")

	server.NewServer()
}
