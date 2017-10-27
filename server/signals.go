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

	go rpcClient.Request("disconnect", nil)

	<-time.NewTimer(5 * time.Second).C

	os.Exit(1)
	//}()
}
