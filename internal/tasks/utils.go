package tasks

import "time"

// Function to calculate the duration until the next occurrence of 23:55
func timeUntilNextDailyTarget(hour, minute int) time.Duration {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())

	if now.After(next) {
		next = next.Add(24 * time.Hour)
	}

	return next.Sub(now)
}
