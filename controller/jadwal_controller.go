package controller

import (
	"fmt"
	"jadwalku/handlers"
	"jadwalku/helper"
	"jadwalku/model"
)

func AddSchedule() {
	fmt.Println("===== Tambah Jadwal =====")
	jadwal := readScheduleInput()
	saved, err := handlers.AddSchedule(jadwal)
	if err != nil {
		helper.SetMessage("Gagal menambahkan jadwal: " + err.Error())
		return
	}

	helper.SetMessage(fmt.Sprintf("Jadwal berhasil ditambahkan dengan ID %d.", saved.ID))
}

func UpdateSchedule() {
	fmt.Println("===== Ubah Jadwal =====")
	schedules := handlers.GetSchedules()
	if len(schedules) == 0 {
		helper.SetMessage("Belum ada data jadwal.")
		return
	}

	displaySchedules(schedules)
	id := helper.ReadInt("Masukan ID jadwal yang ingin diubah: ")
	fmt.Println("Masukan data baru.")

	jadwal := readScheduleInput()
	updated, err := handlers.UpdateSchedule(id, jadwal)
	if err != nil {
		helper.SetMessage("Gagal mengubah jadwal: " + err.Error())
		return
	}

	helper.SetMessage(fmt.Sprintf("Jadwal ID %d berhasil diubah.", updated.ID))
}

func DeleteSchedule() {
	fmt.Println("===== Hapus Jadwal =====")
	schedules := handlers.GetSchedules()
	if len(schedules) == 0 {
		helper.SetMessage("Belum ada data jadwal.")
		return
	}

	displaySchedules(schedules)
	id := helper.ReadInt("Masukan ID jadwal yang ingin dihapus: ")
	err := handlers.DeleteSchedule(id)
	if err != nil {
		helper.SetMessage("Gagal menghapus jadwal: " + err.Error())
		return
	}

	helper.SetMessage(fmt.Sprintf("Jadwal ID %d berhasil dihapus.", id))
}

func ShowSchedules() {
	fmt.Println("===== Daftar Jadwal =====")
	displaySchedules(handlers.GetSchedules())
	helper.Pause()
}

func SequentialSearch() {
	fmt.Println("===== Sequential Search =====")
	keyword := helper.ReadRequired("Masukan nama mata kuliah atau dosen: ")
	result := handlers.SequentialSearch(keyword)
	displaySchedules(result)
	helper.Pause()
}

func BinarySearch() {
	fmt.Println("===== Binary Search =====")
	keyword := helper.ReadRequired("Masukan awalan/nama mata kuliah atau dosen: ")
	result := handlers.BinarySearch(keyword)
	displaySchedules(result)
	helper.Pause()
}

func ShowSelectionSort() {
	fmt.Println("===== Selection Sort =====")
	result := handlers.SelectionSortByStartTime()
	displaySchedules(result)
	helper.Pause()
}

func ShowInsertionSort() {
	fmt.Println("===== Insertion Sort =====")
	result := handlers.InsertionSortByStartTime()
	displaySchedules(result)
	helper.Pause()
}

func ShowStatistics() {
	fmt.Println("===== Statistik =====")
	day := helper.ReadDay("Masukan hari acuan: ")
	currentTime := helper.ReadTime("Masukan jam acuan (HH:MM): ")
	totalMinutes, remaining := handlers.GetStatistics(day, currentTime)

	fmt.Println("+++ JadwalKu +++")
	fmt.Printf("Total jam kuliah per minggu: %d jam %d menit\n", totalMinutes/60, totalMinutes%60)
	fmt.Printf("Jumlah mata kuliah tersisa pada hari %s setelah %s: %d\n", day, helper.FormatTime(currentTime), remaining)
	helper.Pause()
}

func readScheduleInput() model.Jadwal {
	return model.Jadwal{
		MataKuliah: helper.ReadRequired("Nama mata kuliah: "),
		Dosen:      helper.ReadRequired("Nama dosen: "),
		Ruangan:    helper.ReadRequired("Ruangan: "),
		Hari:       helper.ReadDay("Hari: "),
		JamMulai:   helper.ReadTime("Jam mulai (HH:MM): "),
		JamSelesai: helper.ReadTime("Jam selesai (HH:MM): "),
	}
}

func displaySchedules(schedules []model.Jadwal) {
	if len(schedules) == 0 {
		fmt.Println("Belum ada data jadwal.")
		return
	}

	helper.DisplayTableHeader()
	for _, jadwal := range schedules {
		fmt.Printf("| %-3d | %-24s | %-20s | %-10s | %-12s | %-11s | %-12s |\n",
			jadwal.ID,
			jadwal.MataKuliah,
			jadwal.Dosen,
			jadwal.Ruangan,
			jadwal.Hari,
			helper.FormatTime(jadwal.JamMulai),
			helper.FormatTime(jadwal.JamSelesai),
		)
	}
	helper.DisplayTableFooter()
}
