package new

import "github.com/mgoulish/mentat-go-2/internal/types"

func NewRouter(name string) *types.Router {   // ← added name parameter
	return &types.Router{
		Name:          name,
		TopologyCalcs: make([]*types.NextHops, 0),
	}
}

func NewSite() *types.Site {
	return &types.Site{
		Router: NewRouter(""), // name will be set later
	}
}



