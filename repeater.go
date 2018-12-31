package goutils

import "time"

type Repeater struct {
	stopCh chan bool
	tick   <-chan time.Time
}

func NewRepeater(d time.Duration, f func()) (repeater *Repeater) {
	repeater = &Repeater{}
	repeater.stopCh = make(chan bool)
	repeater.tick = time.Tick(d)

	go func() {
		for {
			select {
			case <-repeater.tick:
				f()
			case <-repeater.stopCh:
				return
			}
		}
	}()

	return
}

func (self *Repeater) Stop() {
	self.stopCh <- true
}
