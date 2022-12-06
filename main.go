package main

import (
	"gin/controller"
	"gin/initial"
	"log"
	"os"
)

func main() {
	initial.LoadEnv()
	port := os.Getenv("PORT")
	db := initial.ConnectDB()
	services := controller.Service(db)
	controller.MainRouter(services, port)
	log.Println("server starts on port", port)
}
