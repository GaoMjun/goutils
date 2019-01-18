package goutils

import (
	"log"
	"testing"
	"time"
)

func TestRepeatTimer(t *testing.T) {
	end := make(chan bool)

	RepeatTimer(time.Second*1, func() {
		log.Println("RepeatTimer")
	})

	for range end {
		return
	}
}

func TestCreateFileNotExist(t *testing.T) {
	f, err := CreateFileNotExist("test")
	if err != nil {
		log.Println(err)
		return
	}

	f.WriteString("hjhh")
	f.Close()
}

func TestRandBytes(t *testing.T) {
	bs := RandBytes(8)

	log.Printf("%x", bs)

	s := RandString(8)
	log.Println(s)
}
