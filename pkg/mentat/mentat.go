package mentat

import (
	"fmt"
	"os"

	"github.com/mgoulish/mentat-go-2/internal/config"
	"github.com/mgoulish/mentat-go-2/internal/connectivity"
	"github.com/mgoulish/mentat-go-2/internal/types"
	"github.com/mgoulish/mentat-go-2/internal/utils"
)

var fp = fmt.Printf

func DoEverything(path_arg string) error {
	var site *types.Site
	var err error

	// If the argument path is a directory,
	// we will just use it as-is.
	// But if it's a tar.gz file, we will expand it
	// into a new directory, and then use that directory.
	path_to_use := path_arg

	if utils.IsTarGz(path_arg) {
		new_path, e := utils.ExpandTGZFile(path_arg)
		if e != nil {
			fp("error: %s\n", e)
			os.Exit(1)
		}
		path_to_use = new_path
	}
	if utils.IsDir(path_to_use) {
		site, err = config.ReadSite(path_to_use)
		if err != nil {
			fmt.Printf("Error reading site in dir %s: %v\n", path_to_use, err)
			return nil
		}
	} else {
		fp("%s is not a dir.\n", path_to_use)
		os.Exit(1)
	}

	connectivity.Find(site)
	connectivity.Print(site)
	return nil
}

func CheckAtTime(path_arg string, timestampStr string) error {
	var site *types.Site
	var err error

	// If the argument path is a directory,
	// we will just use it as-is.
	// But if it's a tar.gz file, we will expand it
	// into a new directory, and then use that directory.
	path_to_use := path_arg

	if utils.IsTarGz(path_arg) {
		new_path, e := utils.ExpandTGZFile(path_arg)
		if e != nil {
			fp("error: %s\n", e)
			os.Exit(1)
		}
		path_to_use = new_path
	}
	if utils.IsDir(path_to_use) {
		site, err = config.ReadSite(path_to_use)
		if err != nil {
			fmt.Printf("Error reading site in dir %s: %v\n", path_to_use, err)
			return nil
		}
	} else {
		fp("%s is not a dir.\n", path_to_use)
		os.Exit(1)
	}

	connectivity.Find(site)
	connectivity.Print(site)
	//example: connectivity.Check(site, "2026-04-29 19:11:00")
	connectivity.Check(site, timestampStr)
	return nil
}
