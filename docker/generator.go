package docker

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"text/template"

	"github.com/alexeyco/devtools/version"
)

type Generator struct {
	fs fs.FS
}

func NewGenerator(fs fs.FS) *Generator {
	return &Generator{
		fs: fs,
	}
}

func (g Generator) Dockerfile(ver version.Version, source, dest string) error {
	b, err := fs.ReadFile(g.fs, source)
	if err != nil {
		return fmt.Errorf("error reading file %q: %w", source, err)
	}

	tpl, err := template.New("dockerfile").Parse(string(b))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	var buf bytes.Buffer
	defer buf.Reset()

	data := map[string]any{
		"Version": ver.String(),
	}

	if err = tpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	if err = os.WriteFile(dest, buf.Bytes(), os.ModePerm); err != nil {
		return fmt.Errorf("error writing file %q: %w", dest, err)
	}

	return nil
}
