package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Barang struct {
	Nama   string
	Harga  int
	Jumlah int
}

func bacaTeks(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())

}

func bacaAngka(scanner *bufio.Scanner) int {
	teks := bacaTeks(scanner)
	angka, err := strconv.Atoi(teks)
	if err != nil || angka < 0 {
		fmt.Println("Input Angka nonnegatif")
		return 0
	}
	return angka
}

func hitungSubtotal(barang Barang) int {
	return barang.Harga * barang.Jumlah
}

func cetakStruk(daftarbarang []Barang) {
	total := 0

	fmt.Println("\n===STRUK ANDA===")
	for _, barang := range daftarbarang {
		subtotal := hitungSubtotal(barang)

		fmt.Printf("%s, %dx, %d = %d\n",
			barang.Nama,
			barang.Jumlah,
			barang.Harga,
			subtotal)

		total += subtotal

	}

	fmt.Println("==============")
	fmt.Printf("Total: %d\n", total)

}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukan Jumlah Jenis Barang :")
	n := bacaAngka(scanner)

	daftarBarang := make([]Barang, 0, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nBarang Ke-%d\n", i+1)

		fmt.Println("Masukan Nama Barang :")
		nama := bacaTeks(scanner)

		fmt.Println("Masukan Harga Satuan :")
		harga := bacaAngka(scanner)

		fmt.Println("Masukan Jumlah Barang : ")
		jumlah := bacaAngka(scanner)

		barang := Barang{
			Nama:   nama,
			Harga:  harga,
			Jumlah: jumlah,
		}
		daftarBarang = append(daftarBarang, barang)
	}
	cetakStruk(daftarBarang)

}
