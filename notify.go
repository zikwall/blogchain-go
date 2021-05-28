package main

import (
	"github.com/zikwall/blogchain/src/platform/log"
	"os"
	"os/signal"
	"syscall"
)

func congratulations() {
	log.Info("Congratulations, the Blogchain server has been successfully launched")
}

func wait(onReceiveSignal func()) {
	congratulations()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig

	onReceiveSignal()
}
