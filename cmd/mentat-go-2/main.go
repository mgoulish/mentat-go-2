package main

import (
    "flag"
    "fmt"
    
    "github.com/mgoulish/mentat-go-2/internal/config"
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

    for _, s := range sites {
      fmt.Printf("Site %s was read.\n", s.Name)
    }
}



