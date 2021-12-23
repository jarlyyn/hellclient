package loggers

import (
	"syscall"
)

func init() {
	reopenSignals = append(reopenSignals, syscall.SIGUSR1)
}
