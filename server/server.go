package server

import (
	"crypto/rand"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/LostLaser/recoverE/emitter"
)

// Server is a single entity
type Server struct {
	master            string
	id                string
	state             string
	NeighborServers   map[string]*Server
	electionAlgorithm Election
	electionLock      sync.Mutex
	triggerElection   bool
	emitter           *emitter.Emitter
	heartbeatPause    time.Duration
}

const (
	running = "running"
	stopped = "stopped"
)

// New will create a cluster with the specified number of servers
func New(e *emitter.Emitter, heartbeatPause time.Duration) *Server {
	s := new(Server)
	s.id = generateUniqueID()
	s.state = stopped
	s.NeighborServers = make(map[string]*Server)
	s.emitter = e
	s.heartbeatPause = heartbeatPause

	return s
}

// Start brings up the server
func (s *Server) Start() {
	s.state = running
	go s.run()
}

// Stop the provided server
func (s *Server) Stop() {
	s.state = stopped
}

// Print displays the server information in a readable format
func (s *Server) Print() {
	fmt.Println("ID:", s.id, " Master:", s.master)
}

// GetID returns the server id
func (s *Server) GetID() string {
	return s.id
}

func generateUniqueID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}