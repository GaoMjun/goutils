package goutils

import (
	"errors"
	"os"
	"time"
)

func RepeatTimer(d time.Duration, f func()) {
	go func() {
		for range time.Tick(d) {
			f()
		}
	}()
}

func CreateFileNotExist(filename string) (f *os.File, err error) {
	if _, err = os.Stat(filename); !os.IsNotExist(err) {
		err = errors.New("file is exist")
		return
	}
	
	f, err = os.Create(filename)
	if err != nil {
		return
	}

	return
}
