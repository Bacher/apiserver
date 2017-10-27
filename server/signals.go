package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func addSigTermHandler() {
	sigCh := make(chan os.Signal)

	signal.Notify(sigCh, syscall.SIGINT)

	//go func() {
	<-sigCh

	<-time.NewTimer(5 * time.Second).C

	os.Exit(1)
	//}()
}
