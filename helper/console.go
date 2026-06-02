package helper

import "fmt"

var message string

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func SetMessage(value string) {
	message = value
}

func ShowMessage() {
	if message == "" {
		return
	}

	fmt.Println(message)
	message = ""
}

func DisplayMainMenu() {
	ShowMessage()
	fmt.Println("+++ JadwalKu +++")
	fmt.Println("1. Tambah Jadwal")
	fmt.Println("2. Ubah Jadwal")
	fmt.Println("3. Hapus Jadwal")
	fmt.Println("4. Tampilkan Jadwal")
	fmt.Println("5. Sequential Search")
	fmt.Println("6. Binary Search")
	fmt.Println("7. Selection Sort")
	fmt.Println("8. Insertion Sort")
	fmt.Println("9. Statistik")
	fmt.Println("10. Keluar")
	fmt.Println("=================")
}

func DisplayTableHeader() {
	fmt.Println("==============================================================================================================")
	fmt.Printf("| %-3s | %-24s | %-20s | %-10s | %-12s | %-11s | %-12s |\n", "ID", "Mata Kuliah", "Dosen", "Ruangan", "Hari", "Jam Mulai", "Jam Selesai")
	fmt.Println("==============================================================================================================")
}

func DisplayTableFooter() {
	fmt.Println("==============================================================================================================")
}
