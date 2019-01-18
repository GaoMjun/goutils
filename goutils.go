package goutils

import (
	"errors"
	"fmt"
	"math/rand"
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
		err = errors.New(fmt.Sprint("file exist ", filename))
		return
	}

	f, err = os.Create(filename)
	if err != nil {
		return
	}

	return
}

func RandBytes(length int) []byte {
	rand.Seed(time.Now().UnixNano())
	p := make([]byte, length)
	rand.Read(p)
	return p
}

func RandString(length int) string {
	return fmt.Sprintf("%x", RandBytes(length))
}
