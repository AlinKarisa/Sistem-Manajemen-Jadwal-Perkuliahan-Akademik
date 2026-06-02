package main

import (
	"fmt"
	"jadwalku/controller"
	"jadwalku/helper"
)

func main() {
	var choice int

	for isRunning := false; !isRunning; {
		helper.ClearScreen()
		helper.DisplayMainMenu()

		choice = helper.ReadIntRange("Masukan Pilihan Anda : ", 1, 10)
		helper.ClearScreen()

		switch choice {
		case 1:
			controller.AddSchedule()
		case 2:
			controller.UpdateSchedule()
		case 3:
			controller.DeleteSchedule()
		case 4:
			controller.ShowSchedules()
		case 5:
			controller.SequentialSearch()
		case 6:
			controller.BinarySearch()
		case 7:
			controller.ShowSelectionSort()
		case 8:
			controller.ShowInsertionSort()
		case 9:
			controller.ShowStatistics()
		case 10:
			isRunning = true
			fmt.Println("Terima kasih telah menggunakan JadwalKu.")
		}

		if choice != 10 {
			pilihanEel := 0
			choice = pilihanEel
		}
	}
}
