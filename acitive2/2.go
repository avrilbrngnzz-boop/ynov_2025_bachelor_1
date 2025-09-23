// problème3
package main

import (
	"fmt"
	"time"
)

func SendMessage(personne int) {
	start := time.Now()
	for i := 0; i < personne; i++ {
		fmt.Println("message envoyé à XXX")
	}
	duration := time.Since(start)
	fmt.Println("Temps d'exécution :", duration)
}

func main() {
	SendMessage(12)
}
