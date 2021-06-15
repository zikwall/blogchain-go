package service

// The Notifier interface and the Notify structure implement a custom solution for closing
// and destroying (clearing) resources after the application is terminated (including an emergency,
// with the exception of a power outage ^_^)
type Notifier interface {
	Close() error
	CloseMessage() string
}

type Notify struct {
	notifiers []Notifier
}

func (n *Notify) AddNotifiers(notifiers ...Notifier) {
	for _, notifier := range notifiers {
		n.AddNotify(notifier)
	}
}

func (n *Notify) AddNotify(notify Notifier) {
	n.notifiers = append(n.notifiers, notify)
}
