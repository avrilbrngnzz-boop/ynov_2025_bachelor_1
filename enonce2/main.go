package main

func meteo(prenom string, temperature []int) {
	max := temperature[0]
	min := temperature[0]
	for _, temp := range temperature {
		if temp > max {
			max = temp
		}
		if temp < min {
			min = temp
		}
	}
	println(prenom, "a relevé une température maximale de", max, "et une minimale de", min)
}

func main() {
	meteo("Alice", []int{12, 15, 14, 11, 13})
}
