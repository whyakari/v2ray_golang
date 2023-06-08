package sleep

import "time"

func Sleep(seconds int) {
	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)
}
