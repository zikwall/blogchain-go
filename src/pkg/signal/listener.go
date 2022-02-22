package signal

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/zikwall/blogchain/src/pkg/log"

	"github.com/gofiber/fiber/v2"
)

const (
	ListenerTCP = iota + 1
	ListenerUDS
)

func ResolveListener(ctx context.Context, listener int, uds, tcp string) (net.Listener, error) {
	if listener == ListenerUDS {
		defer maybeChmodSocket(ctx, uds)
		ln, err := listenToUnix(uds)

		return ln, err
	}
	if !strings.Contains(tcp, ":") {
		tcp = ":" + tcp
	}

	ln, err := net.Listen(fiber.NetworkTCP4, tcp)
	if err != nil {
		return nil, err
	}

	return ln, nil
}

func maybeChmodSocket(c context.Context, sock string) {
	// on Linux and similar systems, there may be problems with the rights to the UDS socket
	go func() {
		ctx, cancel := context.WithTimeout(c, 500*time.Millisecond)
		defer cancel()

		var tryCount uint

		log.Info("run chmod")
		defer log.Info("stop chmod")

		for {
			select {
			case <-ctx.Done():
				log.Info("context is canceled")
				return
			case <-time.After(time.Millisecond * 100):
				log.Info(fmt.Sprintf("loop %d for chmod unix socket (%s)", tryCount, sock))

				if err := os.Chmod(sock, 0o666); err != nil {
					log.Warning(err)
					continue
				}

				_, err := os.Stat(sock)
				// if the file exists and it already has permissions
				if err == nil {
					log.Info(fmt.Sprintf("unix socket (%s) is ready for listen", sock))
					return
				}

				tryCount++
				if tryCount > 5 {
					log.Warning("try count is outside for chmod unix socket")
					return
				}
			}
		}
	}()

	_ = os.Chmod(sock, 0o666)
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
