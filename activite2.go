//probleme 1

package main

import (
	"fmt"
	"time"
)

func Prenom(table []string) {
	start := time.Now()
	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
	duration := time.Since(start)
	fmt.Println("Temps d'exécution :", duration)
}

func main() {
	Prenom([]string{"Valéri", "Sandrine", "Arnaud", "Enzo"})
}
