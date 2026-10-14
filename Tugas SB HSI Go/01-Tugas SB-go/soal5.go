package main

import (
	"fmt"
	"strings"
)

func main() {
	var kalimat string
	var kunci string

	fmt.Print("Masukan Kalimat: ")
	fmt.Scan(&kalimat)

	fmt.Print("Masukan Kata Kunci: ")
	fmt.Scan(&kunci)

	kalimatKecil := strings.ToLower(kalimat)
	kunciKecil := strings.ToLower(kunci)

	if strings.Contains(kalimatKecil, kunciKecil) && kunci != "" {
		fmt.Println("Kata Ditemukan")
	} else {
		fmt.Println("Kata tidak ditemukan")
	}

	//	masih kurang mas, bingung kenapa saat nulis satu kata bisa ke detect, tapi sebaliknya pas nulis kalimat
	// wallahu a'lam

}
