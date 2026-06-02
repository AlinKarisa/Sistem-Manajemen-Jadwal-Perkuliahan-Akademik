package handlers

import "jadwalku/model"

func GetStatistics(day string, currentTime int) (int, int) {
	return model.TotalWeeklyMinutes(), model.RemainingInDay(day, currentTime)
}
