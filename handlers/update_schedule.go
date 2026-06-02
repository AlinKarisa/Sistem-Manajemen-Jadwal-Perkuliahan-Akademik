package handlers

import "jadwalku/model"

func UpdateSchedule(id int, jadwal model.Jadwal) (model.Jadwal, error) {
	return model.Update(id, jadwal)
}
