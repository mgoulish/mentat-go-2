package connectivity

import (
    "fmt"

    "github.com/mgoulish/mentat-go-2/internal/new"
    "github.com/mgoulish/mentat-go-2/internal/types"
)



func Connectivity(router *types.Router) {
    if len(router.TopologyCalcs) == 0 {
	fmt.Printf("Connectivity: Router %s has no topology calculations\n", router.Name)
	return
    }

    fmt.Printf("Connectivity: Router %s — %d topology calculations\n", 
	router.Name, len(router.TopologyCalcs))

    for _, calc := range router.TopologyCalcs {
	// Each topology calc becomes one Connectivity Event
	ce := new.NewConnectivityEvent()
        time_str := calc.Timestamp.Format("15:04:05.123")
	ce.Timestamp = time_str
	fmt.Printf("\n%s == %d\n", time_str, calc.Micros)
	ce.Timestamp = time_str
	ce.Micros    = calc.Timestamp.UnixMicro()

	if len(calc.MapData) == 0 {
	    fmt.Println("    NO NEIGHBORS")
	    router.ConnectivityEvents = append(router.ConnectivityEvents, ce)
	    continue
	}

	// This topology calc may contain many next-hop calcs.
	for dest, next_hop := range calc.MapData {
	    fmt.Printf("dest: %s, next_hop: %s\n", dest, next_hop)
	    if dest == next_hop {
	        fmt.Printf("    NEIGHBOR: %s\n", next_hop)
		ce.Neighbors = append(ce.Neighbors, next_hop)
	    }
	}

	router.ConnectivityEvents = append(router.ConnectivityEvents, ce)
    }
}



