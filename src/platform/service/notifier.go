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

func (n *Notify) AddNotifiers(notifiers ...Notifier) {
	for _, notifier := range notifiers {
		n.AddNotify(notifier)
	}
}

func (n *Notify) AddNotify(notify Notifier) {
	n.notifiers = append(n.notifiers, notify)
}
