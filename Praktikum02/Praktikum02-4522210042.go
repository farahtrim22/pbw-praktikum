package main

import "fmt"

type mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	dataMahasiswa := make(map[string]mahasiswa)

	dataMahasiswa["10000321"] = mahasiswa{"Malika", "3521240006", "Akuntansi"}
	dataMahasiswa["10000322"] = mahasiswa{"Tegar", "4402300010", "Teknik Informatika"}
	dataMahasiswa["10000323"] = mahasiswa{"Farel", "1324354657", "Pariwisata"}
	dataMahasiswa["10000324"] = mahasiswa{"Annisa", "9786756453", "Farmasi"}
	dataMahasiswa["10000325"] = mahasiswa{"zaky", "4402300021", "Teknik Informatika"}

	fmt.Println("DAFTAR MAHASISWA")
	for _, mhs := range dataMahasiswa {
		fmt.Println("Nama: ", mhs.Nama)
		fmt.Printf("NPM: %s\n", mhs.NPM)
		fmt.Printf("Jurusan: %s\n", mhs.Jurusan)
		fmt.Println()
	}

	CariNPM := "10000324"
	mahasiswa, found := dataMahasiswa[CariNPM]
	if found {
		fmt.Printf("\nData Mahasiswa dengan NPM %s ditemukan: \n", CariNPM)
		fmt.Println("Nama:", mahasiswa.Nama)
		fmt.Println("NPM:", mahasiswa.NPM)
		fmt.Println("Jurusan:", mahasiswa.Jurusan)
	} else {
		fmt.Printf("\nData Mahasiswa dengan NPM %s tidak ditemukan.\n", CariNPM)
	}
}
