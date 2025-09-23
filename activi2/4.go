package main

import "fmt"

func verifierDoublon(liste []string) {
	for i := 0; i < len(liste); i++ {
		for j := i + 1; j < len(liste); j++ {
			if liste[i] == liste[j] {
				fmt.Println("Doublon trouvé :", liste[i])
				return
			}
		}
	}
	fmt.Println("Aucun doublon trouvé")
}

func main() {
	objets := []string{"clé", "lampe", "stylo", "lampe", "livre"}
	verifierDoublon(objets)
}
