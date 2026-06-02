package helper

import "testing"

func TestParseTime(t *testing.T) {
	got, err := ParseTime("07:30")
	if err != nil {
		t.Fatalf("ParseTime returned error: %v", err)
	}
	if got != 450 {
		t.Fatalf("ParseTime = %d, want 450", got)
	}
}

func TestParseTimeRejectsInvalidHour(t *testing.T) {
	if _, err := ParseTime("24:00"); err == nil {
		t.Fatal("ParseTime should reject invalid hour")
	}
}

func TestFormatTime(t *testing.T) {
	got := FormatTime(570)
	if got != "09:30" {
		t.Fatalf("FormatTime = %q, want %q", got, "09:30")
	}
}
