package service

import (
	"fmt"
	"github.com/zikwall/blogchain/src/platform/log"
	"runtime"
	"strconv"
	"time"
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

func (s *ServiceInstance) AddNotifiers(notifiers ...Notifier) {
	for _, notifier := range notifiers {
		s.AddNotify(notifier)
	}
}

func (s *ServiceInstance) AddNotify(notify Notifier) {
	s.notifiers = append(s.notifiers, notify)
}

func (s ServiceInstance) Shutdown(onError func(error)) {
	log.Info("Shutdown Blogchain Service via System signal")

	// cancel root context
	s.cancelFunc()

	for _, notifier := range s.notifiers {
		log.Info(notifier.CloseMessage())

		if err := notifier.Close(); err != nil {
			onError(err)
		}
	}
}

func (s ServiceInstance) Stacktrace() {
	log.Info("Waiting for the server completion report to be generated")

	<-time.After(time.Second * 2)

	memory := runtime.MemStats{}
	runtime.ReadMemStats(&memory)

	colored := func(category, context string) string {
		return fmt.Sprintf("%s: %s", log.Colored(category, log.Cyan), log.Colored(context, log.Green))
	}

	fmt.Println(
		fmt.Sprintf("%s \n \t - %s \n \t - %s",
			log.Colored("REPORT", log.Green),
			colored("Number of remaining goroutines:", strconv.Itoa(runtime.NumGoroutine())),
			colored("Number of operations of the garbage collector:", strconv.Itoa(int(memory.NumGC))),
		),
	)
}
