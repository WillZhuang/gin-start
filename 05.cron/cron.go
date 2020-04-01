package main

import (
	"time"
	"log"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")
	// Create Cron job runner
	c := cron.New()
	// AddFunc 会向 Cron job runner 添加一个 func, 以按给定的时间表运行
	c.AddFunc("* * * * * *", func() {
		log.Println("Run Cron...")
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}