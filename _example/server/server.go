package main

import (
	"context"
	"fmt"
	"log"

	p2p "github.com/leprosus/golang-p2p"
)

type Hello struct {
	Text string
}

type Buy struct {
	Text string
}

func main() {
	tcp := p2p.NewTCP("localhost", 8080)

	settings := p2p.NewServerSettings()

	server, err := p2p.NewServer(tcp, settings)
	if err != nil {
		log.Panicln(err)
	}

	server.SetHandle("dialog", func(ctx context.Context, req p2p.Request) (res p2p.Response, err error) {
		hello := Hello{}
		err = req.GetGob(&hello)
		if err != nil {
			return
		}

		fmt.Printf("> Hello: %s\n", hello.Text)

		buy := Buy{Text: hello.Text}
		err = res.SetGob(buy)

		return
	})

	err = server.Serve()
	if err != nil {
		log.Panicln(err)
	}
}
