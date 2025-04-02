package saveNonDng

import (
	"os"
)

type File struct {
	path     string
	datatype string
}

var (
	sourcePath string
	targetPath string
)

func init() {
	sourcePath = os.Getenv("SOURCE_PATH")
	targetPath = os.Getenv("TARGET_PATH")
}
