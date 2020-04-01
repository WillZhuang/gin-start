package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	TrapSignal(func() {
		fmt.Println("before exit.")
	})

	engine.Run()
}

// TrapSignal catches the SIGTERM/SIGINT/SIGKILL and executes cb function.
// After that it exits with code 0.
func TrapSignal(cb func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		for sig := range c {
			fmt.Println("captured %v, exiting...", sig)
			if cb != nil {
				cb()
			}
			fmt.Println("exit")
			os.Exit(0)
		}
	}()
}