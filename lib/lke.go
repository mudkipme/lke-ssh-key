package lib

import (
	"context"
	"strings"

	"github.com/linode/linodego"
)

// LKEInstances lists linodes created by LKE
func LKEInstances(c context.Context, linodeClient *linodego.Client) ([]linodego.Instance, error) {
	instances, err := linodeClient.ListInstances(c, nil)
	if err != nil {
		return nil, err
	}

	// As linodego doesn't support LKE API yet, we grep linodes with "lke" prefix and kubernetes images
	var lkeInstances []linodego.Instance
	for _, i := range instances {
		if strings.HasPrefix(i.Label, "lke") && strings.Contains(i.Image, "kube") {
			lkeInstances = append(lkeInstances, i)
		}
	}
	return lkeInstances, nil
}
