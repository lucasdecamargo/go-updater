package updater

import (
	"archive/zip"
	"fmt"
	"io"
	"path/filepath"
	"slices"
	"strings"
)

var (
	ignorePaths = []string{".DS_Store", "__MACOSX"}
)

func ApplyZip(zipfile string, opts Options) error {
	if opts.TargetPath == "" {
		p, err := ExecutableRealPath()
		if err != nil {
			return fmt.Errorf("executable path: %w", err)
		}

		opts.TargetPath = filepath.Dir(p)
	}

	return streamZipFile(zipfile, func(name string, reader io.Reader) error {
		return applyZipFile(name, reader, opts)
	})
}

func applyZipFile(fname string, freader io.Reader, opts Options) error {
	opts.TargetPath = filepath.Join(opts.TargetPath, fname)
	return Apply(freader, opts)
}

// streamZipFile provides a simpler interface using a visitor pattern
func streamZipFile(zipPath string, visitor func(name string, reader io.Reader) error) error {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("zip reader: %w", err)
	}
	defer reader.Close()

	for _, file := range reader.File {
		if shouldIgnore(file) {
			continue
		}

		// Skip directories
		if file.FileInfo().IsDir() {
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return fmt.Errorf("zip open: %w", err)
		}

		// Visit the file
		err = visitor(file.Name, fileReader)
		fileReader.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

func shouldIgnore(file *zip.File) bool {
	for p := range strings.SplitSeq(file.Name, string(filepath.Separator)) {
		if slices.Contains(ignorePaths, p) {
			return true
		}
	}
	return false
}
