package server

import (
	"time"
)

func (s *Server) run() {
	for s.state == running {
		time.Sleep(s.heartbeatPause)
		if !s.pingMaster() || s.triggerElection {
			s.triggerElection = true
			startElection(s)
			s.triggerElection = false
		}
	}
}

func (s *Server) pingMaster() bool {
	s.emitter.Write(s.id, s.master, "HEARTBEAT")
	if s.master == "" || (s.master != s.id && !s.NeighborServers[s.master].isUp()) {
		return false
	}
	return true
}

func (s *Server) setMaster(masterID string) {
	s.electionLock.Lock()
	defer s.electionLock.Unlock()
	s.master = masterID
	s.emitter.Write(s.id, s.master, "SET MASTER")
}

func (s *Server) isUp() bool {
	return s.state == running
}
