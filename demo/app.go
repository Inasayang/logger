//go:build linux

package main

import (
	"github.com/Inasayang/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	reloadCh := make(chan struct{}, 1)
	//logger.Init(".", "test", "debug", nil) //Copy and Truncate
	logger.Init(".", "test", "debug", reloadCh) //Rename and Create
	go func() {
		for {
			logger.Infof("test msg : %s", time.Now().String())
			time.Sleep(time.Hour)
		}
	}()
	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, syscall.SIGUSR1)
	go func() {
		for _ = range sCh {
			reloadCh <- struct{}{}
		}
	}()
	select {}
}
