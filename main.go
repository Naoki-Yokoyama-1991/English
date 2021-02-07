package main

import (
	"github.com/naoki/english/app/api"
)

var server = api.Server{}

func main() {

	server.Run()
}
