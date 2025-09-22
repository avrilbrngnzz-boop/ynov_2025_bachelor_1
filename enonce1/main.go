package main

import "fmt"

func mairi(age int, prenom string) {
	if age >= 18 {
		fmt.Println(prenom, "est majeur")
	} else {
		fmt.Println(prenom, "est mineur")
	}
}

func main() {
	mairi(20, "Alice")
	mairi(17, "Arnaud")
}
