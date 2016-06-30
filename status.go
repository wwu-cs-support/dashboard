package main

///////////////
// Status.go //
///////////////

import (
	"net"
	"sync"
	"time"
)

const (
	LINUX        = 0
	WINDOWS      = 1
	INACCESSIBLE = 2
)

// systemStatus takes a hostname(str) and returns what operating system that
// machine is running based on which port is successfully used to connect.
func systemStatus(hostname string) int {
	if accessible(hostname, "***REMOVED***") {
		return LINUX
	} else if accessible(hostname, "***REMOVED***") {
		return WINDOWS
	} else {
		return INACCESSIBLE
	}
}

// accessible takes a hostname and a port number and tries to establish a
// connection using those parameters.
// It returns an error, if one occurrs.
func accessible(hostn string, port string) bool {
	conn, err := net.DialTimeout("tcp", hostn+":"+port, time.Millisecond*50)
	check := check(err)
	if check {
		conn.Close()
		return true
	}
	return false
}

//
//
func updateStatuses(machines []*Machine) {
	var wg sync.WaitGroup
	for _, machine := range machines {
		wg.Add(1)

		go func(m *Machine) {
			defer wg.Done()
			// m.UpdateStatus()
		}(machine)
	}
	wg.Wait()
}