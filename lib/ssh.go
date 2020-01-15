package lib

import (
	"fmt"
	"time"
	"os"

	"github.com/linode/linodego"
	"golang.org/x/crypto/ssh"
)

// SetAuthorizedKeys sets authorized_keys to a list of ssh keys
func SetAuthorizedKeys(ip string, password string, keys []linodego.SSHKey) error {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: 5 * time.Second,
	}

	conn, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		return err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stdin, err := session.StdinPipe()
	if err != nil {
		return err
	}
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.Shell()
	if err != nil {
		return err
	}

	fmt.Fprintln(stdin, "mkdir -p ~/.ssh")
	fmt.Fprintln(stdin, "rm -f ~/.ssh/authorized_keys")
	fmt.Fprintln(stdin, "touch ~/.ssh/authorized_keys")

	for _, key := range keys {
		fmt.Fprintf(stdin, "echo '%v' >> ~/.ssh/authorized_keys\n", key.SSHKey)
	}

	fmt.Fprintln(stdin, "exit")
	err = session.Wait()
	if err != nil {
		return err
	}
	return nil
}
