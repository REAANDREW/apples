package main

import (
	"os"
	"os/signal"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var address = kingpin.Flag("address", "Address to listen on").Short('a').Required().Default("45000").String()

func main() {
	kingpin.Parse()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	CreateApplication(*address).Run(c)
}
