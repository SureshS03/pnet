package cmd

type Actions string

const (
	StartVPN Actions = "start_vpn"
	StopVPN Actions = "stop_vpn"
)

var CommandMaps = map[Actions]string {
	StartVPN: "sudo systemctl start wg-quick@wg0",
	StopVPN: "sudo systemclt stop wg-quick@wg0",
}