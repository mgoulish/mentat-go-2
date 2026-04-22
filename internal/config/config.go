
package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/mgoulish/mentat-go-2/internal/new"
	"github.com/mgoulish/mentat-go-2/internal/parse"
)


func ReadSites(root string) ([]*new.Site, error) {
    fmt.Printf("Reading sites from root %s\n", root)

    entries, err := os.ReadDir(root)
    if err != nil {
	return nil, err
    }

    sites := make([]*new.Site, 0, len(entries))

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



func ReadRouter(path string, router *new.Router) error {

    entries, err := os.ReadDir(path)
    if err != nil {
	return err
    }
    for _, entry := range entries {
	name := entry.Name()
	if strings.HasPrefix(name, "skupper-router") {
	  router.Name = name
	  log_dir_path := path + "/" + name + "/" + "logs"
	  parse.ParseRouterLog(log_dir_path)
	}
    }
    return nil
}


