package main

import (
	"fmt"
	// "github.com/josip/timewarp"
	"github.com/robig/cron"
	"strconv"
	"sync"
	"time"
)

// const ONE_SECOND = 1*time.Second + 10*time.Millisecond

func main() {

	a := -1

	println(strconv.Itoa(a))

	// cr := time.Now()
	fmt.Println(time.Now().AddDate(0, 0, -1))

	fmt.Println(time.Now())
	c := cron.New()
	c.AddFunc("* * * * * ?", func() { fmt.Println("Every 1 minutes") })
	c.AddFunc("@hourly", func() { fmt.Println(time.Now()) })

	c.Start()

	// select {}
	time.Sleep(1e10)
	// for {

	// select {}
	// println("a")
	// }
	println("End here")

	// wg := &sync.WaitGroup{}
	// wg.Add(1)

	// cron := cron.New()
	// cron.AddFunc("* * * * * ?", func() { println("heeee") })
	// cron.Start()
	// defer cron.Stop()

	// // Give cron 2 seconds to run our job (which is always activated).
	// // select {
	// // case <-time.After(ONE_SECOND):
	// // 	println("aaaaa")
	// // case <-wait(wg):
	// // 	println("hhhhhh")
	// // }

	// select {}

	// for {
	// }

}

func wait(wg *sync.WaitGroup) chan bool {
	ch := make(chan bool)
	go func() {
		wg.Wait()
		ch <- true
	}()
	return ch
}
