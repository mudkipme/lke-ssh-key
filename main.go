package main

import (
	"log"
	"os"

	"github.com/mudkipme/lke-ssh-key/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "lke-ssh-key",
		Usage:  "Add SSH keys in your Linode account to your Linode Kubernetes Engine nodes",
		Action: commands.LKESSHKey,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
