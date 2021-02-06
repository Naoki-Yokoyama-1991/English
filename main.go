package main

import (
	"github.com/naoki/task/app/api"
)

var server = api.Server{}

func main() {

	server.Run()
}
