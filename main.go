package main

import (
	"crypto/aes"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "genkeys":
			cmdGenkeys()
		case "new":
			cmdNew()
		case "aes":
			cmdAes()
		case "help":
			cmdHelp()
		default:
			fmt.Println("Comando <" + os.Args[1] + "> incorrecto. Comprueba la ayuda con el comando \"help\".")
		}
	} else {
		cmdHelp()
	}
}

func cmdNew() {
	if len(os.Args) >= 3 {
		pil, err := GetPilgrimFromTerminal(os.Args[2])
		if err != nil {
			fmt.Println("<Error interno> No se han podido generar los datos del peregrino.")
			pil.GetName() //esto es porque sino no me deja hacer build
			return
		}

	} else {
		fmt.Println("Los argumentos no son suficientes:")
		fmt.Println("compostelas new <nombre_oficina>")
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
	fmt.Println("\tnew <nombre_oficina>\tGenera una compostela.")
	fmt.Println("\taes <encriptar | desencriptar | desencriptarmal>\tMuestra una demostración de encriptado AES de tipo CFB")
	fmt.Println("\thelp\t\tMuestra este dialogo de ayuda.")
}

func cmdAes() {

	if len(os.Args) >= 3 {
		data := []byte{0, 1}
		encript := []byte{117, 125}
		var key = "1234567890123456"
		var keyerr = "6543210987654321"
		var iv = []byte(key)[:aes.BlockSize] // No se como calcular el iv

		switch os.Args[2] {
		case "encriptar":
			resultado := encriptar_aes(data, []byte(key), iv)
			fmt.Printf("Encrypting %v -> %v\n", data, []byte(resultado))

		case "desencriptar":
			resultado := desencriptar_aes(encript, []byte(key), iv)
			fmt.Printf("Decrypting %v -> %v\n", encript, []byte(resultado))

		case "desencriptarmal":
			resultado := desencriptar_aes(encript, []byte(keyerr), iv)
			fmt.Printf("Decrypting %v -> %v\n", encript, []byte(resultado))

		default:
			fmt.Println("Opción <" + os.Args[2] + "> incorrecta. Use <encriptar | desencriptar | desencriptarmal>")
		}

	} else {
		fmt.Println("Los argumentos no son suficientes:")
		fmt.Println("compostelas aes <encriptar | desencriptar | desencriptarmal>")
	}
}
