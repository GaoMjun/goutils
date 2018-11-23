package goutils

import (
	"fmt"
	"time"
)

type Timeout struct {
	changeDurationCh chan time.Duration
	timeoutCh        <-chan time.Time
	stopCh           chan bool
}

func (self *Timeout) ChangeTime(duration time.Duration) {
	self.changeDurationCh <- duration
}

func (self *Timeout) Stop() {
	self.stopCh <- true
}

func NewTimeout(duration time.Duration, f func()) (timeout *Timeout) {
	timeout = &Timeout{}
	timeout.changeDurationCh = make(chan time.Duration)
	timeout.stopCh = make(chan bool)

	timeout.timeoutCh = time.After(duration)

	go func() {
		for {
			select {
			case newDuration, ok := <-timeout.changeDurationCh:
				if !ok {
					return
				}
				timeout.timeoutCh = time.After(newDuration)
			case <-timeout.timeoutCh:
				f()
				return
			case <-timeout.stopCh:
				fmt.Println("timeout.stopCh")
				close(timeout.changeDurationCh)
				return
			}
		}
	}()

	return
}
