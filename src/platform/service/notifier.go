package service

import (
	"github.com/zikwall/blogchain/src/platform/log"
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

func (s *BlogchainServiceInstance) ShutdownBlogchainServer(onError func(error)) {
	log.Info("Shutdown Blogchain Service via System signal")

	for _, notifier := range s.notifiers {
		log.Info(notifier.CloseMessage())

		if err := notifier.Close(); err != nil {
			onError(err)
		}
	}
}
