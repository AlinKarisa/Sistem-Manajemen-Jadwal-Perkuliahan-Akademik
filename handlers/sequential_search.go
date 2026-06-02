package handlers

import "jadwalku/model"

func SequentialSearch(keyword string) []model.Jadwal {
	return model.SequentialSearch(keyword)
}
