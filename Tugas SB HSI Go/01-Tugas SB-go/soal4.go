package main

import "fmt"

func main() {
	// soal 2 x 2 = 2, 3 x 3 =5, 4 x 4 = 8, 10 x 10 =50
	// bilangan genap contoh 2 x 2 = 2 / n x n / 2 = 4
	// rumus genap 4 x 4 = 8 : n x n /2 = 8

	// bilangan ganjil  3 x 3 = 5 = ((n x n) + 1) (/2) = 5
	// 5 x 5 = ((5 x 5) + 1) /2 = 13
	// so rumus buat ganjil (n*n+)/2
	var n int

	fmt.Print("Masukan Nilai N: ")
	fmt.Scan(&n)

	totalpetak := n * n
	Tanaman := totalpetak / 2

	fmt.Println("hasil:", Tanaman)

}
