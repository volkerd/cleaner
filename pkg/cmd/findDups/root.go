package findDups

import (
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"io/fs"
	"log"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

func Exec() {
	fmt.Printf("find duplicates\nsource: %s\n", sourcePath)
	files := traversFS(sourcePath)
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].hash < files[j].hash
	})
	for _, file := range files {
		fmt.Printf("%s %s%s\n", file.hash, file.basename, file.extension)
	}
}

func traversFS(sourcePath string) []File {
	fmt.Printf("TraversFS(%v)\n", sourcePath)
	totalCount := 0
	extensionFilter := map[string]bool{
		".JPG":  true,
		".jpeg": true,
		".jpg":  true,
	}
	extensionFilterKeys := slices.Collect(maps.Keys(extensionFilter))
	fmt.Println("find extensions ", extensionFilterKeys)
	skipFilter := map[string]bool{
		".MOV": true,
		".dop": true,
		".jpg": true,
	}
	fmt.Printf("skip extensions  %v\n", slices.Collect(maps.Keys(skipFilter)))
	fsys := os.DirFS(sourcePath)
	extensions := make(map[string]int)
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && !strings.HasPrefix(d.Name(), ".") {
			totalCount++
			ext := filepath.Ext(path)
			if extensionCount, ok := extensions[ext]; ok {
				extensions[ext] = extensionCount + 1
			} else {
				extensions[ext] = 1
			}
			if _, ok := extensionFilter[ext]; !ok {
				if _, ok := skipFilter[ext]; !ok {
					fmt.Printf("%v\t%v\n", ext, d.Name())
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(extensions)
	files := make([]File, 0)
	err = fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && !strings.HasPrefix(d.Name(), ".") {
			totalCount++
			ext := filepath.Ext(path)
			if _, ok := extensionFilter[ext]; ok {
				if len(files) < 100 {
					absPath := filepath.Join(sourcePath, path)
					reader, err := os.Open(absPath)
					if err != nil {
						panic(err)
					}
					hashValue := createHash(md5.New(), reader)
					file := File{
						path:      absPath,
						basename:  strings.TrimSuffix(filepath.Base(path), ext),
						extension: ext,
						hash:      hashValue}
					fmt.Printf("%v\n", file)
					files = append(files, file)
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func createHash(h hash.Hash, r io.Reader) string {
	_, err := io.Copy(h, r)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	hashValue := fmt.Sprintf("%x", h.Sum(nil))
	return hashValue
}
