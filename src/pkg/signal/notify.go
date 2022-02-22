package signal

import (
	"os"
	"os/signal"
	"syscall"
)

// Notifier function returns functions for waiting and interrupting the signal for graceful termination
// of the application and the ability to clean up all resources
//
// ```go
//    await, stop := signal.Notifier(func() {
//		 stdout.Info("received a system signal to shut down API server, start the shutdown process..")
//    })
//
//    go func() {
//       if err := someService.Run(ctx); err != nil {
//          stop(err) // Gracefully closing the application cleaning up the weight of resources
//       }
//    }()
//
//    ...some code
//
//    if err := await(); err != nil {
//       some code for handle error
//    }
// ```
func Notifier(on ...func()) (wait func() error, stop func(err ...error)) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	err := make(chan error, 1)

	wait = func() error {
		<-sig
		if len(on) > 0 {
			on[0]()
		}
		select {
		case e := <-err:
			return e
		default:
			return nil
		}
	}

	stop = func(e ...error) {
		if len(e) > 0 {
			err <- e[0]
		}
		sig <- syscall.SIGINT
	}

	return wait, stop
}
