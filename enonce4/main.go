package main

import "fmt"

func autoroute(vitesse int, vitesseMAx int) {
	if vitesse > vitesseMAx {
		fmt.Println("vous êtes en excès de vitesse")
	} else {
		fmt.Println("vous êtes dans la bonne vitesse")
	}
}

func main() {
	autoroute(130, 110)
	autoroute(90, 110)
}
