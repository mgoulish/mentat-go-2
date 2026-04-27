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

    // Demonstrate connectivity functionality
    connectivity.Find(sites)
    connectivity.Print(sites)
    connectivity.Check(sites, "2025-09-09 14:00:00")
    connectivity.Check(sites, "2025-09-16 04:20:00")
}



