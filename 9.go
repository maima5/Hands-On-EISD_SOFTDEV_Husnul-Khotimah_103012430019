package main

import "fmt"

func main() {
	var kuota int = 3

	var pendaftar [5]string
	pendaftar[0] = "Alya"
	pendaftar[1] = "Bima"
	pendaftar[2] = "Citra"
	pendaftar[3] = "Doni"
	pendaftar[4] = "Eka"


	var diterima [5]string
	var ditolak [5]string

	var jumlahDiterima int = 0
	var jumlahDitolak int = 0
	var stop bool = false

	for i := 0; i < 5 && !stop; i++ {

		if jumlahDiterima < kuota {
			// Kl masih ada kuota maka diterima
			diterima[jumlahDiterima] = pendaftar[i]
			jumlahDiterima++
		} else {
			// Kl kuota penuh maka ditolak
			ditolak[jumlahDitolak] = pendaftar[i]
			jumlahDitolak++
			stop = true //flag boolean pengganti break
		}
	}

	fmt.Println("Daftar Peserta Diterima:")
	for i := 0; i < jumlahDiterima; i++ {
		fmt.Println("-", diterima[i])
	}

	fmt.Println("\nDaftar Peserta Ditolak:")
	for i := 0; i < jumlahDitolak; i++ {
		fmt.Println("-", ditolak[i])
	}
}
