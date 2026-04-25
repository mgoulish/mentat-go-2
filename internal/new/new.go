package new

import "time"

import "github.com/mgoulish/mentat-go-2/internal/types"


func NewRouter(name string) *types.Router {   
    return &types.Router{
		Name:          name,
		TopologyCalcs: make([]*types.NextHops, 5),
	    }
}


func NewSite() *types.Site {
    return &types.Site{
		Router: NewRouter(""),
	    }
}


func NewConnectivityEvent() *types.ConnectivityEvent {
    now := time.Now()
    return &types.ConnectivityEvent{
		Micros:    now.UnixMicro(),
		Neighbors: make([]string, 0, 3), 
    }
}


