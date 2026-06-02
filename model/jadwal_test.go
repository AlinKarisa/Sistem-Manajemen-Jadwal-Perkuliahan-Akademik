package model

import "testing"

func TestAddRejectsRoomConflict(t *testing.T) {
	ResetForTest()
	_, err := Add(Jadwal{
		MataKuliah: "Algoritma",
		Dosen:      "Dosen A",
		Ruangan:    "R101",
		Hari:       "Senin",
		JamMulai:   480,
		JamSelesai: 600,
	})
	if err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	_, err = Add(Jadwal{
		MataKuliah: "Basis Data",
		Dosen:      "Dosen B",
		Ruangan:    "R101",
		Hari:       "Senin",
		JamMulai:   540,
		JamSelesai: 660,
	})
	if err == nil {
		t.Fatal("Add should reject room conflict")
	}
}

func TestSequentialSearchFindsCourseAndLecturer(t *testing.T) {
	ResetForTest()
	Add(Jadwal{MataKuliah: "Algoritma", Dosen: "Dosen A", Ruangan: "R101", Hari: "Senin", JamMulai: 480, JamSelesai: 600})
	Add(Jadwal{MataKuliah: "Basis Data", Dosen: "Dosen B", Ruangan: "R102", Hari: "Selasa", JamMulai: 600, JamSelesai: 720})

	courseResult := SequentialSearch("algo")
	if len(courseResult) != 1 || courseResult[0].MataKuliah != "Algoritma" {
		t.Fatalf("SequentialSearch course result = %+v", courseResult)
	}

	lecturerResult := SequentialSearch("dosen b")
	if len(lecturerResult) != 1 || lecturerResult[0].Dosen != "Dosen B" {
		t.Fatalf("SequentialSearch lecturer result = %+v", lecturerResult)
	}
}

func TestBinarySearchFindsCourseAndLecturerPrefix(t *testing.T) {
	ResetForTest()
	Add(Jadwal{MataKuliah: "Algoritma", Dosen: "Dosen A", Ruangan: "R101", Hari: "Senin", JamMulai: 480, JamSelesai: 600})
	Add(Jadwal{MataKuliah: "Basis Data", Dosen: "Dosen B", Ruangan: "R102", Hari: "Selasa", JamMulai: 600, JamSelesai: 720})

	courseResult := BinarySearch("Algo")
	if len(courseResult) != 1 || courseResult[0].MataKuliah != "Algoritma" {
		t.Fatalf("BinarySearch course result = %+v", courseResult)
	}

	lecturerResult := BinarySearch("Dosen B")
	if len(lecturerResult) != 1 || lecturerResult[0].Dosen != "Dosen B" {
		t.Fatalf("BinarySearch lecturer result = %+v", lecturerResult)
	}
}

func TestSelectionAndInsertionSortByStartTime(t *testing.T) {
	ResetForTest()
	Add(Jadwal{MataKuliah: "Siang", Dosen: "Dosen A", Ruangan: "R101", Hari: "Senin", JamMulai: 780, JamSelesai: 840})
	Add(Jadwal{MataKuliah: "Pagi", Dosen: "Dosen B", Ruangan: "R102", Hari: "Senin", JamMulai: 420, JamSelesai: 480})
	Add(Jadwal{MataKuliah: "Sore", Dosen: "Dosen C", Ruangan: "R103", Hari: "Senin", JamMulai: 900, JamSelesai: 960})

	selection := SelectionSortByStartTime()
	if selection[0].MataKuliah != "Pagi" || selection[1].MataKuliah != "Siang" || selection[2].MataKuliah != "Sore" {
		t.Fatalf("SelectionSortByStartTime result = %+v", selection)
	}

	insertion := InsertionSortByStartTime()
	if insertion[0].MataKuliah != "Pagi" || insertion[1].MataKuliah != "Siang" || insertion[2].MataKuliah != "Sore" {
		t.Fatalf("InsertionSortByStartTime result = %+v", insertion)
	}
}

func TestStatistics(t *testing.T) {
	ResetForTest()
	Add(Jadwal{MataKuliah: "Pagi", Dosen: "Dosen A", Ruangan: "R101", Hari: "Senin", JamMulai: 480, JamSelesai: 600})
	Add(Jadwal{MataKuliah: "Siang", Dosen: "Dosen B", Ruangan: "R102", Hari: "Senin", JamMulai: 780, JamSelesai: 900})

	if got := TotalWeeklyMinutes(); got != 240 {
		t.Fatalf("TotalWeeklyMinutes = %d, want 240", got)
	}

	if got := RemainingInDay("Senin", 600); got != 1 {
		t.Fatalf("RemainingInDay = %d, want 1", got)
	}
}
