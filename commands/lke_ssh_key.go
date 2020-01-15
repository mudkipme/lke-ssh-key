package commands

import (
	"fmt"
	"time"

	"github.com/mudkipme/lke-ssh-key/lib"
	"github.com/urfave/cli/v2"
)

// LKESSHKey adds SSH keys in Linode account to Linode Kubernetes Engine nodes
func LKESSHKey(c *cli.Context) error {
	linodeClient, err := lib.GetClient()
	if err != nil {
		return err
	}
	lkeInstances, err := lib.LKEInstances(c.Context, linodeClient)
	if err != nil {
		return err
	}

	sshKeys, err := linodeClient.ListSSHKeys(c.Context, nil)
	if err != nil {
		return err
	}

	for _, i := range lkeInstances {
		err := lib.PowerOff(c.Context, linodeClient, i.ID)
		if err != nil {
			return err
		}
		password, err := lib.ResetPassword(c.Context, linodeClient, i.ID)
		if err != nil {
			return err
		}
		err = lib.PowerOn(c.Context, linodeClient, i.ID)
		if err != nil {
			return err
		}
		retry := 0
		fmt.Printf("Waiting for %v ready.", i.Label)
		for {
			err = lib.SetAuthorizedKeys(i.IPv4[0].String(), password, sshKeys)
			if err == nil {
				fmt.Printf("\nssh root@%v\nPassword: %v\n", i.IPv4[0].String(), password)
				fmt.Printf("Configured authorized keys for %v\n", i.Label)
				break
			}
			retry++
			if retry == 10 {
				return err
			}
			fmt.Print(".")
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}
