package main

import "fmt"

func caisse(montant int, somme int) {
	if somme < montant {
		fmt.Println("Il manque", montant-somme, "euros")
	}
	if somme == montant {
		fmt.Println("Merci, au revoir")
	} else {
		fmt.Println("Voici votre monnaie:", somme-montant, "euros")
	}
}

func main() {
	caisse(50, 20)
	caisse(50, 50)
	caisse(50, 70)
}
