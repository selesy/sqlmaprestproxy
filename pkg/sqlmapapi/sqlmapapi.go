package sqlmapapi

import (
	"context"
	"os/exec"
)

// Server launches the SQLMap API server by exec'ing the python command
// found at cmd and using the options (if any) passed as opts.  If no
// error is returned, a cancelFunc is provided to stop the server when
// the caller is done with it.
func Server(cmd string, opts ...option) (context.CancelFunc, error) {
	args := arguments{"--server"}

	for _, opt := range opts {
		opt(&args)
	}

	ctx, cancel := context.WithCancel(context.Background())
	srv := exec.CommandContext(ctx, cmd, args...)

	return cancel, srv.Start()
}
