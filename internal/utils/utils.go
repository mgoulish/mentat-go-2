package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func IsTarGz(path string) bool {
	if path == "" {
		return false
	}
	// Handle both .tar.gz and .tgz
	lower := strings.ToLower(path)
	return strings.HasSuffix(lower, ".tar.gz") || strings.HasSuffix(lower, ".tgz")
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		// Path doesn't exist or is inaccessible
		return false
	}
	return info.IsDir()
}

func ExpandTGZFile(tgzFile string) (string, error) {
	// Create a temp directory for extraction
	tmpDir, err := os.MkdirTemp("", "skupper-mentat-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp dir: %w", err)
	}
	//defer os.RemoveAll(tmpDir)

	fmt.Printf("📦 Extracting tgz %s to %s...\n", tgzFile, tmpDir)

	if err := extractTarGz(tgzFile, tmpDir); err != nil {
		return "", err
	}

	return tmpDir, nil
}

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
