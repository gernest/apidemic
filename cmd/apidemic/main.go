package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/cli"
	"github.com/gernest/apidemic"
)

func server(ctx *cli.Context) {
	port := ctx.Int("port")
	s := apidemic.NewServer()

	log.Println("starting server on port :", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), s))
}

func main() {
	app := cli.NewApp()
	app.Name = "apidemic"
	app.Usage = "Fake JSON API Responses"
	app.Authors = []cli.Author{
		{"Geofrey Ernest", "geofreyernest@live.com"},
	}
	app.Version = apidemic.Version
	app.Commands = []cli.Command{
		cli.Command{
			Name:      "start",
			ShortName: "s",
			Usage:     "starts apidemic server",
			Action:    server,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:   "port",
					Usage:  "http port to run",
					Value:  3000,
					EnvVar: "PORT",
				},
			},
		},
	}
	app.RunAndExitOnError()
}
