package goutils

import "time"

func RepeatTimer(d time.Duration, f func()) {
	go func() {
		for range time.Tick(d) {
			f()
		}
	}()
}
