package main

import (
	"github.com/dimiro1/banner"
	"github.com/volkerd/cleaner/pkg/cmd/findDups"
	"os"
)

func main() {
	showBanner()
	findDups.Exec()
}

func showBanner() {
	templ := `{{ "findDups" }} {{ .GoVersion }} {{ .GOOS }}/{{ .GOARCH }} {{ .Now "2.1.2006 15:04" }}
`
	banner.InitString(os.Stdout, true, false, templ)

}
