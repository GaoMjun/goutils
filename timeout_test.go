package goutils

import (
	"log"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	end := make(chan bool)

	timeout := NewTimeout(time.Second*5, func() {
		log.Println("timeout")
		log.Println(time.Now())
		end <- true
	})

	log.Println(time.Now())

	time.Sleep(time.Second * 2)
	log.Println(time.Now())
	timeout.ChangeTime(time.Second * 10)

	time.Sleep(time.Second * 2)
	log.Println(time.Now())

	go func() {
		timeout.Stop()
		end <- true
	}()

	for range end {
		return
	}
}
