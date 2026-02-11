package main

import "fmt"

func main() {

	peserta := [4]int{18, 15, 10, 10}

	total := 0
	for i := 0; i < 4; i++ {
		total = total + peserta[i]
	}

	rataRata := float64(total) / float64(len(peserta))

	fmt.Println("Total seluruh peserta:", total)
	fmt.Println("Rata-rata peserta per Study Group:", rataRata)
}
