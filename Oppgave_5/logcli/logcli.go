package main

import (
	"os"
	"strconv"
	"fmt"
	"math"
)

/**
Vi importerer os.Args pakken for å hente input fra brukeren.
Så printer vi ut svaret under ved hjelp av sammen funksjon som i oppgave 4, "math.Log2".
Bare at denne gangen bruker vi "argInput" i steden for et predefinert tall.
 */
func Log2(){
	fmt.Println("Skriv inn et tall som du vil ta Logaritmen til 2 av:")

	arg1 := os.Args[1]
	argInput, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Svaret ditt er:  ")
	fmt.Println(math.Log2(argInput))
}
