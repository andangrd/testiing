package internal

import (
	"fmt"
	"strings"
)

type Service struct {
	HelloCount int
}

type ServiceImpl interface {
	SayHello(counter int) string
	AddCounter(param int) int
}

func ProvideService(counter int) *Service {
	return &Service{
		HelloCount: counter,
	}
}

func (s *Service) SayHello(counter int) string {
	newVal := s.AddCounter(counter)
	return fmt.Sprintf("%s%s", strings.Repeat("Hello ", newVal), "World!")
}

func (s *Service) AddCounter(param int) int {
	return s.HelloCount + param
}
