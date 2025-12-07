package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/SureshS03/pnet/internal"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use: "pnet",
}

var start_vpn_cmd = &cobra.Command{
	Use: "start_vpn",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("im at running state")
		client := internal.ConnectToServer()
		defer client.Close()
		output, err := Start_vpn(client)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(output)

	},
}

func Exe() {
	root.AddCommand(start_vpn_cmd)
	err := root.Execute()
	if err != nil {
		log.Fatal(err)
	}
}