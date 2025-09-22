package main

import "fmt"

func evaluerNote(note int, salle string) {
	if note < 0 || note > 20 {
		fmt.Println("Erreur : note invalide")
		return
	}
	fmt.Println("Note :", note)
	if note < 10 {
		fmt.Println("Échec")
	} else if note < 13 {
		fmt.Println("Passable")
	} else if note < 16 {
		fmt.Println("Bien")
	} else {
		fmt.Println("Excellent")
	}
	fmt.Println("Salle :", salle)
}

func main() {
	evaluerNote(15, "B204")
	evaluerNote(8, "A101")
	evaluerNote(21, "C302")
	evaluerNote(17, "D105")
}
