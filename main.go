package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	startDir string
	format   string
	print    bool
)

func init() {
	flag.StringVar(&startDir, "d", ".", "starting directory")
	flag.StringVar(&format, "f", "txt", "output format (txt or json)")
	flag.BoolVar(&print, "p", false, "print to console")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

type DirEntry struct {
	Name    string      `json:"name"`
	IsDir   bool        `json:"is_dir"`
	Entries []*DirEntry `json:"entries,omitempty"`
}

func main() {
	if print || (format == "" && startDir == "." && !print) {
		walkDir(startDir, 0, "", os.Stdout)
	} else if format == "json" {
		root := &DirEntry{Name: startDir, IsDir: true}
		walkDirJSON(startDir, root)

		f, _ := os.Create("dir_tree.json")
		defer f.Close()
		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")
		enc.Encode(root)
	} else {
		f, _ := os.Create("dir_tree.txt")
		defer f.Close()
		walkDir(startDir, 0, "", f)
	}
}

func walkDirJSON(path string, parent *DirEntry) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		entry := &DirEntry{Name: file.Name(), IsDir: file.IsDir()}
		parent.Entries = append(parent.Entries, entry)
		if file.IsDir() {
			walkDirJSON(filepath.Join(path, file.Name()), entry)
		}
	}
}

func walkDir(path string, level int, prefix string, w io.Writer) {
	files, _ := ioutil.ReadDir(path)
	for i, file := range files {
		if i == len(files)-1 {
			fmt.Fprintf(w, "%s└── %s\n", prefix, file.Name())
			if file.IsDir() {
				walkDir(filepath.Join(path, file.Name()), level+1, prefix+"    ", w)
			}
		} else {
			fmt.Fprintf(w, "%s├── %s\n", prefix, file.Name())
			if file.IsDir() {
				walkDir(filepath.Join(path, file.Name()), level+1, prefix+"│   ", w)
			}
		}
	}
}
