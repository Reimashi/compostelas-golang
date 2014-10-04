package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Pilgrim struct {
	name          string
	dni           string
	address       string
	creationTime  time.Time
	creationPlace string
	motivations   string
}

func (t *Pilgrim) GetName() string {
	return t.name
}

func (t *Pilgrim) GetDni() string {
	return t.dni
}

func (t *Pilgrim) GetAddress() string {
	return t.address
}

func (t *Pilgrim) GetCreationTime() time.Time {
	return t.creationTime
}

func (t *Pilgrim) GetCreationPlace() string {
	return t.creationPlace
}

func (t *Pilgrim) GetMotivations() string {
	return t.motivations
}

func GetPilgrimFromJson(data []byte) (Pilgrim, error) {
	var pil Pilgrim

	err := json.Unmarshal(data, &pil)
	if err != nil {
		return pil, err
	}

	return pil, nil
}

// TODO Acabar el método
func GetPilgrimFromTerminal(hostelid string) (Pilgrim, error) {
	var pil Pilgrim

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Formulario: Alta de Peregrinos ===")

	fmt.Print("Nombre: ")
	namer, _ := reader.ReadString('\n')

	fmt.Print("DNI: ")
	dnir, _ := reader.ReadString('\n')

	fmt.Print("Dirección: ")
	addressr, _ := reader.ReadString('\n')

	fmt.Print("Motivaciones: ")
	motivationsr, _ := reader.ReadString('\n')

	pil = Pilgrim{namer, dnir, addressr, time.Now(), hostelid, motivationsr}

	return pil, nil
}

func (t *Pilgrim) GetJson() ([]byte, error) {
	return json.Marshal(t)
}
