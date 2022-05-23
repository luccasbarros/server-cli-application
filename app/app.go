package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executado
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca ips e nomes de servidor"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{{
		Name:   "ip",
		Usage:  "Busca ips de endereços na internet",
		Flags:  flags,
		Action: getIps,
	},
		{
			Name:   "servidores",
			Usage:  "Busca servidores pelo endereço web",
			Flags:  flags,
			Action: getServers,
		},
	}

	return app
}

func getIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
