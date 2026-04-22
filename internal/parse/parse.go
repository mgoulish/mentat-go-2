package parse

import (
         "fmt"
	 "os"
)


func ParseRouterLog(path string) error {
    fmt.Printf ( "Read router log at %s\n", path )

    entries, err := os.ReadDir(path)
    if err != nil {
        return err
    }
    for _, entry := range entries {
        name := entry.Name()
        fmt.Printf ( "ParseRouterLog: %s\n", name )
    }
    return nil
}


