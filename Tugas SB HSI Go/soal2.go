package main

import "fmt"

// fmt.Println(checkDuplicates([]string{"apple", "mango", "orange", "orange", "banana", "apple", "guava"})) // [apple orange]

// func checkDuplikasi(items []string) []string {
// 	ditemukan := make(map[string]bool)
// 	sudahDuplikasi := make(map[string]bool)
// 	hasil := []string{}

// 	for _, item := range items {
// 		if ditemukan[item] {
// 			if !sudahDuplikasi[item] {
// 				hasil = append(hasil, item)
// 				sudahDuplikasi[item] = true
// 			}
// 		} else {
// 			ditemukan[item] = true
// 		}
// 	}
// 	return hasil

// }

// func main() {
// 	buah := []string{
// 		"apple", "mango", "orange", "orange", "banana", "apple", "guava",
// 	}
// 	fmt.Println("Data Buah :", buah)
// 	fmt.Println("Buah Duplikasi :", checkDuplikasi(buah))

// }

// fmt.Println(checkDuplicates([]string{}))                                                                 // []

// func checkDuplikasi(items []string) []string {
// 	ditemukan := make(map[string]bool)
// 	sudahDuplikasi := make(map[string]bool)
// 	hasil := []string{}

// 	for _, item := range items {
// 		if ditemukan[item] {
// 			if !sudahDuplikasi[item] {
// 				hasil = append(hasil, item)
// 				sudahDuplikasi[item] = true
// 			}
// 		} else {
// 			ditemukan[item] = true
// 		}
// 	}
// 	return hasil

// }

// func main() {
// 	none := []string{}
// 	fmt.Println("Data Buah :", none)
// 	fmt.Println("Buah Duplikasi :", checkDuplikasi(none))

// }

//fmt.Println(checkDuplicates([]string{"kiwi", "mango", "banana"}))                                        // []

// func checkDuplikasi(items []string) []string {
// 	ditemukan := make(map[string]bool)
// 	sudahDuplikasi := make(map[string]bool)
// 	hasil := []string{}

// 	for _, item := range items {
// 		if ditemukan[item] {
// 			if !sudahDuplikasi[item] {
// 				hasil = append(hasil, item)
// 				sudahDuplikasi[item] = true
// 			}
// 		} else {
// 			ditemukan[item] = true
// 		}
// 	}
// 	return hasil

// }

// func main() {
// 	buah := []string{"kiwi", "mango", "banana"}
// 	fmt.Println("Data Buah :", buah)
// 	fmt.Println("Buah Duplikasi :", checkDuplikasi(buah))

// }

// fmt.Println(checkDuplicates([]string{"kiwi", "kiwi"}))                                                   // [kiwi]
// func checkDuplikasi(items []string) []string {
// 	ditemukan := make(map[string]bool)
// 	sudahDuplikasi := make(map[string]bool)
// 	hasil := []string{}

// 	for _, item := range items {
// 		if ditemukan[item] {
// 			if !sudahDuplikasi[item] {
// 				hasil = append(hasil, item)
// 				sudahDuplikasi[item] = true
// 			}
// 		} else {
// 			ditemukan[item] = true
// 		}
// 	}
// 	return hasil

// }

// func main() {
// 	buah := []string{"kiwi", "kiwi"}
// 	fmt.Println("Data Buah :", buah)
// 	fmt.Println("Buah Duplikasi :", checkDuplikasi(buah))

// }

// fmt.Println(checkDuplicates([]string{"a", "b", "a", "b", "c"}))                                          // [a b]

// func checkDuplikasi(items []string) []string {
// 	ditemukan := make(map[string]bool)
// 	sudahDuplikasi := make(map[string]bool)
// 	hasil := []string{}

// 	for _, item := range items {
// 		if ditemukan[item] {
// 			if !sudahDuplikasi[item] {
// 				hasil = append(hasil, item)
// 				sudahDuplikasi[item] = true
// 			}
// 		} else {
// 			ditemukan[item] = true
// 		}
// 	}
// 	return hasil

// }

// func main() {
// 	alphabet := []string{"a", "b", "a", "b", "c"}
// 	fmt.Println("Data Buah :", alphabet)
// 	fmt.Println("Buah Duplikasi :", checkDuplikasi(alphabet))

// }

// fmt.Println(checkDuplicates([]string{"x", "y", "z", "x", "z", "y", "x"}))                                // [x y z]

func checkDuplikasi(items []string) []string {
	ditemukan := make(map[string]bool)
	sudahDuplikasi := make(map[string]bool)
	hasil := []string{}

	for _, item := range items {
		if ditemukan[item] {
			if !sudahDuplikasi[item] {
				hasil = append(hasil, item)
				sudahDuplikasi[item] = true
			}
		} else {
			ditemukan[item] = true
		}
	}
	return hasil

}

func main() {
	alphabet := []string{"x", "y", "z", "x", "z", "y", "x"}
	fmt.Println("Data Buah :", alphabet)
	fmt.Println("Buah Duplikasi :", checkDuplikasi(alphabet))

}
