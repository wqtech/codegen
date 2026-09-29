package entcodegen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"entgo.io/ent/entc/gen"
)

const annotationName = "entcodegen"

type Annotation struct {
	TemplateWrite map[string]string
}

func (a *Annotation) Name() string {
	return annotationName
}
func (a *Annotation) Template(name string, out string) *Annotation {
	a.TemplateWrite[name] = out
	return a
}

func (a *Annotation) Generate(t *Template, g *gen.Type) error {
	var buf bytes.Buffer
	for name, out := range a.TemplateWrite {
		buf.Reset()
		_, err := os.Stat(out)
		if err == nil {
			continue
		}

		if err = t.ExecuteTemplate(&buf, name, g); err != nil {
			return fmt.Errorf("execute template: name %s %w", name, err)
		}
		if err = os.WriteFile(out, buf.Bytes(), 0644); err != nil {
			return fmt.Errorf("write file name: name %s  %w", name, err)
		}
	}
	return nil
}
func NewAnnotation() *Annotation {
	return &Annotation{
		TemplateWrite: make(map[string]string),
	}
}

func parserAnnotation(g *gen.Type) (*Annotation, bool) {
	a, ok := g.Annotations[annotationName]
	if !ok {
		return nil, false
	}
	b, err := json.Marshal(a)
	if err != nil {
		return nil, false
	}
	var annotation Annotation
	if err = json.Unmarshal(b, &annotation); err != nil {
		return nil, false
	}
	return &annotation, true
}
