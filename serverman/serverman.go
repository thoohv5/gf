package serverman

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type (
	Server interface {
		Serve()
		Stop()
	}
	ServerMan struct {
		svrs []Server
	}
)

var (
	defaultSeverMan = NewServerMan()
)

func NewServerMan() *ServerMan {
	return &ServerMan{
		svrs: make([]Server, 0, 1),
	}
}

func (s *ServerMan) RegisterServer(server Server) {
	s.svrs = append(s.svrs, server)
}

func (s *ServerMan) Start() {
	go handleSysSignal()
	wg := sync.WaitGroup{}
	wg.Add(len(s.svrs))
	for _, svr := range s.svrs {
		go func(s Server) {
			defer wg.Done()
			s.Serve()
		}(svr)
	}
	wg.Wait()
}

func (s *ServerMan) Stop() {
	for _, svr := range s.svrs {
		svr.Stop()
	}
}

func RegisterServer(server Server) {
	defaultSeverMan.RegisterServer(server)
}

func Start() {
	defaultSeverMan.Start()
}

func Stop() {
	defaultSeverMan.Stop()
}

func handleSysSignal() {
	sChan := make(chan os.Signal)
	for {
		signal.Notify(sChan, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		sig := <-sChan
		switch sig {
		case os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			Stop()
		}

	}

}
