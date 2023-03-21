package internal

import (
	"fmt"
	"strings"

	"github.com/gocraft/web"
)

type HandlerImpl struct {
	HelloCount int
}

type Handler interface {
	SayHello(rw web.ResponseWriter, req *web.Request)
}

func ProvideHandler(counter int) *HandlerImpl {
	return &HandlerImpl{
		HelloCount: counter,
	}
}

func (s *HandlerImpl) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", s.HelloCount), "World!")
}
