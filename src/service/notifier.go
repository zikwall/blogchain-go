package service

import (
	"os"
	"os/signal"
	"syscall"
)

type (
	Notify struct {
		notifiers []Notifier
	}
	Notifier interface {
		Close() error
		CloseMessage() string
	}
)

func (s *BlogchainServiceInstance) AddNotifiers(notifiers ...Notifier) {
	for _, notifier := range notifiers {
		s.AddNotify(notifier)
	}
}

func (s *BlogchainServiceInstance) AddNotify(notify Notifier) {
	s.notifiers = append(s.notifiers, notify)
}

func (s *BlogchainServiceInstance) WaitBlogchainSystemNotify() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig
}

func (s *BlogchainServiceInstance) ShutdownBlogchainServer() {
	for _, notifier := range s.notifiers {
		if err := notifier.Close(); err != nil {
			s.logger.Warning(err)
		}
	}
}
