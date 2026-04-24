package connectivity

import (
	"fmt"
	"sort"

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
		fmt.Printf("\n%s  %s\n", 
			calc.Timestamp.Format("15:04:05"),
			calc.Message)

		if len(calc.MapData) == 0 {
			fmt.Println("    (empty map)")
			continue
		}

		// Sorted keys for nice output
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
