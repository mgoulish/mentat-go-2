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

func DoEverything(path_arg string) { 
        var site *types.Site
        var err error

	// If the argument path is a directory,
	// we will just use it as-is.
	// But if it's a tar.gz file, we will expand it
	// into a new directory, and then use that directory.
	path_to_use := path_arg


        if utils.IsTarGz(path_arg) {
                fp("That's a tar file!\n")
                new_path, e := utils.ExpandTGZFile(path_arg)
                if e != nil {
                        fp("error: %s\n", e)
                        os.Exit(1)
                }
                fp("new_path is %s\n", new_path)
                path_to_use = new_path
        }
        if utils.IsDir(path_to_use) {
                fp("%s is a dir!\n", path_to_use)
                site, err = config.ReadSite(path_to_use)
                fp("Main gets site: %+v\n", site)
                if err != nil {
                        fmt.Printf("Error reading site in dir %s: %v\n", path_to_use, err)
                        return
                }
        } else {
                fp("%s is not a dir.\n", path_to_use)
                os.Exit(1)
        }

        fmt.Printf("pkg: site: %s\n", site.Name)
        connectivity.Find(site)
        connectivity.Print(site)
        connectivity.Check(site, "2026-04-29 18:06:00")
        connectivity.Check(site, "2026-04-29 19:11:00")
 
}




