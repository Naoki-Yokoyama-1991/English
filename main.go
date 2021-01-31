package main

import (
	"github.com/naoki/gacha/app/api"
)

var server = api.Server{}

func main() {

	server.Run()
}
