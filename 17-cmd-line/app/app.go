package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Build() will return the command line ready to run
func Build() *cli.App {
	app := cli.NewApp()

	app.Name = "Go CLI"
	app.Usage = "Get IPs and Hostnames"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "get IPs on the internet",
			Flags: flags,
			Action: getIPs,
		},
		{
			Name: "servers",
			Usage: "get server hostname on the internet",
			Flags: flags,
			Action: getServers,
		},
	}

	return app
}

func getIPs(c *cli.Context) {
	host := c.String("host")
	ips, erro :=  net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServers(c *cli.Context) {
	host := c.String("host")
	servers, erro :=  net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}