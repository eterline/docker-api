package main

import (
	"flag"
	"log"

	"github.com/eterline/docker-api/internal/api"
)

type Settings struct {
	Port uint
	Ip   string
}

func main() {
	var sets Settings
	sets.mustArgs()
	api.StartServer(sets.Port, sets.Ip)
}

func (s *Settings) mustArgs() {
	flag.UintVar(&s.Port, "port", 9000, "Set port")
	flag.StringVar(&s.Ip, "ip", "0.0.0.0", "Set ip")
	flag.Parse()
	if s.Port > 65535 {
		log.Fatal("Incorrect port.")
	}
	log.Printf("Server args set.")
}
