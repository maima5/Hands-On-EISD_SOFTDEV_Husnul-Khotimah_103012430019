package main

import "fmt"

type Pendaftar struct {
	Nama   string
	Status string
	Nilai  int
}

type Hasil struct {
	Nama  string
	Nilai int
}

func main() {
	var data [5]Pendaftar

	data[0] = Pendaftar{"Alya", "active", 85}
	data[1] = Pendaftar{"Bima", "inactive", 90}
	data[2] = Pendaftar{"Citra", "active", 70}
	data[3] = Pendaftar{"Doni", "active", 95}
	data[4] = Pendaftar{"Eka", "inactive", 60}

	var mapped [5]Hasil
	var count int = 0

	for i := 0; i < 5; i++ {
		if data[i].Status == "active" && data[i].Nilai >= 80 {

			mapped[count].Nama = data[i].Nama
			mapped[count].Nilai = data[i].Nilai

			count++
		}
	}

	fmt.Println("Hasil Filter:")
	for i := 0; i < count; i++ {
		fmt.Println("nama:", mapped[i].Nama)
		fmt.Println("nilai:", mapped[i].Nilai)
		fmt.Println()
	}
}
