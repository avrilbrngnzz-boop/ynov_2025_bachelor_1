//probleme 2

package main

import (
	"fmt"
	"time"
)

func Depense(table []int) {
	start := time.Now()
	somme := 0
	for i := 0; i < len(table); i++ {
		somme += table[i]
	}
	fmt.Println(somme)
	duration := time.Since(start)
	fmt.Println("Temps d'exécution :", duration)
}

func main() {
	Depense([]int{12, 14, 12, 36, 95, 1, 2})
}
