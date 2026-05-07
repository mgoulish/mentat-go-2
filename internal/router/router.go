package router

import (
  "fmt"
  "os"
  "strings"

  "github.com/mgoulish/mentat-go-2/internal/parse"
  "github.com/mgoulish/mentat-go-2/internal/types"
)

var fp = fmt.Printf


func Read(path string, router *types.Router) error {

        fp("router.Read: my path is %s\n", path)

        entries, err := os.ReadDir(path)
        if err != nil {
                return err
        }
	// skupper-router-6b5cf5c8f4-wtjxd-router.txt
        for _, entry := range entries {
                name := entry.Name()
		if strings.HasPrefix(name, "skupper-router") && strings.HasSuffix(name, "router.txt") {
		  fp("I found my file! %s\n", name)
		  log_path := path + "/" + name
		  fp("log_path %s\n", log_path)
                  router.Name = name
                 parse.ParseRouterLogs(log_path, router)
		}
		/*
                if strings.HasPrefix(name, "skupper-router") {
                        router.Name = name
                        //fmt.Printf ( "ReadRouter:  reading router %s\n", router.Name )
                        log_dir_path := path + "/" + name + "/" + "logs"
                        parse.ParseRouterLogs(log_dir_path, router)
                }
		*/
        }
        return nil
}

