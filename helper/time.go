package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseTime(value string) (int, error) {
	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("format jam harus HH:MM")
	}

	hour, hourErr := strconv.Atoi(parts[0])
	minute, minuteErr := strconv.Atoi(parts[1])
	if hourErr != nil || minuteErr != nil {
		return 0, fmt.Errorf("jam harus berisi angka")
	}

	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return 0, fmt.Errorf("jam harus berada pada rentang 00:00 sampai 23:59")
	}

	return hour*60 + minute, nil
}

func FormatTime(minutes int) string {
	hour := minutes / 60
	minute := minutes % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

func ReadTime(prompt string) int {
	for {
		value := ReadLine(prompt)
		minutes, err := ParseTime(value)
		if err == nil {
			return minutes
		}

		fmt.Println(err.Error())
	}
}

func ReadDay(prompt string) string {
	for {
		value := strings.ToLower(ReadRequired(prompt))
		switch value {
		case "senin":
			return "Senin"
		case "selasa":
			return "Selasa"
		case "rabu":
			return "Rabu"
		case "kamis":
			return "Kamis"
		case "jumat", "jum'at":
			return "Jumat"
		case "sabtu":
			return "Sabtu"
		case "minggu":
			return "Minggu"
		default:
			fmt.Println("Hari harus Senin, Selasa, Rabu, Kamis, Jumat, Sabtu, atau Minggu.")
		}
	}
}
