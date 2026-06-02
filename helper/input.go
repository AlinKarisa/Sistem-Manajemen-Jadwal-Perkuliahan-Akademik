package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadLine(prompt string) string {
	fmt.Print(prompt)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func ReadRequired(prompt string) string {
	for {
		value := ReadLine(prompt)
		if value != "" {
			return value
		}

		fmt.Println("Input tidak boleh kosong.")
	}
}

func ReadInt(prompt string) int {
	for {
		value := ReadLine(prompt)
		number, err := strconv.Atoi(value)
		if err == nil {
			return number
		}

		fmt.Println("Masukan harus berupa angka.")
	}
}

func ReadIntRange(prompt string, minValue, maxValue int) int {
	for {
		value := ReadInt(prompt)
		if value >= minValue && value <= maxValue {
			return value
		}

		fmt.Printf("Pilihan harus antara %d sampai %d.\n", minValue, maxValue)
	}
}

func Pause() {
	fmt.Println("Tekan Enter untuk kembali ke menu...")
	reader.ReadString('\n')
}
