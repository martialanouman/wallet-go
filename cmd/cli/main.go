package main

import (
	"fmt"
	"os"
	"slices"
)

var (
	version = "dev"
)

func main() {
	args := os.Args
	cmd := ""

	if len(args) > 1 {
		cmd = args[1]
	}

	help := `Usage: wallet <command> [<args>]
The commands are:
	help        Show this help message
	version     Show the version of the wallet
	withdraw    Withdraw funds from the wallet
	deposit     Deposit funds into the wallet
	balance     Show the balance of the wallet
`

	if cmd == "" {
		fmt.Fprintln(os.Stderr, help)
		os.Exit(2)
	}

	if cmd == "help" {
		fmt.Println(help)
		os.Exit(0)
	}

	if cmd == "version" {
		fmt.Println("wallet " + version)
		os.Exit(0)
	}

	if slices.Contains([]string{"withdraw", "deposit", "balance"}, cmd) {
		fmt.Fprintf(os.Stderr, "%s: not implemented\n", cmd)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd)
	fmt.Fprint(os.Stderr, help)
	os.Exit(2)
}
