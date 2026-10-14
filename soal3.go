package main

import "fmt"

func main() {
	var n int
	var jumlah int

	fmt.Print("masukan nilai : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			jumlah += i
			fmt.Print(i, "")

		}
	}
	fmt.Println("\njumlah faktor: ", jumlah)

}
