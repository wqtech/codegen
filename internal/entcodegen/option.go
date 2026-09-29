package entcodegen

import (
	"path/filepath"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func WithGenConfig(fn func(g *gen.Config)) Option {
	return func(cfg *Config) error {
		fn(&cfg.genConfig)
		return nil
	}
}
func WithTemplateDir(path string) Option {
	return WithTemplateGlob(filepath.Join(path, "/*.tmpl"))
}
func WithTemplateGlob(p string) Option {
	return func(cfg *Config) error {
		ts, err := TemplateGlob(p)
		if err != nil {
			return err
		}
		cfg.rootTemplate = ts
		return nil
	}
}
func WithGenOptions(options ...entc.Option) Option {
	return func(cfg *Config) error {
		cfg.entcOptions = append(cfg.entcOptions, options...)
		return nil
	}
}
