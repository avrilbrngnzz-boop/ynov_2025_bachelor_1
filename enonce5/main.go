package main

import "fmt"

func verifierMachine(compteur int, marche bool, temperature float64, pression float64) {
	if !marche {
		fmt.Println("Machine déjà arrêtée")
	} else if compteur > 10000 || temperature > 80 || pression > 120 {
		fmt.Println("Maintenance requise")
	} else {
		fmt.Println("Machine OK")
	}
}

func main() {
	verifierMachine(9500, true, 75.0, 110.0)
	verifierMachine(12000, true, 85.0, 130.0)
	verifierMachine(8000, false, 60.0, 100.0)
}
