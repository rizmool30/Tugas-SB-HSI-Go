package main

import "fmt"

func intersect(a, b []string) []string {
	barang1 := make(map[string]bool)
	barang2 := make(map[string]bool)
	hasil := []string{}

	for _, item := range b {
		barang1[item] = true
	}

	for _, item := range a {
		if barang1[item] && !barang2[item] {
			hasil = append(hasil, item)
			barang2[item] = true

		}
	}

	return hasil
}

func main() {
	snack1 := []string{"apple", "mango", "banana", "guava"}
	snack2 := []string{"apple", "biscuit", "cereal", "guava"}

	fmt.Print(intersect(snack1, snack2))

}
