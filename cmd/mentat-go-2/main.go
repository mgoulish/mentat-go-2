package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mgoulish/mentat-go-2/internal/config"
	"github.com/mgoulish/mentat-go-2/internal/connectivity"
	"github.com/mgoulish/mentat-go-2/internal/types"
	"github.com/mgoulish/mentat-go-2/internal/utils"
)

var fp = fmt.Printf

func main() {
	path_arg := flag.String("root", ".", "either a directory or a tar.gz file")
	flag.Parse()
	data_path := *path_arg

	var site *types.Site
	var err error

	if utils.IsTarGz(*path_arg) {
		new_path, e := utils.ExpandTGZFile(*path_arg)
		if e != nil {
			fp("error: %s\n", e)
			os.Exit(1)
		}
		data_path = new_path
	}

	if utils.IsDir(data_path) {
		site, err = config.ReadSite(data_path)
			if err != nil {
			fmt.Printf("Error reading site in dir %s: %v\n", data_path, err)
			return
		}
	} else {
		fp("%s is not a dir.\n", data_path)
		os.Exit(1)
	}

	fmt.Printf("main: site: %s\n", site.Name)
	connectivity.Find(site)
	connectivity.Print(site)
	connectivity.Check(site, "2026-04-29 18:06:00")
	connectivity.Check(site, "2026-04-29 19:11:00")
}
