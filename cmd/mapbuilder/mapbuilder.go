package main

import (
	"flag"
	"github.com/AndreyGermanov/mapbuilder_backend/internal/mapbuilder"
)

var p = flag.Int("p", 6001, "Port number")

func main() {
	flag.Parse()
	mapbuilder.RunWebServer(*p)
}
