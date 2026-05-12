package types

import "time"

// NextHops represents one "Computed next hops" entry from the logs
type NextHops struct {
	Timestamp time.Time
	Micros    int64
	MapData   map[string]string
	Message   string
	RawLine   string
	LineNum   int
}

type ConnectivityEvent struct {
	Timestamp string
	Micros    int64
	Neighbors []string
}

// Router is the main container for a router's parsed data
type Router struct {
	Name               string
	TopologyCalcs      []*NextHops
	ConnectivityEvents []*ConnectivityEvent
}

// Site represents a site containing one router
type Site struct {
	Name   string
	Path   string
	Router *Router
}
