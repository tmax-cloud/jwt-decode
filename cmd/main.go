package main

import (
	"github.com/tmax-cloud/jwt-decode/config"
)

func main() {
	c, _ := config.NewConfig().RunServer()
	panic(<-c)
}
