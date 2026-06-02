package model

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Jadwal struct {
	ID         int
	MataKuliah string
	Dosen      string
	Ruangan    string
	Hari       string
	JamMulai   int
	JamSelesai int
}

var daftarJadwal = []Jadwal{
	{ID: 1, MataKuliah: "Algoritma Pemrograman", Dosen: "Budi Santoso", Ruangan: "R101", Hari: "Senin", JamMulai: 480, JamSelesai: 600},
	{ID: 2, MataKuliah: "Basis Data", Dosen: "Siti Aminah", Ruangan: "Lab 1", Hari: "Senin", JamMulai: 630, JamSelesai: 750},
	{ID: 3, MataKuliah: "Struktur Data", Dosen: "Rudi Hartono", Ruangan: "R202", Hari: "Selasa", JamMulai: 540, JamSelesai: 660},
	{ID: 4, MataKuliah: "Pemrograman Web", Dosen: "Dewi Lestari", Ruangan: "Lab 2", Hari: "Rabu", JamMulai: 780, JamSelesai: 900},
	{ID: 5, MataKuliah: "Sistem Operasi", Dosen: "Agus Wijaya", Ruangan: "R303", Hari: "Kamis", JamMulai: 480, JamSelesai: 600},
}
var nextID = len(daftarJadwal) + 1

func GetAll() []Jadwal {
	result := make([]Jadwal, len(daftarJadwal))
	copy(result, daftarJadwal)
	return result
}

func Add(jadwal Jadwal) (Jadwal, error) {
	if err := validate(jadwal, 0); err != nil {
		return Jadwal{}, err
	}

	jadwal.ID = nextID
	nextID++
	daftarJadwal = append(daftarJadwal, jadwal)
	return jadwal, nil
}

func Update(id int, jadwal Jadwal) (Jadwal, error) {
	index := findIndexByID(id)
	if index == -1 {
		return Jadwal{}, fmt.Errorf("jadwal dengan ID %d tidak ditemukan", id)
	}

	if err := validate(jadwal, id); err != nil {
		return Jadwal{}, err
	}

	jadwal.ID = id
	daftarJadwal[index] = jadwal
	return jadwal, nil
}

func Delete(id int) error {
	index := findIndexByID(id)
	if index == -1 {
		return fmt.Errorf("jadwal dengan ID %d tidak ditemukan", id)
	}

	daftarJadwal = append(daftarJadwal[:index], daftarJadwal[index+1:]...)
	return nil
}

func FindByID(id int) (Jadwal, bool) {
	index := findIndexByID(id)
	if index == -1 {
		return Jadwal{}, false
	}

	return daftarJadwal[index], true
}

func SequentialSearch(keyword string) []Jadwal {
	var result []Jadwal
	normalizedKeyword := normalize(keyword)
	if normalizedKeyword == "" {
		return result
	}

	for _, jadwal := range daftarJadwal {
		if strings.Contains(normalize(jadwal.MataKuliah), normalizedKeyword) || strings.Contains(normalize(jadwal.Dosen), normalizedKeyword) {
			result = append(result, jadwal)
		}
	}

	return result
}

func BinarySearch(keyword string) []Jadwal {
	normalizedKeyword := normalize(keyword)
	if normalizedKeyword == "" {
		return nil
	}

	resultByID := map[int]Jadwal{}

	byCourse := GetAll()
	sort.Slice(byCourse, func(i, j int) bool {
		return normalize(byCourse[i].MataKuliah) < normalize(byCourse[j].MataKuliah)
	})
	collectBinaryPrefix(byCourse, normalizedKeyword, func(jadwal Jadwal) string {
		return normalize(jadwal.MataKuliah)
	}, resultByID)

	byLecturer := GetAll()
	sort.Slice(byLecturer, func(i, j int) bool {
		return normalize(byLecturer[i].Dosen) < normalize(byLecturer[j].Dosen)
	})
	collectBinaryPrefix(byLecturer, normalizedKeyword, func(jadwal Jadwal) string {
		return normalize(jadwal.Dosen)
	}, resultByID)

	if len(resultByID) == 0 {
		return nil
	}

	var result []Jadwal
	for _, jadwal := range resultByID {
		result = append(result, jadwal)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func SelectionSortByStartTime() []Jadwal {
	sorted := GetAll()
	n := len(sorted)

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if compareSchedule(sorted[j], sorted[minIndex]) < 0 {
				minIndex = j
			}
		}
		if minIndex != i {
			sorted[i], sorted[minIndex] = sorted[minIndex], sorted[i]
		}
	}

	return sorted
}

func InsertionSortByStartTime() []Jadwal {
	sorted := GetAll()

	for i := 1; i < len(sorted); i++ {
		jadwalEel := sorted[i]
		j := i - 1

		for j >= 0 && compareSchedule(sorted[j], jadwalEel) > 0 {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = jadwalEel
	}

	return sorted
}

func TotalWeeklyMinutes() int {
	total := 0
	for _, jadwal := range daftarJadwal {
		total += jadwal.JamSelesai - jadwal.JamMulai
	}
	return total
}

func RemainingInDay(day string, currentMinutes int) int {
	count := 0
	for _, jadwal := range daftarJadwal {
		if strings.EqualFold(jadwal.Hari, day) && jadwal.JamMulai >= currentMinutes {
			count++
		}
	}
	return count
}

func ResetForTest() {
	daftarJadwal = nil
	nextID = 1
}

func validate(jadwal Jadwal, ignoredID int) error {
	if strings.TrimSpace(jadwal.MataKuliah) == "" {
		return errors.New("nama mata kuliah tidak boleh kosong")
	}
	if strings.TrimSpace(jadwal.Dosen) == "" {
		return errors.New("nama dosen tidak boleh kosong")
	}
	if strings.TrimSpace(jadwal.Ruangan) == "" {
		return errors.New("nama ruangan tidak boleh kosong")
	}
	if strings.TrimSpace(jadwal.Hari) == "" {
		return errors.New("hari tidak boleh kosong")
	}
	if jadwal.JamSelesai <= jadwal.JamMulai {
		return errors.New("jam selesai harus lebih besar dari jam mulai")
	}
	if conflict := findConflict(jadwal, ignoredID); conflict != nil {
		return fmt.Errorf("jadwal bentrok dengan ID %d pada hari %s", conflict.ID, conflict.Hari)
	}

	return nil
}

func findConflict(candidate Jadwal, ignoredID int) *Jadwal {
	for i := range daftarJadwal {
		current := &daftarJadwal[i]
		if current.ID == ignoredID || !strings.EqualFold(current.Hari, candidate.Hari) {
			continue
		}

		sameRoom := strings.EqualFold(current.Ruangan, candidate.Ruangan)
		sameLecturer := strings.EqualFold(current.Dosen, candidate.Dosen)
		overlap := candidate.JamMulai < current.JamSelesai && candidate.JamSelesai > current.JamMulai

		if overlap && (sameRoom || sameLecturer) {
			return current
		}
	}

	return nil
}

func findIndexByID(id int) int {
	for i, jadwal := range daftarJadwal {
		if jadwal.ID == id {
			return i
		}
	}

	return -1
}

func compareSchedule(left, right Jadwal) int {
	if left.JamMulai != right.JamMulai {
		return left.JamMulai - right.JamMulai
	}
	return strings.Compare(normalize(left.MataKuliah), normalize(right.MataKuliah))
}

func collectBinaryPrefix(sorted []Jadwal, keyword string, field func(Jadwal) string, result map[int]Jadwal) {
	index := lowerBound(sorted, keyword, field)
	for i := index; i < len(sorted); i++ {
		if !strings.HasPrefix(field(sorted[i]), keyword) {
			break
		}
		result[sorted[i].ID] = sorted[i]
	}
}

func lowerBound(sorted []Jadwal, keyword string, field func(Jadwal) string) int {
	low := 0
	high := len(sorted)

	for low < high {
		mid := (low + high) / 2
		if field(sorted[mid]) < keyword {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func normalize(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}
