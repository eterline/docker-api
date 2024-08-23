package settings

import (
	"flag"
	"log"
)

type Settings struct {
	Port uint
	Ip   string
	Pass string
}

func (s *Settings) MustArgs() {
	flag.UintVar(&s.Port, "port", 9000, "Set port")
	flag.StringVar(&s.Ip, "ip", "0.0.0.0", "Set ip")
	flag.StringVar(&s.Pass, "pass", "1234", "Set pass")
	flag.Parse()
	if s.Port > 65535 {
		log.Fatal("Incorrect port.")
	}
	log.Printf("Server args set.")
}
