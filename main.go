package main

import (
	"fmt"

	"github.com/SureshS03/pnet/cmd"
)

func main() {
	fmt.Println(cmd.CommandMaps[cmd.StartVPN])
	cmd.Exe()
}