package cmd

import (
	"golang.org/x/crypto/ssh"
	"bytes"
	"fmt"
	"log"
)

func Start_vpn(client *ssh.Client) (string, error) {
	session, err := client.NewSession()
	var out, in bytes.Buffer
	session.Stdout = &out
	session.Stdin = &in
	fmt.Println(CommandMaps[StartVPN])
	err = session.Run(CommandMaps[StartVPN])
	if err != nil {
		fmt.Println("err in runing",err)
		log.Fatal(err)
		return "", err
	}
	return out.String(), nil
}