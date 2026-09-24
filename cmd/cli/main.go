package main

import (
	"fmt"
	"os"
	"slices"
)

var (
	Version = "dev"
)

func main() {

	args := os.Args[1:]
	cmd := args[0]

	help := `Usage: wallet <command> [<args>]
The commands are:
	help        Show this help message
	version     Show the version of the wallet
	withdraw    Withdraw funds from the wallet
	deposit     Deposit funds into the wallet
`

	if len(args) == 0 {
		fmt.Println(help)
		os.Exit(2)
	}

	if cmd == "help" {
		fmt.Println(help)
		os.Exit(0)
	}

	if cmd == "version" {
		fmt.Println("wallet " + Version)
		os.Exit(0)
	}

	if slices.Contains([]string{"withdraw", "deposit"}, cmd) {
		fmt.Printf("<%s>: not implemented\n", cmd)
		os.Exit(0)
	}

	fmt.Printf("unknown command: %s\n", cmd)
	os.Exit(2)

}
