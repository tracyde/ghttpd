package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"net/http"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghttpd"
	app.Usage = "serve a directory using HTTP"

	app.Flags = []cli.Flag{
		cli.IntFlag{"port, p", 8080, "port ghttpd listens on"},
	}

	app.Action = func(c *cli.Context) {
		// set default values for arguments
		dir := "."

		if len(c.Args()) > 0 {
			dir = c.Args()[0]
		}
		port := fmt.Sprintf(":%d", c.Int("port"))
		log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(dir))))
	}

	app.Run(os.Args)
}
