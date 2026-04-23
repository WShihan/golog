package handler

import (
	"fmt"

	"golog/entity"
	"golog/util"

	"github.com/urfave/cli/v2"
)

var (
	injection entity.Injection
)

func Start(c *cli.Context, inject *entity.Injection) error {
	injection = *inject
	if c.String("tls-crt") != "" && c.String("tls-key") != "" {
		url := fmt.Sprintf("https://localhost%s", c.String("port"))
		fmt.Printf("👋 Visitf %s to use Golog\n", url)
		util.OpenBrowser(url)
		return Router.RunTLS(c.String("port"), c.String("tls-crt"), c.String("tls-key"))
	} else {
		url := fmt.Sprintf("http://localhost%s", c.String("port"))
		fmt.Printf("👋 Visitf %s to use Golog\n", url)
		util.OpenBrowser(url)
		return Router.Run(c.String("port"))
	}
}
