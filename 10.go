package main

import "fmt"

func main() {
	var stok int = 4
	var habis bool = false
	var peminjam [6]string
	peminjam[0] = "Alya"
	peminjam[1] = "Bima"
	peminjam[2] = "Citra"
	peminjam[3] = "Doni"
	peminjam[4] = "Eka"
	peminjam[5] = "Farah"

	var berhasil [6]string
	var gagal [6]string

	var jumlahBerhasil int = 0
	var jumlahGagal int = 0

	for i := 0; i < 6 && !habis; i++ {

		if stok > 0 {
			berhasil[jumlahBerhasil] = peminjam[i]
			jumlahBerhasil++
			stok--
		} else {
			gagal[jumlahGagal] = peminjam[i]
			jumlahGagal++
			habis = true // set flag habis jika stok habis
		}
	}

	fmt.Println("Peminjaman Berhasil:")
	for i := 0; i < jumlahBerhasil; i++ {
		fmt.Println("-", berhasil[i])
	}

	fmt.Println("\nPeminjaman Gagal:")
	for i := 0; i < jumlahGagal; i++ {
		fmt.Println("-", gagal[i])
	}
}
