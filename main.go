package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func getTemanByAbsen(absen int) Teman {
	var teman Teman
	switch absen {
	case 1:
		teman = Teman{
			Nama:      "Andi",
			Alamat:    "Jakarta",
			Pekerjaan: "Mahasiswa",
			Alasan:    "ingin mengembangkan ilmu",
		}
	case 2:
		teman = Teman{
			Nama:      "Budi",
			Alamat:    "Bekasi",
			Pekerjaan: "Programmer",
			Alasan:    "belajar dunia baru",
		}
	case 3:
		teman = Teman{
			Nama:      "Cindy",
			Alamat:    "Depok",
			Pekerjaan: "Marketing",
			Alasan:    "tertarik di dunia koding",
		}
	default:
		teman = Teman{}
	}
	return teman
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Masukkan nomor absen teman yang ingin ditampilkan")
		return
	}

	absen, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	teman := getTemanByAbsen(absen)
	if teman.Nama == "" {
		fmt.Println("Tidak ditemukan teman dengan nomor absen tersebut")
		return
	}

	fmt.Printf("Nama: %s\n", teman.Nama)
	fmt.Printf("Alamat: %s\n", teman.Alamat)
	fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
	fmt.Printf("Alasan: %s\n", teman.Alasan)
}
