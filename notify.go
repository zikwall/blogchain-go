package main

import (
	"context"
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

type conf struct {
	onSignal func()
}

func congratulations() {
	log.Info("congratulations, the Blogchain server has been successfully launched")
}

func notifier(r conf) (wait func() error, stop func(err ...error)) {
	congratulations()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	err := make(chan error, 1)

	wait = func() error {
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
	stop = func(e ...error) {
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

func maybeChmodSocket(c context.Context, sock string) {
	ctx, cancel := context.WithTimeout(c, 500*time.Millisecond)
	defer cancel()

	// on Linux and similar systems, there may be problems with the rights to the UDS socket
	go func() {
		var tryCount uint

		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Millisecond * 100):
				if err := os.Chmod(sock, 0666); err == nil {
					log.Warning(err)
				} else {
					_, err := os.Stat(sock)
					// if the file exists and it already has permissions
					if err == nil {
						return
					}
				}

				tryCount++
				if tryCount > 5 {
					return
				}
			}
		}
	}()

	_ = os.Chmod(sock, 0666)
}
