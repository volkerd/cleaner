package main

import (
	"github.com/dimiro1/banner"
	"github.com/volkerd/cleaner/pkg/cmd/saveNonDng"
	"os"
)

func main() {
	showBanner()
	saveNonDng.Exec()
}

func showBanner() {
	templ := `{{ "saveNonDng" }} {{ .GoVersion }} {{ .GOOS }}/{{ .GOARCH }} {{ .Now "2.1.2006 15:04" }}
`
	banner.InitString(os.Stdout, true, false, templ)

}
