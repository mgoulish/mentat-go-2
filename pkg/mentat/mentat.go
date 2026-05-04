package mentat

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mgoulish/mentat-go-2/internal/config"
	"github.com/mgoulish/mentat-go-2/internal/connectivity"
	"github.com/mgoulish/mentat-go-2/internal/types"
)


// Analyze takes the root directory (the date folder like "2025_04_30")
// and returns the parsed sites after connectivity analysis.
func Analyze(rootDir string) ([]*types.Site, error) {   // adjust type if needed
	sites, err := config.ReadSites(rootDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read sites: %w", err)
	}

	connectivity.Find(sites)
	return sites, nil
}

// AnalyzeAndPrint does the full workflow and prints (like your main does)
func AnalyzeAndPrint(rootDir string) error {
	sites, err := Analyze(rootDir)
	if err != nil {
		return err
	}

	connectivity.Print(sites)

	// Optional: run the example checks
	connectivity.Check(sites, "2025-09-09 14:00:00")  // you can make this configurable later
	return nil
}

/*
// AnalyzeDump is the one you'll call from Skupper (handles tar.gz)
func AnalyzeDump(dumpFile string) error {
	// TODO: extract tar.gz to temp dir (I gave you the helper earlier)
	// For now, assume user gave the already-extracted root dir:
	// extractTarGz(dumpFile, tmpRoot)
	// return AnalyzeAndPrint(tmpRoot)

	fmt.Printf("Would analyze dump: %s\n", dumpFile)
	return nil // placeholder until we add extraction
}
*/


// AnalyzeDump takes a path to a skupper debug dump .tar.gz file,
// extracts it to a temp directory, and runs the full analysis.
func AnalyzeDump(dumpFile string) error {
	// Create a temp directory for extraction
	tmpDir, err := os.MkdirTemp("", "skupper-mentat-*")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir) // clean up when done (or remove this line if you want to keep it)

	fmt.Printf("📦 Extracting dump %s to %s...\n", dumpFile, tmpDir)

	if err := extractTarGz(dumpFile, tmpDir); err != nil {
		return err
	}

	// Now run your existing analysis on the extracted tree
	// (adjust the call to match your internal packages)
	sites, err := analyzeExtractedDir(tmpDir) // your existing logic, possibly wrapped
	if err != nil {
		return err
	}

	PrintReport(sites)
	return nil
}

// extractTarGz is a simple helper
func extractTarGz(src, dst string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(dst, header.Name)
		if header.Typeflag == tar.TypeDir {
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return err
			}
			f.Close()
		}
	}
	return nil
}

func PrintReport(sites interface{}) { 
  /* call your existing print logic */ }


func analyzeExtractedDir(root string) (interface{}, error) { 
  /* your existing walk + parse */ 
  fmt.Printf("analyzeExtractedDir: %s\n", root )
  return nil, nil }



