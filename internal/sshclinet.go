package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func ConnectToServer() *ssh.Client {
	err := godotenv.Load(".env")
	keypath := os.Getenv("KEY_PATH")
	IP := os.Getenv("PUBLIC_IP")
	key, err := os.ReadFile(keypath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		fmt.Println("err in parsing")
		log.Fatal(err)
		return nil
	}
	auth := ssh.PublicKeys(signer)

	hostkey, err := knownhosts.New("/home/suresh/.ssh/known_hosts")
	if err != nil {
		fmt.Println("err in hostkey verify config")
		log.Fatal(err)
		return nil
	}

	config := &ssh.ClientConfig{
		User: "vpnadmin",
		Auth: []ssh.AuthMethod{
			auth,
		},
		HostKeyCallback: hostkey,
	}

	client, err := ssh.Dial("tcp", IP+":22", config)
	if err != nil {
		fmt.Println("err in dail")
		log.Fatal(err)
		return nil
	}
	fmt.Println("ssh done")
	return client
}