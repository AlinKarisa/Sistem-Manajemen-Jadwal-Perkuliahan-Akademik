package handlers

import "jadwalku/model"

func DeleteSchedule(id int) error {
	return model.Delete(id)
}
