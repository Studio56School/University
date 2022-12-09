package server

import (
	"github.com/Studio56School/university/internal/handler"
)

type ServerHandlers struct {
	university handler.IHandler
}

func newHandlers(svc *ServerServices, logger *logger.Logger) (*ServerHandlers, error) {

	h := &ServerHandlers{}
	h.university = handler.NewHandler(svc.Srv, logger)

	return h, nil
}
