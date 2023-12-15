package main

import (
	"fmt"
	"math"
)

func main() {
	// SOAL 1
	var luasLigkaran float64
	var kelilingLingkaran float64
	calcLuasLingkaran(7, &luasLigkaran)
	calcKelilingLingkaran(7, &kelilingLingkaran)
	fmt.Println("Luas Lingkaran: ", luasLigkaran)
	fmt.Println("Keliling Lingkaran: ", kelilingLingkaran)
	fmt.Println()

	// SOAL 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
	fmt.Println()

	// SOAL 3
	// INISIASI ARRAY
	var buah = []string{}

	addFruit(&buah, "Jeruk")
	addFruit(&buah, "Semangka")
	addFruit(&buah, "Mangga")
	addFruit(&buah, "Strawberry")
	addFruit(&buah, "Durian")
	addFruit(&buah, "Manggis")
	addFruit(&buah, "Alpukat")

	for i, fruit := range buah {
		fmt.Printf("%d. %s\n", i+1, fruit)
	}
	fmt.Println()

	// SOAL 4
	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. title    :  %s\n", i+1, film["title"])
		fmt.Println("   duration : ", film["duration"])
		fmt.Println("   genre    : ", film["genre"])
		fmt.Println("   year     : ", film["year"])
		fmt.Println()
	}

}

// FUNGSI SOAL 1
func calcLuasLingkaran(r float64, luas *float64) {
	*luas = r * r * math.Pi
}

func calcKelilingLingkaran(r float64, keliling *float64) {
	*keliling = 2 * r * math.Pi
}

// FUNGSI SOAL 2
func introduce(sentence *string, nama, gender, job, age string) {
	title := "Pak"
	if gender == "perempuan" {
		title = "Bu"
	}

	*sentence = title + " " + nama + " adalah seorang " + job + " yang berusia " + age + " tahun"
}

// FUNGSI SOAL 3
func addFruit(fruit *[]string, newFruit string) {
	*fruit = append(*fruit, newFruit)
}

// FUNGSI SOAL 4
func tambahDataFilm(judul, durasi, genre, tahun string, film *[]map[string]string) {
	*film = append(*film, map[string]string{
		"title":    judul,
		"year":     tahun,
		"genre":    genre,
		"duration": durasi,
	})
}
