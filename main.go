package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // inicializamos un lector asi podemos introducir la patente que queramos
	fmt.Print("Ingrese la patente: ")
	patenteRaw, _ := reader.ReadString('\n')
	patente := strings.TrimSpace(patenteRaw) //separamos cada simbolo del string ingresado en un slice

	fmt.Print("Ingrese el valor de k: ")             // ingresamos la K cantidad de veces
	kString, _ := reader.ReadString('\n')            // leemos el k ingresado
	k, _ := strconv.Atoi(strings.TrimSpace(kString)) //convertimos el k en un int

	p := siguiente(patente, k)
	fmt.Println(p)
}

func siguiente(patente string, k int) string {
	for ; k > 0; k-- {
		patente = rotarUltimaLetra(strings.ToUpper(patente))
	}

	return patente
}

func rotarUltimaLetra(patente string) string {
	ultimaLetra := patente[len(patente)-1]
	patenteSinUltimaLetra := patente[:len(patente)-1]
	if ultimaLetra == 'Z' {
		return rotarUltimaLetra(patenteSinUltimaLetra) + "A"
	}

	if ultimaLetra == '9' {
		return rotarUltimaLetra(patenteSinUltimaLetra) + "0"
	}

	return patenteSinUltimaLetra + string(ultimaLetra+1)
}
