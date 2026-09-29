package codegen

import "codegen/internal/entcodegen"

type EntOption = entcodegen.Option

func EntGenerate(schemaPath string, tmplDir string, options ...EntOption) error {
	if tmplDir != "" {
		options = append(options, entcodegen.WithTemplateDir(tmplDir))
	}
	return entcodegen.Generate(schemaPath, options...)
}

func NewEntAnnotation() *entcodegen.Annotation {
	return entcodegen.NewAnnotation()
}
