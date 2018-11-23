package goutils

import (
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
			case newDuration := <-timeout.changeDurationCh:
				timeout.timeoutCh = time.After(newDuration)
			case <-timeout.timeoutCh:
				f()
				return
			case <-timeout.stopCh:
				return
			}
		}
	}()

	return
}
