package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

// Generate will return the command line ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "search for IPs and server names on the internet"

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "amazon.com.br",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:   "ip",
			Usage:  "Search ip from internet",
			Flags:  flags,
			Action: SearchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search name severs from internet",
			Flags:  flags,
			Action: SearchServers,
		},
	}

	return app
}

func SearchIps(c *cli.Context) error {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil

}

func SearchServers(c *cli.Context) error {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

	return nil

}

