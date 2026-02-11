package main

import "fmt"

type Peserta struct {
	Nama     string
	Status   string
	Kategori string
	Nilai    int
}

func main() {
	var data = [5]Peserta{
		{"Alya", "active", "Web", 85},
		{"Bima", "active", "Data", 75},
		{"Citra", "inactive", "Web", 901},
		{"Doni", "active", "Web", 95},
		{"Eka", "active", "Data", 80},
	}

	fmt.Println("Peserta yang lolos kelas lanjutan:")

	for i := 0; i < len(data); i++ {
		if data[i].Status == "active" &&
			data[i].Kategori == "Web" &&
			data[i].Nilai >= 80 {

			fmt.Println("Nama:", data[i].Nama)
			fmt.Println("Nilai:", data[i].Nilai)
			fmt.Println()
		}
	}
}
