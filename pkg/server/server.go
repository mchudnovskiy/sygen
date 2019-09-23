package server

import (
	"fmt"
	"github.com/mchudnovskiy/sygen/pkg/server/sender"
	"github.com/mchudnovskiy/sygen/pkg/server/settings"
	"time"
)

// Server is a main entry point in Sygen Code
type Server struct {
	args *settings.Args
}

// New is a method creates a new Server
func New(a *settings.Args) *Server {
	return &Server{
		args: a,
	}
}

// Start the server.
func (s *Server) Start() error {
	fmt.Println("Sygen started")
	tick(s.args.ExecutionTime, s.args.RequestRate)

	return nil
}

// Stop cleans up resources used by the server.
func (s *Server) Stop() {
	fmt.Println("Sygen stopped")
}

func tick(timeInSecs int, rate int) {
	ticker := time.NewTicker(time.Duration(1000/rate) * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				go func() {
					fmt.Printf("run query at: %v\n", t)
					s, _ := sender.NewSender("HTTP||sdfsdfsdfsdfsdfdfsdf")
					s.Send("payload", map[string]string{
						"header": "value",
					})
				}()
			}
		}
	}()
	time.Sleep(time.Duration(timeInSecs) * time.Second)
	ticker.Stop()
	done <- true
}
