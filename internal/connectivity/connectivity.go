package connectivity


import (
    "fmt"
    "github.com/mgoulish/mentat-go-2/internal/new"
)


func Connectivity(router *new.Router) {
    _, ok := router.Data["topology calcs"]

    if ok {
        fmt.Printf ( "Connectivity: Router %s has topo\n", router.Name )
	//     var nextHops_list []*NextHops
	// router.Data["topology calcs"] = topology_calcs

	topology_calcs := router.Data["topology calcs"]
	fmt.Printf ( "Connectivity: there are %d topo calcs.\n", len(topology_calcs) )

	for _, topology_calc := range topology_calcs {
	  fmt.Printf ( "calc: %v\n", topology_calc )
	}

    }
}



