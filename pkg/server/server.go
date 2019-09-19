package server

import (
	"fmt"
	"time"
)

// Server is a main entry point in Sygen Code
type Server struct {
	runTime     int
	requestRate int
}

// New is a method creates a new Server
func New() *Server {
	return &Server{}
}

// Start the server.
func (s *Server) Start() error {
	fmt.Println("Sygen started")
	tick(5,100)

	return nil
}

// Stop cleans up resources used by the server.
func (s *Server) Stop() {
	fmt.Println("Sygen stopped")
}

func tick(timeInSecs int, rate int) {

	ticker := time.NewTicker(time.Duration(1000 / rate) * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				go func() {
					fmt.Println("Run at", t)
				}()
			}
		}
	}()
	time.Sleep(time.Duration(timeInSecs) * time.Second)
	ticker.Stop()
	done <- true
}
