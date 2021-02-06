package main

import "fmt"

func main() {
	fmt.Println("Asallik kontrolu icin bir sayi girin: ")
	var sayi int
	fmt.Scanln(&sayi)
	if asalMi(sayi) {
		fmt.Println("Sayi asaldir")
	} else {
		fmt.Println("Sayi asal degildir")
	}

}
func asalMi(sayi int) bool {
	var a int = sayi / 2
	if sayi%2 == 0 {
		return false
	}
	for i := 3; i < a; i++ {
		if sayi%i == 0 {
			return false
		}

	}
	return true
}
