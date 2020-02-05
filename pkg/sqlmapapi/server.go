package sqlmapapi

import (
	"context"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// Server launches the SQLMap API server by exec'ing the python command
// found at cmd and using the options (if any) passed as opts.  If no
// error is returned, a cancelFunc is provided to stop the server when
// the caller is done with it.
func Server(cmd string, opts ...option) (context.CancelFunc, error) {
	log.Info("Starting SQLMap API server")
	args := arguments{"--server"}

	for _, opt := range opts {
		opt(&args)
	}

	srv := exec.Command(cmd, args...)

	cancel := func() {
		log.Info("Stopping SQLMap API server")
		err := srv.Process.Kill()
		if err != nil {
			log.Fatal("Failed to stop the SQLMap API server")
		}
	}

	return cancel, srv.Start()
}
