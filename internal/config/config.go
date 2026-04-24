package config

import (
	//"fmt"
	"os"
	"strings"

	"github.com/mgoulish/mentat-go-2/internal/new"
	"github.com/mgoulish/mentat-go-2/internal/types"
	"github.com/mgoulish/mentat-go-2/internal/parse"
)


func ReadSites(root string) ([]*types.Site, error) {
    //fmt.Printf("Reading sites from root %s\n", root)

    entries, err := os.ReadDir(root)
    if err != nil {
	return nil, err
    }

    sites := make([]*types.Site, 0, len(entries))

    for _, entry := range entries {
	site := new.NewSite()        
	site.Name = entry.Name()    
	site.Path = root + "/" + site.Name

	routerPath := site.Path + "/pods"
	ReadRouter(routerPath, site.Router)

	sites = append(sites, site)
    }

    return sites, nil
}



func ReadRouter(path string, router *types.Router) error {

    entries, err := os.ReadDir(path)
    if err != nil {
	return err
    }
    for _, entry := range entries {
	name := entry.Name()
	if strings.HasPrefix(name, "skupper-router") {
	  router.Name = name
	  //fmt.Printf ( "ReadRouter:  reading router %s\n", router.Name )
	  log_dir_path := path + "/" + name + "/" + "logs"
	  parse.ParseRouterLogs(log_dir_path, router)
	}
    }
    return nil
}


