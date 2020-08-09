package base

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type (
	OpenCloser interface {
		Opener
		Closer
	}

	Opener interface {
		Open() error
	}

	Closer interface {
		Close() error
		Message() string
	}
)

func Notifier(openers ...Closer) {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig

		fmt.Println("Stoped Server by signal")

		for _, opener := range openers {
			switch opener.(type) {
			case Closer:
				o := opener.(Closer)
				fmt.Println(o.Message())
			}

			if err := opener.Close(); err != nil {
				// notify
			}
		}

		os.Exit(0)
	}()
}
