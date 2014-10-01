package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "genkeys":
			cmdGenkeys()
		case "help":
			cmdHelp()
		default:
			fmt.Println("Comando <" + os.Args[1] + "> incorrecto. Comprueba la ayuda con el comando \"help\".")
		}
	} else {
		cmdHelp()
	}
}

func cmdGenkeys() {
	if len(os.Args) >= 3 {
		path, err := os.Getwd()
		if err != nil {
			fmt.Println("<Error interno> No se ha podido determinar el directorio actual.")
			return
		}

		err = GenerateRsaKeys(path, os.Args[2])
		if err != nil {
			fmt.Println("<Error interno> Las llaves no han podido ser generadas.")
			return
		}

		fmt.Println("Las llaves han sido creadas con éxito.")
	} else {
		fmt.Println("Los argumentos no son suficientes:")
		fmt.Println("compostelas genkeys <usuario>")
	}
}

func cmdHelp() {
	fmt.Println("compostelas <comando> [Opciones]")
	fmt.Println()
	fmt.Println("Comandos:")
	fmt.Println("\tgenkeys <usuario>\tGenera un par de llaves publicas/privadas para el <usuario>.")
	fmt.Println("\thelp\t\tMuestra este dialogo de ayuda.")
}
