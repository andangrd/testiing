package internal

import (
	"fmt"
	"strings"
)

type ServiceImpl struct {
	HelloCount int
}

type Service interface {
	SayHello(counter int) string
}

func ProvideService(counter int) *ServiceImpl {
	return &ServiceImpl{
		HelloCount: counter,
	}
}

func (s *ServiceImpl) SayHello(counter int) string {
	// newVal := s.AddCounter(counter)
	return fmt.Sprintf("%s%s", strings.Repeat("Hello ", counter), "World!")
}

func (s *ServiceImpl) AddCounter(param int) int {
	return s.HelloCount + param
}
