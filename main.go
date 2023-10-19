package main

import "github.com/ysrckr/filestorage/api/config"

var PORT = "8080"

func main() {
	config.Config.StartServer(PORT)
}