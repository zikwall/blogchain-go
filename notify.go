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
	err := make(chan error, 1)

	wait := func() error {
		// wait signal for the close application
		<-sig

		if r.onSignal != nil {
			r.onSignal()
		}

		// receive err from stop function
		select {
		case e := <-err:
			return e
		default:
			return nil
		}
	}

	// add send error to await function
	stop := func(e ...error) {
		if len(e) > 0 {
			err <- e[0]
		}

		// Send a signal to end the application
		sig <- syscall.SIGINT
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
