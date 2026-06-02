package handlers

import "jadwalku/model"

func GetSchedules() []model.Jadwal {
	return model.GetAll()
}
