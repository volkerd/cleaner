package findDups

import (
	"os"
)

type File struct {
	basename  string
	extension string
	hash      string
	path      string
}

var (
	sourcePath string
)

func init() {
	sourcePath = os.Getenv("SOURCE_PATH")
}
