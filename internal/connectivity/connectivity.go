package connectivity

import (
    "fmt"
    "sort"
    "time"

    "github.com/mgoulish/mentat-go-2/internal/new"
    "github.com/mgoulish/mentat-go-2/internal/types"
)



func Find(sites []*types.Site) {
    for _, site := range sites {
        fmt.Printf("Finding connectivity for site %s\n", site.Name)
	router := site.Router
	if len(router.TopologyCalcs) == 0 {
	    fmt.Printf("Connectivity: Router %s has no topology calculations\n", router.Name)
	    return
	}

	//fmt.Printf("Connectivity: Router %s — %d topology calculations\n", 
	    //router.Name, len(router.TopologyCalcs))

	for _, calc := range router.TopologyCalcs {
	    // Each topology calc becomes one Connectivity Event
	    ce := new.NewConnectivityEvent()
	    //time_str := calc.Timestamp.Format("15:04:05.123")
	    time_str := calc.Timestamp.Format("2006-01-02 15:04:05.123")
	    ce.Timestamp = time_str
	    //fmt.Printf("\n%s == %d\n", time_str, calc.Micros)
	    ce.Timestamp = time_str
	    ce.Micros    = calc.Timestamp.UnixMicro()

	    if len(calc.MapData) == 0 {
		//fmt.Println("    NO NEIGHBORS")
		router.ConnectivityEvents = append(router.ConnectivityEvents, ce)
		continue
	    }

	    // This topology calc may contain many next-hop calcs.
	    for dest, next_hop := range calc.MapData {
		//fmt.Printf("dest: %s, next_hop: %s\n", dest, next_hop)
		if dest == next_hop {
		    //fmt.Printf("    NEIGHBOR: %s\n", next_hop)
		    ce.Neighbors = append(ce.Neighbors, next_hop)
		}
	    }

	    router.ConnectivityEvents = append(router.ConnectivityEvents, ce)
	}
        
	// Make sure the array of Connectivity Events is sorted in 
	// order of ascending miscroseconds.
	// I think it already should be, but let's make sure.
	sort.Slice(router.ConnectivityEvents, func(i, j int) bool {
            return router.ConnectivityEvents[i].Micros < router.ConnectivityEvents[j].Micros
        })
    }
}



func Print(sites []*types.Site) {
    shift_width := "    " // 4 spaces
    fmt.Printf("\nConnection History\n------------------------------\n")
    for _, site := range sites {
        indent := ""
        fmt.Printf("Site: %s\n", site.Name)
        router := site.Router
        r_indent := indent + shift_width
        fmt.Printf("%sRouter: %s\n", r_indent, router.Name)
        for _, ce := range router.ConnectivityEvents {
	    t_indent := r_indent + shift_width
            fmt.Printf("%sat time %s neigbors are:\n", t_indent, ce.Timestamp)
	    n_indent := t_indent + shift_width
	    if len(ce.Neighbors) < 1 {
	            fmt.Printf("%sNO NEIGHBORS\n", n_indent)
	    } else {
	        for _, neighbor := range ce.Neighbors {
	            fmt.Printf("%s%s\n", n_indent, neighbor)
		}
	    }
        }
    }
    fmt.Printf("\n\n")
}



// Example Timestamps:
//     2025-09-09 14:00:00.000
//     2025-09-16 04:20:00.0000
func Check(sites []*types.Site, timestr string) (error) {
    fmt.Printf("\nChecking Connectivity at time %s\n", timestr)
    fmt.Printf("----------------------------------------------------\n")
    t, err := time.Parse("2006-01-02 15:04:05", timestr)
    if err != nil {
	fmt.Printf ("err\n")
        return fmt.Errorf("timestamp parse error: %w", err)
    }
    check_micros := t.UnixMicro()
    for _, site := range sites {
        router := site.Router
	fmt.Printf ( "\n  Checking site: %s   router: %s\n", site.Name, router.Name)
	if router.ConnectivityEvents[0].Micros > check_micros {
	  fmt.Printf("    timestamp is before first connectivity event.\n")
	  continue
	}
	latest_ce := router.ConnectivityEvents[0]
        for _, ce := range router.ConnectivityEvents {
	    //fmt.Printf("    check: %d ce: %d\n", check_micros, ce.Micros)
	    if check_micros < ce.Micros {
	      //fmt.Printf("    that ce was bigger!\n")
	      break
	    }
	    latest_ce = ce
        }
	//fmt.Printf ( "latest_ce was: %v\n", latest_ce )
	//fmt.Printf("n neighbors in previous CE is %d\n", len(latest_ce.Neighbors) )
	if 0 == len(latest_ce.Neighbors) {
	    colorRed   := "\033[31m"
            colorReset := "\033[0m"
	    fmt.Printf("%s    NO CONNECTIVITY\n%s", colorRed, colorReset)
	} else {
	    fmt.Printf("    Neighbors:\n")
	    for _, neighbor := range(latest_ce.Neighbors) {
	        fmt.Printf("        %s\n", neighbor)
	    }
	}
	fmt.Printf("\n\n")
    }
    
    return nil
}



