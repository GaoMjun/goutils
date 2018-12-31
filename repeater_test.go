package goutils

import (
	"log"
	"testing"
	"time"
)

func TestRepeater(t *testing.T) {
	r := NewRepeater(time.Second*1, func() {
		log.Println("tick")
	})

	time.Sleep(time.Second * 10)
	r.Stop()
}
