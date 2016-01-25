// Package discard implements a no-op handler useful for benchmarks and tests.
package discard

import (
	"github.com/narqo/go-dice/Godeps/_workspace/src/github.com/apex/log"
)

// Handler implementation.
type Handler struct{}

// New handler.
func New() *Handler {
	return &Handler{}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	return nil
}
