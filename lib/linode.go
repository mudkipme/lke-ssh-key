package lib

import (
	"errors"
	"net/http"
	"os"
	"sync"

	"github.com/linode/linodego"
	"golang.org/x/oauth2"
)

var (
	instance     *linodego.Client
	mutex        sync.Mutex
)

// GetClient returns a singleton linode client
func GetClient() (*linodego.Client, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if instance != nil {
		return instance, nil
	}
	apiKey, ok := os.LookupEnv("LINODE_TOKEN")
	if !ok {
		return nil, errors.New("could not find LINODE_TOKEN, please assert it is set")
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey})

	oauth2Client := &http.Client{
		Transport: &oauth2.Transport{
			Source: tokenSource,
		},
	}
	client := linodego.NewClient(oauth2Client)
	instance = &client
	return instance, nil
}
