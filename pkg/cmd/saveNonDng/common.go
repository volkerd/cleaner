package saveNonDng

import (
	"os"
)

var (
	sourcePath string
	targetPath string
)

func init() {
	sourcePath = os.Getenv("SOURCE_PATH")
	targetPath = os.Getenv("TARGET_PATH")
}
