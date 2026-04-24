package connectivity

import (
	"fmt"
	"sort"

	"github.com/mgoulish/mentat-go-2/internal/new"
	"github.com/mgoulish/mentat-go-2/internal/parse"
)



func Connectivity(router *new.Router) {
    tc, ok := router.Data["topology calcs"]
    if !ok {
	fmt.Printf("Connectivity: Router %s has no topology calcs\n", router.Name)
	return
    }

    // It's an array of empty interfaces,
    // so we need to cast it to a specific type.
    topology_calcs, ok := tc.([]*parse.NextHops)
    if !ok {
	fmt.Printf("Connectivity: Router %s - wrong type for topology calcs\n", router.Name)
	return
    }

    fmt.Printf("Connectivity: Router %s — %d topology calculations\n", 
               router.Name, len(topology_calcs))

    for _, calc := range topology_calcs {
	fmt.Printf("\n%s  %s\n",
	    calc.Timestamp.Format("15:04:05"),
	    calc.Message)

	if len(calc.MapData) == 0 {
	    fmt.Println("    (empty map)")
	    continue
	}

	// Print keys in sorted order for consistent output
	keys := make([]string, 0, len(calc.MapData))
	for k := range calc.MapData {
	    keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
	    fmt.Printf("    %-50s → %s\n", key, calc.MapData[key])
	}
    }
}



