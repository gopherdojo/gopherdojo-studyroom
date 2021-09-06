package fortune

import "time"

func SetTime(t time.Time) {
	timeNow = func() time.Time { return t }
}
