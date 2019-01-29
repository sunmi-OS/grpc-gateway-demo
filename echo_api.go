package main

import (
	"os"
	"sort"
	"strings"
	"net/http"
	"io/ioutil"

	"github.com/urfave/cli"
	"github.com/labstack/echo"
	"github.com/sunmi-OS/gocore/api"
)

type EchoApi struct {
}

var eApi EchoApi

func (a *EchoApi) echoStart(c *cli.Context) error {
	// Echo instance
	e := echo.New()

	// Route => handler
	e.GET("/", func(c echo.Context) error {

		//request := api.NewRequest(c)
		response := api.NewResponse(c)

		resp, err := http.Post("http://localhost:8080/v1/example/echo", "application/x-www-form-urlencoded",
			strings.NewReader(`{"value":" world"}`))
		if err != nil {
			// handle error
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}

		return response.RetSuccess(string(body))
	})

	// Start server
	e.Logger.Fatal(e.Start(":1333"))
	return nil
}

func main() {

	app := cli.NewApp()

	// 指定对于的命令
	app.Commands = []cli.Command{
		{
			Name:    "api",
			Aliases: []string{"a"},
			Usage:   "api",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "开启ECHO-API",
					Action: eApi.echoStart,
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
