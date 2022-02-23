package main

import (
	"github.com/zieckey/etcdsync"
	"log"
	"time"
)

func main() {
	m, err := etcdsync.New("/lock1", 100, []string{"http://192.168.21.118:12379","http://192.168.21.118:22379","http://192.168.21.118:32379"})
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}
	err = m.Lock()
	if err != nil {
		log.Printf("etcdsync.Lock failed")
		return
	}
	log.Printf("etcdsync.Lock OK")
	time.Sleep(10 * time.Second)
	log.Printf("Get the lock. Do something here.")

	err = m.Unlock()
	if err != nil {
		log.Printf("etcdsync.Unlock failed")
	} else {
		log.Printf("etcdsync.Unlock OK")
	}
}
