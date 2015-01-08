package main

import (
	"fmt"
	"sync"
	"time"
)

type Chat struct {
	sync.RWMutex
	ActiveConvs map[string]string
}

func (this *Chat) Read() {
	for i := 0; i < 10; i++ {

		go func() {
			this.RLock()
			defer this.Unlock()
			for key, value := range this.ActiveConvs {

				fmt.Printf("Key: %v  Value: %v \n", key, value)

			}

		}()

	}

}

func (this *Chat) Write(key, value string) {
	this.Lock()
	defer this.Unlock()
	this.ActiveConvs[key] = value
}

func main() {
	chat := &Chat{
		ActiveConvs: map[string]string{
			"1": "work",
			"2": "study",
		},
	}

	go chat.Write("3", "play")
	chat.Read()
	time.Sleep(1 * time.Second)

}
