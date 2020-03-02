package utils

import "time"

func FormatTime(t time.Time) string {
	return t.Format("2006.01.02-15.04.05") // format: YYYY.MM.DD-hh.mm.ss
}
