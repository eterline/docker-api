package main

import (
	"github.com/eterline/docker-api/internal/api"
)

func main() {
	api.Sets.MustArgs()
	api.StartServer(api.Sets.Port, api.Sets.Ip)
}
