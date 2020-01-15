package lib

import (
	"context"

	"github.com/google/uuid"
	"github.com/linode/linodego"
)

// GenPassword generates a random password
func GenPassword() string {
	return uuid.New().String()
}

// ResetPassword resets the root password of a linode
func ResetPassword(c context.Context, linodeClient *linodego.Client, linodeID int) (string, error) {
	disks, err := linodeClient.ListInstanceDisks(c, linodeID, nil)
	if err != nil {
		return "", err
	}
	if len(disks) == 0 {
		return "", nil
	}
	password := GenPassword()
	err = linodeClient.PasswordResetInstanceDisk(c, linodeID, disks[0].ID, password)
	if err != nil {
		return "", err
	}
	return password, nil
}
