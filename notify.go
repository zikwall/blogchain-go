package main

import (
	"github.com/zikwall/blogchain/src/platform/log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ListenerTCP = iota + 1
	ListenerUDS
)

type signalReceiver struct {
	onSignal func()
}

func congratulations() {
	log.Info("Congratulations, the Blogchain server has been successfully launched")
}

func awaiter(r signalReceiver) (func() error, func(err ...error)) {
	congratulations()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	wait := func() error {
		// wait signal for the close application
		<-sig

		if r.onSignal != nil {
			r.onSignal()
		}

		// receive err from stop function
		return nil
	}

	// add send error to await function
	stop := func(err ...error) {
		// Send a signal to end the application
		sig <- syscall.SIGINT

		if len(err) > 0 {
			log.Warning(err[0])
		}
	}

	return wait, stop
}

func listenToUnix(bind string) (net.Listener, error) {
	_, err := os.Stat(bind)

	if err == nil {
		err = os.Remove(bind)

		if err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	return net.Listen("unix", bind)
}

func chmodSocket(sock string) {
	// problem with unix socket permissions
	go func() {
		// wait complete create socket
		<-time.After(time.Second * 2)

		if err := os.Chmod(sock, 0666); err != nil {
			log.Warning(err)
		}
	}()
}
