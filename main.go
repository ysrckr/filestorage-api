package main

import (
	"github.com/ysrckr/filestorage/api/config"
)




func main() {
	config.Config.StartServer(config.PORT)
}