package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mgoulish/mentat-go-2/internal/utils"
	//"github.com/mgoulish/mentat-go-2/internal/connectivity"
)

var fp = fmt.Printf

func main() {
	path_arg := flag.String("root", ".", "either a directory or a tar.gz file")
	flag.Parse()
	data_path := *path_arg

	if utils.IsTarGz(*path_arg) {
	  fp("That's a tar file!\n")
	  new_path, err := utils.ExpandTGZFile(*path_arg)
	  if err != nil {
	    fp("error: %s\n", err)
	    os.Exit(1)
	  }
	  fp("new_path is %s\n", new_path)
	  data_path = new_path
	}

	if utils.IsDir(data_path) {
	  fp("%s is a dir!\n", data_path )
	} else {
	  fp("%s is not a dir.\n", data_path)
	  os.Exit(1)
	}

	os.Exit(0)


	/*
	sites, err := config.ReadSites(*root)
	if err != nil {
		fmt.Printf("Error reading sites: %v\n", err)
		return
	}

	for _, site := range sites {
		fmt.Printf("main: site: %s\n", site.Name)
		connectivity.Find(site)
		connectivity.Print(site)
		// Demonstrate connectivity functionality
		connectivity.Check(site, "2025-09-09 14:00:00")
		connectivity.Check(site, "2025-09-16 04:20:00")
	}
	*/

}


