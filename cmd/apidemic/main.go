package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/codegangsta/cli"
	"github.com/gernest/apidemic"
	"github.com/gorilla/mux"
)

func server(ctx *cli.Context) {
	port := ctx.Int("port")
	endpointDir := ctx.String("endpoint-dir")
	s := apidemic.NewServer()

	err := addEndpoints(s, endpointDir)
	if err != nil {
		log.Fatalln(err)
	}

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
					Usage:  "HTTP port to run",
					Value:  3000,
					EnvVar: "PORT",
				},
				cli.StringFlag{
					Name:   "endpoint-dir",
					Usage:  "Directory to scan for endpoint",
					Value:  "./endpoints/",
					EnvVar: "ENDPOINT_DIR",
				},
			},
		},
	}
	app.RunAndExitOnError()
}

// addEndpoints iterates over a directory and tries to register each file as an endpoint
// if directory does not exist or not readable it will simply return without registration
// failing registration however causes an error response
func addEndpoints(s *mux.Router, endpointDir string) error {
	var registerPayload map[string]interface{}
	files, err := ioutil.ReadDir(endpointDir)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := endpointDir + file.Name()

		registerPayload, err = getRegisterPayload(filePath)
		if err != nil {
			return err
		}

		w := httptest.NewRecorder()
		req := apidemic.JsonRequest("POST", "/register", registerPayload)
		s.ServeHTTP(w, req)

		log.Printf("%s is registered\n", filePath)

		if w.Code != http.StatusOK {
			return errors.New("registering " + filePath + " failed")
		}
	}

	return nil
}

func getRegisterPayload(endpoint string) (map[string]interface{}, error) {
	var api map[string]interface{}

	content, err := ioutil.ReadFile(endpoint)
	if err != nil {
		return api, err
	}

	err = json.NewDecoder(bytes.NewReader(content)).Decode(&api)
	if err != nil {
		return api, err
	}

	return api, nil
}
