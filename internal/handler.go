package internal

import (
	"fmt"
	"strconv"

	"github.com/gocraft/web"
)

type HandlerImpl struct {
	Service Service
}

type Handler interface {
	SayHello(rw web.ResponseWriter, req *web.Request)
}

func ProvideHandler(counter int) *HandlerImpl {

	svc := ProvideService(counter)
	return &HandlerImpl{
		Service: *svc,
	}
}

func (s *HandlerImpl) SayHello(rw web.ResponseWriter, req *web.Request) {
	counterStr := req.PathParams["counter"]
	counter, _ := strconv.Atoi(counterStr)
	if counter == 0 {
		counter = 3
	}
	result := s.Service.SayHello(counter)

	fmt.Fprint(rw, result)
}
