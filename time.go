package NoteSystem

import (
	"strconv"
	"time"
)

func timeStr() string {
	now := time.Now().UTC()
	return strconv.Itoa(now.Day()) + "." + strconv.Itoa(int(now.Month())) + "." + strconv.Itoa(now.Year()) +
		"_" + strconv.Itoa(now.Hour()) + "-" + strconv.Itoa(now.Minute())
}
