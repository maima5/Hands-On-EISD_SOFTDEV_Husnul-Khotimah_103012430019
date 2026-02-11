package main

import "fmt"

const max int = 100


func main() {

	var divisi [max]string
	var kuota [max]int
	var terdaftar [max]int

	divisi[0] = "Software Development"
	divisi[1] = "UI/UX Design"
	divisi[2] = "Technopreneur"
	divisi[3] = "Intelligence System"

	kuota[0] = 20
	kuota[1] = 15
	kuota[2] = 25
	kuota[3] = 10

	terdaftar[0] = 18
	terdaftar[1] = 15
	terdaftar[2] = 10
	terdaftar[3] = 10

	nData := 4

	fmt.Println("Study Group yang Masih Tersedia:")
	for i := 0; i < nData; i++ {
		if terdaftar[i] < kuota[i] {
			fmt.Println("-", divisi[i])
		}
	}

	fmt.Println("\nStudy Group yang Sudah Penuh:")
	for i := 0; i < nData; i++ {
		if terdaftar[i] >= kuota[i] {
			fmt.Println("-", divisi[i])
		}
	}
}

