package goutils

import "time"

type Timeout struct {
	changeDurationCh chan time.Duration
	timeoutCh        <-chan time.Time
}

func (self *Timeout) ChangeTime(duration time.Duration) {
	self.changeDurationCh <- duration
}

func NewTimeout(duration time.Duration, f func()) (timeout *Timeout) {
	timeout = &Timeout{}
	timeout.changeDurationCh = make(chan time.Duration)

	timeout.timeoutCh = time.After(duration)

	go func() {
		for {
			select {
			case newDuration := <-timeout.changeDurationCh:
				timeout.timeoutCh = time.After(newDuration)
			case <-timeout.timeoutCh:
				f()
				return
			}
		}
	}()

	return
}
