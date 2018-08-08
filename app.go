package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

// NewApp returns a new cli
func NewApp() *cli.App {
	var (
		swarmHost string
		swarmPort string
		swarmHash string
		localPort string
	)
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "swarm-port",
			Value:       "8500",
			Usage:       "swarm port",
			Destination: &swarmPort,
		},
		cli.StringFlag{
			Name:        "swarm-ip",
			Value:       "localhost",
			Usage:       "swarm ip address",
			Destination: &swarmHost,
		},
		cli.StringFlag{
			Name:        "swarm-hash",
			Usage:       "swarm hash",
			Destination: &swarmHash,
		},
		cli.StringFlag{
			Name:        "local-port",
			Value:       "3000",
			Usage:       "local http port",
			Destination: &localPort,
		},
	}

	app.Action = func(c *cli.Context) error {
		if swarmHash == "" {
			cli.ShowAppHelp(c)
			return cli.NewExitError("", -1)
		}

		http.HandleFunc("/", SwarmHandler(swarmHost, swarmPort, swarmHash))
		fmt.Println("Open http://localhost:" + localPort)
		return http.ListenAndServe(":"+localPort, nil)
	}
	return app
}
