package main

import "fmt"

func main() {
	emails := [6]string{
		"Andi@gmail.com",
		"Budi@gmail.com",
		"Sari@gmail.com",
		"Andi@gmail.com",
		"Rina@gmail.com",
		"Budi@gmail.com",
	}

	var adaDuplikat, stop bool
	adaDuplikat = false
	stop = false
	fmt.Println("Email duplikat:")

	for i := 0; i < 6  ; i++ {
		stop = false
		for j := i + 1; j < 6 && stop == false; j++ {
			if emails[i] == emails[j] {
				fmt.Println("-", emails[i])
				adaDuplikat = true
				stop = true
			}
		}
	}

	if !adaDuplikat {
		fmt.Println("Tidak ada data duplikat")
	}
}
