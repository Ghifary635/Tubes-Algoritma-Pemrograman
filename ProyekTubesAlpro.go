package main

import (
	"fmt"
)

type Peserta struct {
	Nama, AsalDaerah                                             string
	WawasanKebangsaan, PotensiAkademik, PengetahuanKorupsi, Skor float64
	ID_Peserta                                                   int
}

const final int = 10
const temp int = 50

type PesertaLolos [final]Peserta
type PenampungCalon [temp]Peserta
type Null [1]Peserta

func Home() {
	/*	I.S -
		F.S mencetak menu home setelah menginput data (Beranda Utama) yang berisi pilihan menu*/
	fmt.Println("\033[32m_____________ \033[0m")
	fmt.Println("\033[32;1m   H O M E    \033[0m")
	fmt.Println("\033[32m------------- \033[0m")
	fmt.Println("\033[33;1m1. Tambah Lembar Baru Calon Peserta\033[0m")
	fmt.Println("\033[33;1m2. Edit Data Calon Peserta\033[0m")
	fmt.Println("\033[33;1m3. Hapus Data Calon Peserta\033[0m")
	fmt.Println("\033[33;1m4. Tampilkan Peserta Yang Lolos\033[0m")
	fmt.Println("\033[33;1m5. Cari Data Peserta\033[0m")
	fmt.Println("\033[33;1m0.Keluar\033[0m")
}
func Home2() {
	/*	I.S -
		F.S mencetak menu home sebelum menginput data*/
	fmt.Println("\033[32m_____________\033[0m")
	fmt.Println("\033[32m   H O M E   \033[0m")
	fmt.Println("\033[32m-------------\033[0m")
	fmt.Println("\033[32mAnda belum memasukkan data sama sekali, Silahkan masukkan data terlebih dahulu\033[0m")
}
func Opsi1() {
	/*	I.S -
		F.S mencetak opsi lanjutan setelah user menginput opsi 1 pada beranda utama*/
	fmt.Println("11. Isi calon peserta pada lembar baru")
	fmt.Println("12. Tambah calon peserta pada lembar yang telah disimpan")
}

func Opsi2() {
	/*	I.S -
		F.S mencetak opsi lanjutan setelah user menginput opsi 2 pada beranda utama*/
	fmt.Println("21. Ubah nama calon peserta")
	fmt.Println("22. Ubah asal daerah calon peserta")
	fmt.Println("23. Ubah nilai calon peserta")
}

func Opsi3() {
	/*	I.S -
		F.S mencetak opsi lanjutan setelah user menginput opsi 3 pada beranda utama*/
	fmt.Println("31. Hapus data calon peserta")
	fmt.Println("32. Hapus nilai calon peserta saja")
}

func isValidNama(nama string) bool {
	/*	mengembalikan true jika nama tidak mengandung angka dan false jika sebaliknya*/
	for i := 0; i < len(nama); i++ {
		if nama[i] >= '0' && nama[i] <= '9' {
			return false
		}
	}
	return true
}

func isValidNilai(wawasan, potensi, pengetahuan float64) bool {
	/*	mengembalikan false jika ketiga nilai bernilai negatif dan true jika sebaliknya */
	return wawasan >= 0 && potensi >= 0 && pengetahuan >= 0
}

func isValidAsalDaerah(daerah string) bool {
	/*	mengembalikan true jika asal daerah yang diinput adalah JawaTimur, JawaTengah, JawaBarat, Banten, DKIJakarta, dan false jika selain itu*/
	validDaerah := []string{"JawaTimur", "JawaTengah", "JawaBarat", "Banten", "DKIJakarta"}
	for i := 0; i < len(validDaerah); i++ {
		if daerah == validDaerah[i] {
			return true
		}
	}
	return false
}

func Tambah_LembarBaru(A *PenampungCalon, n *int) {
	/*	I.S terdefinisi array A sebanyak n sembarang dan array A kosong
		F.S array A terisi sebanyak n dimana n <= 50*/
	var temp, i int
	for *n > 50 {
		fmt.Println("\033[31;1mAnda hanya dapat menambahkan maksimal 50 calon peserta, \nsilahkan masukkan ulang jumlah calon yang akan ditambahkan\033[0m")
		fmt.Scan(&temp)
		*n = temp
	}
	fmt.Println("\033[33;1mSilahkan masukkan data calon peserta\033[0m")
	fmt.Println("\033[33;1mContoh: Ujang_Ratau_Adus Banten 87 78 87\033[0m")
	for i = 0; i < *n; i++ {
		var nama, asalDaerah string
		var wawasanKebangsaan, potensiAkademik, pengetahuanKorupsi float64

		fmt.Scan(&nama, &asalDaerah, &wawasanKebangsaan, &potensiAkademik, &pengetahuanKorupsi)

		if !isValidNama(nama) {
			fmt.Println("\033[31;1mNama tidak valid, mengandung angka\033[0m")
			i--
			continue
		}
		if !isValidNilai(wawasanKebangsaan, potensiAkademik, pengetahuanKorupsi) {
			fmt.Println("\033[31;1mNilai tidak valid, kurang dari 0\033[0m")
			i--
			continue
		}
		if !isValidAsalDaerah(asalDaerah) {
			fmt.Println("\033[31;1mAsal daerah tidak valid\033[0m")
			i--
			continue
		}

		A[i].Nama = nama
		A[i].AsalDaerah = asalDaerah
		A[i].WawasanKebangsaan = wawasanKebangsaan
		A[i].PotensiAkademik = potensiAkademik
		A[i].PengetahuanKorupsi = pengetahuanKorupsi
		A[i].Skor = 0.3*A[i].WawasanKebangsaan + 0.3*A[i].PotensiAkademik + 0.6*A[i].PengetahuanKorupsi
		A[i].ID_Peserta = 545000 + i
	}
}

func Tambah_LembarLanjut(A *PenampungCalon, n *int, nadd *int) {
	/*	I.S terdefinisi array A sebanyak n dan nadd adalah tambahan data yang akan ditambah
		F.S array A bertambah sebanyak n + nadd dimana n + nadd <= 50*/
	var i, total, tempadd int
	total = *n + *nadd
	for total > 50 {
		fmt.Println("\033[31;1mMohon maaf, Kami hanya dapat menampung maksimal 50 data, Silahkan masukkan ulang jumlah data yang akan ditambahkan\033[0m")
		fmt.Scan(&tempadd)
		*nadd = tempadd
		total = *n + *nadd
	}
	for i = *n; i < total; i++ {
		var nama, asalDaerah string
		var wawasanKebangsaan, potensiAkademik, pengetahuanKorupsi float64

		fmt.Println("\033[33;1mSilahkan isi data calon\033[0m")
		fmt.Println("\033[33;1mContoh: Udin_Oli_Samping Banten 80 80 80\033[0m")
		fmt.Scan(&nama, &asalDaerah, &wawasanKebangsaan, &potensiAkademik, &pengetahuanKorupsi)

		if !isValidNama(nama) {
			fmt.Println("\033[31;1mNama tidak valid, mengandung angka\033[0m")
			i--
			continue
		}
		if !isValidNilai(wawasanKebangsaan, potensiAkademik, pengetahuanKorupsi) {
			fmt.Println("\033[31;1mNilai tidak valid, kurang dari 0\033[0m")
			i--
			continue
		}
		if !isValidAsalDaerah(asalDaerah) {
			fmt.Println("\033[31;1mAsal daerah tidak valid\033[0m")
			i--
			continue
		}

		A[i].Nama = nama
		A[i].AsalDaerah = asalDaerah
		A[i].WawasanKebangsaan = wawasanKebangsaan
		A[i].PotensiAkademik = potensiAkademik
		A[i].PengetahuanKorupsi = pengetahuanKorupsi
		A[i].Skor = 0.3*A[i].WawasanKebangsaan + 0.3*A[i].PotensiAkademik + 0.6*A[i].PengetahuanKorupsi
		A[i].ID_Peserta = 545000 + i
	}
	*n = total
}

func Edit_AsalDaerah(A *PenampungCalon, n int) {
	/*	I.S terdefinisi array A sebanyak n
		F.S mengubah data asal daerah dari array A*/
	var i, id, warn int
	fmt.Println("\033[33;1mMasukkan ID Peserta untuk mengubah asal daerah calon peserta\033[0m")
	fmt.Println("\033[33;1mContoh: 545999\033[0m")
	fmt.Scan(&id)
	warn = -1
	for i = 0; i < n; i++ {
		if A[i].ID_Peserta == id {
			fmt.Println("\033[33;1mAsal daerah calon peserta sebelum diubah:\033[0m", A[i].AsalDaerah)
			fmt.Println("\033[33;1mAsal daerah yang valid (JawaTimur/JawaBarat/JawaTengah/Bnaten/DKIJakarta)\033[0m")
			fmt.Print("\033[33;1mMasukkan Data: \033[0m")
			fmt.Scan(&A[i].AsalDaerah)
			isValidAsalDaerah(A[i].AsalDaerah)
			if !isValidAsalDaerah(A[i].AsalDaerah) {
				fmt.Println("\033[31;1mAsal daerah tidak valid\033[0m")
				i--
				continue
			}
			warn = 1
		}
	}
	if warn == -1 {
		fmt.Println("\033[31;1mID tidak valid\033[0m")
	}
}
func Edit_Nama(A *PenampungCalon, n int) {
	/*	I.S terdefinisi array A sebanyak n
		F.S mengubah data nama dari array A*/
	var i, id, warn int
	fmt.Println("\033[33;1mMasukkan ID Peserta untuk mengubah nama calon peserta\033[0m")
	fmt.Println("\033[33;1mContoh: 545999\033[0m")
	fmt.Scan(&id)
	warn = -1
	for i = 0; i < n; i++ {
		if A[i].ID_Peserta == id {
			fmt.Println("Nama calon peserta sebelum diubah:", A[i].Nama)
			fmt.Print("\033[33;1mMasukkan Anda: \033[0m")
			fmt.Scan(&A[i].Nama)
			isValidNama(A[i].Nama)
			if !isValidNama(A[i].Nama) {
				fmt.Println("\033[31;1mNama tidak valid, mengandung angka\033[0m")
				i--
				continue
			}
			warn = 1
		}
	}
	if warn == -1 {
		fmt.Println("\033[31;1mID tidak valid\033[0m")
	}
}
func Edit_Nilai(A *PenampungCalon, n int) {
	/*	I.S terdefinisi array A sebanyak n
		F.S mengubah data nilai dari array A*/
	var i, id, warn int
	fmt.Println("\033[33;1mMasukkan ID Peserta untuk mengubah nilai peserta\033[0m")
	fmt.Scan(&id)
	warn = -1
	for i = 0; i < n; i++ {
		if A[i].ID_Peserta == id {
			fmt.Println("Nilai calon peserta sebelum diubah:", A[i].WawasanKebangsaan, A[i].PotensiAkademik, A[i].PengetahuanKorupsi)
			fmt.Print("\033[33;1mMasukkan Anda: \033[0m")
			fmt.Scan(&A[i].WawasanKebangsaan, &A[i].PotensiAkademik, &A[i].PengetahuanKorupsi)
			A[i].Skor = 0.3*A[i].WawasanKebangsaan + 0.3*A[i].PotensiAkademik + 0.6*A[i].PengetahuanKorupsi
			isValidNilai(A[i].WawasanKebangsaan, A[i].PotensiAkademik, A[i].PengetahuanKorupsi)
			if !isValidNilai(A[i].WawasanKebangsaan, A[i].PotensiAkademik, A[i].PengetahuanKorupsi) {
				fmt.Println("\033[31;1mNilai tidak valid, kurang dari 0\033[0m")
				i--
				continue
			}
			warn = 1
		}
	}
	if warn == -1 {
		fmt.Println("\033[31;1mID tidak valid\033[0m")
	}
}
func Hapus_AllData(A *PenampungCalon, n *int) {
	/*	I.S terdefinisi array A sebanyak n
		F.S menghapus keseluruhan data yang dimiliki oleh seorang peserta, dan nilai n berkurang 1*/
	var i, id, idx int
	fmt.Println("\033[33;1mMasukkan ID Peserta pada data calon peserta yang akan dihapus\033[0m")
	fmt.Println("\033[33;1mContoh: 545999\033[0m")
	fmt.Scan(&id)
	for i = 0; i < *n; i++ {
		if A[i].ID_Peserta == id {
			idx = i
			A[i].ID_Peserta = -99999
		}
	}
	for i = idx; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n = *n - 1
}
func Hapus_Nilai(A *PenampungCalon, n int) {
	/*	I.S terdefinisi array A sebanyak n
		F.S menghapus data nilai yang dimiliki oleh seorang peserta sehingga nilai peserta menjadi 0*/
	var i, id int
	fmt.Println("\033[33;1mMasukkan ID Peserta pada data calon peserta yang nilainya akan dihapus\033[0m")
	fmt.Println("\033[33;1mContoh: 545999\033[0m")
	fmt.Scan(&id)
	for i = 0; i < n; i++ {
		if A[i].ID_Peserta == id {
			A[i].WawasanKebangsaan = 0
			A[i].PotensiAkademik = 0
			A[i].PengetahuanKorupsi = 0
			A[i].Skor = 0
		}
	}
}

func Lolos(Z *PesertaLolos, A PenampungCalon, n int) {
	/*	I.S terdefinisi array A sebanyak n dan array Z kosong
		F.S mencetak array Z yang berisi 10 peserta yang lolos seleksi
		berdasarkan nilai skor tertinggi tiap daerah dimana tiap daerah mendapat 2 kuota */
	var i, x, y, jatim, jateng, jabar, btn, dki int
	var skor float64
	x = 0
	jatim = 0
	jateng = 0
	jabar = 0
	dki = 0
	btn = 0
	for i = 0; i < n; i++ {
		if A[i].AsalDaerah == "JawaTimur" {
			jatim = jatim + 1
		} else if A[i].AsalDaerah == "JawaTengah" {
			jateng = jateng + 1
		} else if A[i].AsalDaerah == "JawaBarat" {
			jabar = jabar + 1
		} else if A[i].AsalDaerah == "DKIJakarta" {
			dki = dki + 1
		} else if A[i].AsalDaerah == "Banten" {
			btn = btn + 1
		}
	}

	if jatim != 0 {
		skor = -1
		for i = 0; i < n; i++ {
			if A[i].AsalDaerah == "JawaTimur" && A[i].Skor >= skor {
				skor = A[i].Skor
				y = i
			}
		}
		Z[x] = A[y]
		x = x + 1
	}
	if jatim > 1 {
		skor = -1
		for i = 0; i < n; i++ {
			if A[i].AsalDaerah == "JawaTimur" && A[i].Skor >= skor && A[i].ID_Peserta != Z[x-1].ID_Peserta {
				skor = A[i].Skor
				y = i
			}
		}
		Z[x] = A[y]
		x = x + 1
	}

	if jateng != 0 {
		skor = -1
		for i = 0; i < n; i++ {
			if A[i].AsalDaerah == "JawaTengah" && A[i].Skor >= skor {
				skor = A[i].Skor
				y = i
			}
		}
		Z[x] = A[y]
		x = x + 1
	}
	if jateng > 1 {
		skor = -1
		for i = 0; i < n; i++ {
			if A[i].AsalDaerah == "JawaTengah" && A[i].Skor >= skor && A[i].ID_Peserta != Z[x-1].ID_Peserta {
				skor = A[i].Skor
				y = i
			}
		}
		Z[x] = A[y]
		x = x + 1
	}

	if jabar != 0 {
		skor = -1
		for i = 0; i < n; i++ {
			if A[i].AsalDaerah == "JawaBarat" && A[i].Skor >= skor {
				skor = A[i].Skor
				y = i
			}
		}
		Z[x] = A[y]
		x = x + 1
	}
	if jabar > 1 {
		skor = -1
		for i = 0; i < n; i++ {
			if A[i].AsalDaerah == "JawaBarat" && A[i].Skor >= skor && A[i].ID_Peserta != Z[x-1].ID_Peserta {
				skor = A[i].Skor
				y = i
			}
		}
		Z[x] = A[y]
		x = x + 1
	}

	if dki != 0 {
		skor = -1
		for i = 0; i < n; i++ {
			if A[i].AsalDaerah == "DKIJakarta" && A[i].Skor >= skor {
				skor = A[i].Skor
				y = i
			}
		}
		Z[x] = A[y]
		x = x + 1
	}
	if dki > 1 {
		skor = -1
		for i = 0; i < n; i++ {
			if A[i].AsalDaerah == "DKIJakarta" && A[i].Skor >= skor && A[i].ID_Peserta != Z[x-1].ID_Peserta {
				skor = A[i].Skor
				y = i
			}
		}
		Z[x] = A[y]
		x = x + 1
	}

	if btn != 0 {
		skor = -1
		for i = 0; i < n; i++ {
			if A[i].AsalDaerah == "Banten" && A[i].Skor >= skor {
				skor = A[i].Skor
				y = i
			}
		}
		Z[x] = A[y]
		x = x + 1
	}
	if btn > 1 {
		skor = -1
		for i = 0; i < n; i++ {
			if A[i].AsalDaerah == "Banten" && A[i].Skor >= skor && A[i].ID_Peserta != Z[x-1].ID_Peserta {
				skor = A[i].Skor
				y = i
			}
		}
		Z[x] = A[y]
		x = x + 1
	}
	fmt.Printf("%20s %15s %15s %15s %15s %15s %15s", "Nama", "ID_Peserta", "Asal Daerah", "Nilai WK", "Nilai PA", "Nilai PK", "Skor Keseluruhan")
	fmt.Println()
	for i = 0; i < x; i++ {
		fmt.Printf("%20s %15d %15s %15.2f %15.2f %15.2f %15.2f", Z[i].Nama, Z[i].ID_Peserta, Z[i].AsalDaerah, Z[i].WawasanKebangsaan, Z[i].PotensiAkademik, Z[i].PengetahuanKorupsi, Z[i].Skor)
		fmt.Println()
	}
}
func cari(A PenampungCalon, n int, namacalon, daerahcalon string) {
	/*	I.S terdefinisi array A sebanyak n serta nama calon dan daerah asal calon yang akan dicari datanya
		F.S menampilkan keseluruhan data dari peserta yang dicari*/
	var i, dupe, j, idcalon int
	dupe = 0
	for i = 0; i < n; i++ {
		if A[i].Nama == namacalon && A[i].AsalDaerah == daerahcalon {
			dupe = dupe + 1
			j = i
		}
	}
	if dupe == 0 {
		fmt.Println("\033[33;1mData peserta tidak ditemukan\033[0m")
	} else if dupe == 1 {
		fmt.Println("Data Calon Peserta:")
		fmt.Println("Nama:", A[j].Nama)
		fmt.Println("ID:", A[j].ID_Peserta)
		fmt.Println("Asal Daerah:", A[j].AsalDaerah)
		fmt.Println("Skor Tes Wawasan Kebangsaan:", A[j].WawasanKebangsaan)
		fmt.Println("Skor Tes Potensi Akademik:", A[j].PotensiAkademik)
		fmt.Println("Skor Tes Pengetahuan Korupsi:", A[j].PengetahuanKorupsi)
		fmt.Printf("Skor Keseluruhan: %.2f", A[j].Skor)
		fmt.Println()
	} else {
		fmt.Println("Ditemukan data peserta lebih dari 1, silahkan masukkan id peserta")
		fmt.Scan(&idcalon)
		if idcalon >= 545000 {
			for i = 0; i < n; i++ {
				if A[i].ID_Peserta == idcalon {
					j = i
					fmt.Println("Data Calon Peserta:")
					fmt.Println("Nama:", A[i].Nama)
					fmt.Println("ID:", A[i].ID_Peserta)
					fmt.Println("Asal Daerah:", A[i].AsalDaerah)
					fmt.Println("Skor Tes Wawasan Kebangsaan:", A[i].WawasanKebangsaan)
					fmt.Println("Skor Tes Potensi Akademik:", A[i].PotensiAkademik)
					fmt.Println("Skor Tes Pengetahuan Korupsi:", A[i].PengetahuanKorupsi)
					fmt.Printf("Skor Keseluruhan: %.2f", A[i].Skor)
					fmt.Println()
				}
			}
		}
	}
}

func main() {
	var pilih, ndata, ntambahan int
	var X PenampungCalon
	var Z PesertaLolos
	var namaCalon, daerahCalon string
	Home2()
	fmt.Println("\033[33;1mSilahkan masukkan jumlah data calon peserta yang akan dimasukkan\033[0m")
	fmt.Scan(&ndata)
	Tambah_LembarBaru(&X, &ndata)
	Home()
	fmt.Scan(&pilih)
	for pilih != 0 {
		if pilih == 1 {
			Opsi1()
			fmt.Scan(&pilih)
			if pilih == 11 {
				fmt.Println("\033[33;1mSilahkan masukkan jumlah data calon peserta yang akan dimasukkan\033[0m")
				fmt.Scan(&ndata)
				Tambah_LembarBaru(&X, &ndata)
			} else if pilih == 12 {
				fmt.Println("\033[33;1mMasukkan jumlah data yang akan ditambahkan\033[0m")
				fmt.Scan(&ntambahan)
				Tambah_LembarLanjut(&X, &ndata, &ntambahan)
			} else {
				fmt.Println("\033[31;1mNomor menu yang anda masukkan salah, program segera ke menu\033[0m")
			}

		} else if pilih == 2 {
			Opsi2()
			fmt.Scan(&pilih)
			if pilih == 21 {
				Edit_Nama(&X, ndata)
			} else if pilih == 22 {
				Edit_AsalDaerah(&X, ndata)
			} else if pilih == 23 {
				Edit_Nilai(&X, ndata)
			} else {
				fmt.Println("\033[31;1mNomor menu yang anda masukkan salah, program segera ke menu\033[0m")
			}
		} else if pilih == 3 {
			Opsi3()
			fmt.Scan(&pilih)
			if pilih == 31 {
				Hapus_AllData(&X, &ndata)
			} else if pilih == 32 {
				Hapus_Nilai(&X, ndata)
			} else {
				fmt.Println("\033[31;1mNomor menu yang anda masukkan salah, program segera ke menu\033[0m")
			}
		} else if pilih == 4 {
			Lolos(&Z, X, ndata)
		} else if pilih == 5 {
			fmt.Println("\033[33;1mSilahkan masukkan nama dan daerah calon peserta\033[0m")
			fmt.Print("\033[33;1mNama: \033[0m")
			fmt.Scan(&namaCalon)
			fmt.Print("\033[33;1mAsal Daerah: \033[0m")
			fmt.Scan(&daerahCalon)
			cari(X, ndata, namaCalon, daerahCalon)
		} else {
			fmt.Println("\033[31;1mNomor menu yang anda masukkan salah, program segera ke menu\033[0m")
		}
		Home()
		fmt.Scan(&pilih)
	}
	fmt.Println("Terima kasih telah menggunakan layanan kami")
}
