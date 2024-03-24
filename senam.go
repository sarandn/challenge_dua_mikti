package main

import "fmt"

// Menghitung rata-rata skor dari sebuah tim
func hitungRataRata(skors []int) float64 {
	var total int
	for _, skor := range skors {
		total += skor
	}
	return float64(total) / float64(len(skors))
}

// Fungsi untuk mencetak skor rata-rata dan menentukan pemenang
func cetakHasilDanPemenang(skorsLumbalumba, skorsKoala [][]int) {
	for i, skors := range skorsLumbalumba {
		rataRataLumbaLumba := hitungRataRata(skors)
		rataRataKoala := hitungRataRata(skorsKoala[i])

		fmt.Printf("Data %d:\n", i+1)
		fmt.Printf("Skor rata-rata Lumba-lumba: %.2f\n", rataRataLumbaLumba)
		fmt.Printf("Skor rata-rata Koala: %.2f\n", rataRataKoala)

		switch {
		case rataRataLumbaLumba > rataRataKoala:
			fmt.Println("Lumba-lumba menang.")
		case rataRataKoala > rataRataLumbaLumba:
			fmt.Println("Koala menang.")
		case rataRataLumbaLumba >= 100 && rataRataKoala >= 100:
			fmt.Println("Hasil seri.")
		default:
			fmt.Println("Tidak ada pemenang.")
		}

		fmt.Println()
	}
}

func main() {
	skorsLumbalumba := [][]int{{96, 108, 89}, {97, 112, 101}, {97, 112, 101}}
	skorsKoala := [][]int{{88, 91, 110}, {109, 95, 123}, {109, 95, 106}}

	cetakHasilDanPemenang(skorsLumbalumba, skorsKoala)
}
