package types

import "time"

// NextHops represents one "Computed next hops" entry from the logs
type NextHops struct {
	Timestamp time.Time
	MapData   map[string]string
	Message   string
	RawLine   string
	LineNum   int
}

// Router is the main container for a router's parsed data
type Router struct {
	Name          string
	TopologyCalcs []*NextHops
}

// Site represents a site containing one router
type Site struct {
	Name   string
	Path   string
	Router *Router
}
