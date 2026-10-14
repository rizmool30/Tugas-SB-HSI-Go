package main

import (
	"fmt"
)

func main() {
	var harian, pekanan, akhir, NA float64
	fmt.Print("masukan nilai harian: ")
	fmt.Scan(&harian)

	fmt.Print("masukan nilai pekanan: ")
	fmt.Scan(&pekanan)

	fmt.Print("masukan nilai akhir: ")
	fmt.Scan(&akhir)

	NA = (0.20 * harian) + (0.25 * pekanan) + (0.55 * akhir)

	if NA >= 91 && NA <= 100 {
		fmt.Println("Sangat Baik")
	} else if NA >= 80 && NA <= 90 {
		fmt.Println("Baik")
	} else if NA >= 66 && NA <= 79 {
		fmt.Println("Cukup")
	} else if NA >= 51 && NA <= 65 {
		fmt.Print("Kurang")
	} else if NA >= 0 && NA <= 50 {
		fmt.Print("Sangat Kurang")
	} else {
		fmt.Println("eror")
	}

}
