// compostelas project main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "genkeys":
			calcularClaves()
		case "help":
			cmdHelp()
		default:
			fmt.Println("Command " + os.Args[1] + "incorrect. Check help with \"help\" command.")
		}
	} else {
		cmdHelp()
	}
}

func cmdHelp() {
	fmt.Println("compostelas <command> [Options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("\tgenkeys <user>\tGenerate public/private keys for <user>.")
	fmt.Println("\thelp\t\tShow this help.")
}
