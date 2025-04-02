package saveNonDng

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Exec() {
	fmt.Printf("save non dng file\nsource: %s\ntarget: %s\n", sourcePath, targetPath)
	traversFS(sourcePath)
}

func traversFS(sourcePath string) {
	fmt.Printf("TraversFS(%v)\n", sourcePath)
	totalCount := 0
	findCount := 0
	fsys := os.DirFS(sourcePath)
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && !strings.HasPrefix(d.Name(), ".") {
			totalCount++
			ext := filepath.Ext(path)
			if ext != ".dng" {
				if ext != ".jpg" {
					findCount++
					fmt.Printf("%v of %v (%v): %v\n", findCount, totalCount, path, d.Name())
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
