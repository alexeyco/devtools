// Package main contains tool entrypoint.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kingpin/v2"

	"github.com/alexeyco/devtools/version"
)

var (
	app = kingpin.New("", "")

	ver  = app.Command("ver", "returns actual go version")
	tags = app.Command("tags", "returns tags")
)

const separator = "\n"

func main() {
	v, err := version.NewProvider(version.NewDefaultGetter()).Version()
	if err != nil {
		log.Fatalln(err)
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case ver.FullCommand():
		fmt.Println(v.String())
	case tags.FullCommand():
		t, err := v.Tags()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(strings.Join(t, separator))
	}
}
