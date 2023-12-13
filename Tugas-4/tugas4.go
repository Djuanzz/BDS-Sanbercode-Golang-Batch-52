package main

import "fmt"

func main() {
	// SOAL 1
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 != 0 {
			fmt.Println(i, "- I Love Coding")
		} else if i%2 == 0 {
			fmt.Println(i, "- Santai")
		} else {
			fmt.Println(i, "- Berkualitas")
		}
	}
	fmt.Println()

	// SOAL 2
	for i := 0; i <= 7; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
	fmt.Println()

	// SOAL 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	kalimatNew := kalimat[2:]
	fmt.Println(kalimatNew)
	fmt.Println()

	// SOAL 4
	var sayuran = []string{}
	sayuranNeW := append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
	fmt.Println(sayuranNeW)
	fmt.Println()

	// SOAL 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
	}
	fmt.Println("Volume Balok =", volume)
}
