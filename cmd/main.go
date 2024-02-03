// Package main contains tool entrypoint.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kingpin/v2"

	"github.com/alexeyco/devtools/docker"
	"github.com/alexeyco/devtools/version"
)

var (
	app = kingpin.New("", "")

	tags       = app.Command("tags", "returns tags")
	dockerfile = app.Command("dockerfile", "makes Dockerfile from template")
)

const (
	separator = "\n"

	source = "Dockerfile.template"
	dest   = "Dockerfile"
)

func main() {
	v, err := version.NewProvider(version.NewDefaultGetter()).Version()
	if err != nil {
		log.Fatalln(err)
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case tags.FullCommand():
		t, err := v.Tags()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Print(strings.Join(t, separator))
	case dockerfile.FullCommand():
		if err = docker.NewGenerator(os.DirFS(".")).Dockerfile(v, source, dest); err != nil {
			log.Fatalln(err)
		}
	}
}
