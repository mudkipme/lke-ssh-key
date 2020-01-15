package lib

import (
	"context"
	"fmt"
	"time"

	"github.com/linode/linodego"
)

// PowerOff shuts down a linode
func PowerOff(c context.Context, linodeClient *linodego.Client, linodeID int) error {
	i, err := linodeClient.GetInstance(c, linodeID)
	if err != nil {
		return err
	}
	if i.Status == linodego.InstanceOffline {
		return nil
	}
	err = linodeClient.ShutdownInstance(c, i.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Waiting for %v shutting down.", i.Label)
	for {
		time.Sleep(5 * time.Second)
		instance, err := linodeClient.GetInstance(c, i.ID)
		if err != nil {
			return err
		}
		if instance.Status == linodego.InstanceOffline {
			fmt.Printf("\nShut down %v.\n", i.Label)
			return nil
		}
		fmt.Print(".")
	}
}

// PowerOn boots up a linode
func PowerOn(c context.Context, linodeClient *linodego.Client, linodeID int) error {
	i, err := linodeClient.GetInstance(c, linodeID)
	if err != nil {
		return err
	}
	if i.Status == linodego.InstanceRunning {
		return nil
	}
	config, err := linodeClient.ListInstanceConfigs(c, i.ID, nil)
	if err != nil {
		return err
	}
	if len(config) == 0 {
		return nil
	}
	err = linodeClient.BootInstance(c, i.ID, config[0].ID)
	if err != nil {
		return err
	}
	fmt.Printf("Waiting for %v booting up.", i.Label)
	for {
		time.Sleep(5 * time.Second)
		instance, err := linodeClient.GetInstance(c, i.ID)
		if err != nil {
			return err
		}
		if instance.Status == linodego.InstanceRunning {
			fmt.Printf("\nBooted up %v.\n", i.Label)
			return nil
		}
		fmt.Print(".")
	}
}
