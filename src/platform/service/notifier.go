package service

type (
	Notify struct {
		notifiers []Notifier
	}
	Notifier interface {
		Close() error
		CloseMessage() string
	}
)

func (s *Notify) AddNotifiers(notifiers ...Notifier) {
	for _, notifier := range notifiers {
		s.AddNotify(notifier)
	}
}

func (s *Notify) AddNotify(notify Notifier) {
	s.notifiers = append(s.notifiers, notify)
}
