package goutils

import (
	"errors"
	"fmt"
	"os"
	"path"
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
	var (
		base = path.Dir(filename)
	)

	_, err = os.Stat(base)
	if err != nil {
		err = os.MkdirAll(base, os.ModePerm)
		if err != nil {
			return
		}
	}

	_, err = os.Stat(filename)
	if err == nil {
		err = errors.New(fmt.Sprint("%s is exist", filename))
		return
	}

	f, err = os.Create(filename)
	if err != nil {
		return
	}

	return
}
