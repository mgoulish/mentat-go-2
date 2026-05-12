package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"

	"github.com/mgoulish/mentat-go-2/internal/new"
	//"github.com/mgoulish/mentat-go-2/internal/parse"
	"github.com/mgoulish/mentat-go-2/internal/router"
	"github.com/mgoulish/mentat-go-2/internal/types"
)

var fp = fmt.Printf

type YAML_Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type YAML_Site struct {
	Metadata YAML_Metadata `yaml:"metadata"`
}

func ReadSite(root string) (*types.Site, error) {
	//fmt.Printf("Reading site from root %s\n", root)

	dir := root + "/site-namespace/resources"

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	site := new.NewSite()

	for _, entry := range entries {
		if entry.IsDir() {
			continue // skip directories
		}
		name := entry.Name()
		if strings.HasPrefix(name, "Site") && strings.HasSuffix(name, "yaml") {
			full_path := filepath.Join(dir, name)
			data, err := os.ReadFile(full_path)
			if err != nil {
				panic(fmt.Errorf("failed to read file: %w", err))
			}
			var yaml_site YAML_Site
			if err := yaml.Unmarshal(data, &yaml_site); err != nil {
				panic(fmt.Errorf("failed to parse YAML: %w", err))
			}

			fmt.Println("Name:", yaml_site.Metadata.Name)
			site.Name = yaml_site.Metadata.Name
			site.Path = full_path
			//fmt.Println("Namespace:", yaml_site.Metadata.Namespace)
		}
	}

	router.Read(root+"/site-namespace/logs", site.Router)

	return site, nil
}
