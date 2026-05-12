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

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	// example: skupper-router-6b5cf5c8f4-wtjxd-router.txt
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, "skupper-router") && strings.HasSuffix(name, "router.txt") {
			log_path := path + "/" + name
			fp("log_path %s\n", log_path)
			router.Name = name
			parse.ParseRouterLogs(log_path, router)
		}
	}
	return nil
}
