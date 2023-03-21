package internal

import (
	"fmt"
	"strings"

	"github.com/gocraft/web"
)

type ServiceImpl struct {
	HelloCount int
}

type Service interface {
	SayHello(rw web.ResponseWriter, req *web.Request)
}

func ProvideService(counter int) *ServiceImpl {
	return &ServiceImpl{
		HelloCount: counter,
	}
}

func (s *ServiceImpl) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", s.HelloCount), "World!")
}
