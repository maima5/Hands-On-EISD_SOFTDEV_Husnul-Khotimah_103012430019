package main

import "fmt"

func main() {
	var kuota, terdaftar int

	fmt.Print("Masukkan kuota maksimal: ")
	fmt.Scan(&kuota)

	fmt.Print("Masukkan jumlah peserta terdaftar: ")
	fmt.Scan(&terdaftar)

	if terdaftar > kuota {
		fmt.Println("Status: Pendaftaran ditolak")
	} else if terdaftar == kuota {
		fmt.Println("Status: Pendaftaran ditutup")
	} else if terdaftar >= kuota-3 {
		fmt.Println("Status: Hampir penuh")
	} else {
		fmt.Println("Status: Pendaftaran diterima")
	}
}
