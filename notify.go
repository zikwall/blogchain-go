package main

import (
	"github.com/zikwall/blogchain/src/platform/log"
	"os"
	"os/signal"
	"syscall"
)

type receiver struct {
	onSignal func()
}

func congratulations() {
	log.Info("Congratulations, the Blogchain server has been successfully launched")
}

func wait(r receiver) {
	congratulations()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig

	if r.onSignal != nil {
		r.onSignal()
	}
}
