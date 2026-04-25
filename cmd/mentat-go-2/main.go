package main

import (
    "flag"
    "fmt"
    
    "github.com/mgoulish/mentat-go-2/internal/config"
    "github.com/mgoulish/mentat-go-2/internal/connectivity"
)

func main() {
    root := flag.String("root", ".", "root dir of data, should be date like 2025_05_30")

    flag.Parse()

    fmt.Printf("Data root dir is %s\n", *root)

    sites, err := config.ReadSites ( *root )
    if err != nil {
      fmt.Printf("Error reading sites: %v\n", err)
      return
    }

    for _, site := range sites {
        fmt.Printf("Site %s was read. Router: %s\n", site.Name, site.Router.Name)
	connectivity.Connectivity(site.Router)

	for _, ce := range site.Router.ConnectivityEvents {
	  fmt.Printf ( "    CE: %d\n", ce.Micros )
	  fmt.Printf ( "        %v\n", ce.Neighbors )
	}
    }
}



