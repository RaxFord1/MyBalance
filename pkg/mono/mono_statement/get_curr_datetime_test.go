package mono_statement

import (
	"testing"
)

func TestCurrDateTime(t *testing.T) {
	unixStartOfDay, now := GetTimeStartAndNowUnix()

	if unixStartOfDay > now {
		t.Fatal("unixStartOfDay > now")
	}
	t.Log(unixStartOfDay, now)
}
