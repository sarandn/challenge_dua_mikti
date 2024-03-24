package main

import (
	"fmt"
)

func main() {
	// Inisialisasi data untuk Mark dan John pada dua kondisi berbeda
	// Kondisi 1
	beratMark1 := 78.0  // Berat Mark dalam kg
	tinggiMark1 := 1.69 // Tinggi Mark dalam m
	beratJohn1 := 92.0  // Berat John dalam kg
	tinggiJohn1 := 1.95 // Tinggi John dalam m

	// Kondisi 2
	beratMark2 := 95.0  // Berat Mark dalam kg
	tinggiMark2 := 1.88 // Tinggi Mark dalam m
	beratJohn2 := 85.0  // Berat John dalam kg
	tinggiJohn2 := 1.76 // Tinggi John dalam m

	// Fungsi untuk menghitung BMI
	hitungBMI := func(berat, tinggi float64) float64 {
		return berat / (tinggi * tinggi)
	}

	// Menghitung BMI untuk Mark dan John pada kedua kondisi
	bmiMark1 := hitungBMI(beratMark1, tinggiMark1)
	bmiJohn1 := hitungBMI(beratJohn1, tinggiJohn1)

	bmiMark2 := hitungBMI(beratMark2, tinggiMark2)
	bmiJohn2 := hitungBMI(beratJohn2, tinggiJohn2)

	// Membandingkan BMI Mark dan John
	markHigherBMI1 := bmiMark1 > bmiJohn1
	markHigherBMI2 := bmiMark2 > bmiJohn2

	// Menampilkan hasil perbandingan BMI
	fmt.Println("Data 1:")
	fmt.Printf("BMI Mark: %.2f\n", bmiMark1)
	fmt.Printf("BMI John: %.2f\n", bmiJohn1)
	fmt.Printf("Apakah BMI Mark lebih tinggi dari John? %t\n\n", markHigherBMI1)

	fmt.Println("Data 2:")
	fmt.Printf("BMI Mark: %.2f\n", bmiMark2)
	fmt.Printf("BMI John: %.2f\n", bmiJohn2)
	fmt.Printf("Apakah BMI Mark lebih tinggi dari John? %t\n", markHigherBMI2)
}
