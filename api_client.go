package main

import (
	"os"
	"log"
	"sort"

	pb "grpc-gateway-demo/gateway"

	"github.com/urfave/cli"
	"github.com/labstack/echo"
	"github.com/sunmi-OS/gocore/api"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

type EchoApi struct {
}

var eApi EchoApi

const (
	address     = "localhost:9192"
	defaultName = "world"
)

func (a *EchoApi) echoStart(c *cli.Context) error {
	// Echo instance
	e := echo.New()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	pbc := pb.NewGatewayClient(conn)

	// Route => handler
	e.GET("/", func(c echo.Context) error {

		response := api.NewResponse(c)

		r, err := pbc.Echo(context.Background(), &pb.StringMessage{Value: defaultName})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Value)

		return response.RetSuccess(r.Value)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}

func main() {

	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "api",
			Aliases: []string{"a"},
			Usage:   "api",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "开启API-Client",
					Action: eApi.echoStart,
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
