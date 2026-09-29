package entcodegen

import "entgo.io/ent/entc/gen"

type Template struct {
	*gen.Template
}

func TemplateGlob(p string) (*Template, error) {
	ts, err := gen.NewTemplate("root").ParseGlob(p)
	if err != nil {
		return nil, err
	}
	return &Template{ts}, nil
}
