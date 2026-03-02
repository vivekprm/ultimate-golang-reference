package main

import "log"

// Server defines a contract for tcp servers.
type Server interface {
	Start() error
	Stop() error
	Wait() error
}

// server is our Server implementation.
type server struct {
	host string
}

// NewServer returns an interface value of type Server
// with an xServer implementation.
func NewServer(host string) Server {
	return &server{host}
}

// Start allows the server to begin to accept requests.
func (s *server) Start() error {
	log.Println("Starting server on", s.host)
	return nil
}

// Stop shuts the server down.
func (s *server) Stop() error {
	log.Println("Stopping server on", s.host)
	return nil
}

// Wait prevents the server from accepting new connections.
func (s *server) Wait() error {
	return nil
}

func main() {
	srv := NewServer("localhost:8080")
	if err := srv.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}

	// Wait for a signal to stop the server, e.g., os.Interrupt
	// ...
	if err := srv.Stop(); err != nil {
		log.Fatal("Failed to stop server:", err)
	}
}
