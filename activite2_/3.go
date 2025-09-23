package main

import "fmt"

func sendMessage(destinataire string, expediteur string) {
	fmt.Printf("Message envoyé de %s à %s\n", expediteur, destinataire)
}

func main() {
	groupe := []string{"Valéri", "Sandrine", "Arnaud", "Enzo"}

	for _, expediteur := range groupe {
		for _, destinataire := range groupe {
			if destinataire != expediteur {
				sendMessage(destinataire, expediteur)
			}
		}
	}
}
