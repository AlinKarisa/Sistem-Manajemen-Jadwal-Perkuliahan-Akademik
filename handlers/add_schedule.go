package handlers

import "jadwalku/model"

func AddSchedule(jadwal model.Jadwal) (model.Jadwal, error) {
	return model.Add(jadwal)
}
