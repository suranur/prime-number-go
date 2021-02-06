package main

import "fmt"

func main() {

	var sayi int
	fmt.Println("Asal kontrolu icin bir sayi girin: ")
	fmt.Scanln(&sayi)
	if asalMi(sayi) == 1 {
		fmt.Println("Sayi asaldir")
	} else if asalMi(sayi) == 0 {
		fmt.Println("Sayi asal degildir")
	}

}//67280421310721
func asalMi(sayi int) int {
	//var a int =sayi-1
	if sayi < 2 {

		return 0
	}
	if sayi == 2 {
		fmt.Println("Sayi asaldir")
		return 1
	} else {
		for i := 2; i < sayi-1; i++ {
			if sayi%i == 0 {
				return 0
			}
		}
	}
	return 1

}
