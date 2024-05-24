package mono_statement

import (
	"log"
	"time"
)

func GetTimeStartAndNowUnix() (int64, int64) {
	now := time.Now()

	// Create a time object for the start of the current day (midnight)
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Get the Unix timestamp for the start of the day
	unixStartOfDay := startOfDay.Unix()

	log.Println(unixStartOfDay)
	log.Println(now.Unix())

	return unixStartOfDay, now.Unix()
}
