package main

import (
	"fmt"
	"strings"
)

// FUNGSI NOMOR 1 - PERSEGI PANJANG
func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

// FUNGSI NOMOR 2 - INTRODUCE
func introduce(nama, gender, pekerjaan, umur string) (kalimat string) {
	var title string
	if gender == "perempuan" {
		title = "Bu "
	} else {
		title = "Pak "
	}
	kalimat = title + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	return
}

// FUNGSI NOMOR 3 - VARIADIC FUNCTION
func buahFavorit(nama string, buah ...string) (kalimat string) {
	kalimat = "halo nama saya " + nama + " dan buah favorit saya adalah "
	// for _, buah := range buah {
	// 	kalimat += buah + ", "
	// }
	if len(buah) > 0 {
		kalimat += strings.Join(buah, ", ")
	}
	return
}

func main() {
	// SOAL 1
	fmt.Println("Soal 1")
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("luas :", luas)
	fmt.Println("keliling :", keliling)
	fmt.Println("volume :", volume)
	fmt.Println()

	// SOAL 2
	fmt.Println("Soal 2")
	john := introduce("John", "laki-laki", "penulis", "30")
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(john)
	fmt.Println(sarah)
	fmt.Println()

	// SOAL 3
	fmt.Println("Soal 3")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("john", buah...)
	fmt.Println(buahFavoritJohn)
	fmt.Println()

	// SOAL 4
	fmt.Println("Soal 4")
	var dataFilm = []map[string]string{}
	// --- CLOSURE ---
	tambahDataFilm := func(judul, durasi, genre, tahun string) {
		dataFilm = append(dataFilm, map[string]string{
			"title": judul,
			"jam":   durasi,
			"genre": genre,
			"tahun": tahun,
		})
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
	fmt.Println()

}
