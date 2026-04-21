
package config

import (
	"fmt"
	"os"

	"github.com/mgoulish/mentat-go-2/internal/new"
)

// ReadSites reads the directory and returns a slice of *new.Site for each entry.
func ReadSites(root string) ([]*new.Site, error) {
    fmt.Printf("Reading sites from root %s\n", root)

    entries, err := os.ReadDir(root)
    if err != nil {
	return nil, err
    }

    sites := make([]*new.Site, 0, len(entries))

    for _, entry := range entries {
	s := new.NewSite()        // returns *Site
	s.Name = entry.Name()     // set the name

	sites = append(sites, s)

	//fmt.Println("Read Site :", s.Name)
    }

    return sites, nil
}
