package entcodegen

import (
	"errors"
	"fmt"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

type Option func(c *Config) error
type Config struct {
	entcOptions  []entc.Option
	genConfig    gen.Config
	rootTemplate *Template
}

func Generate(schemaPath string, options ...Option) error {
	cfg := &Config{
		genConfig: gen.Config{
			Features: []gen.Feature{
				gen.FeatureLock,
				gen.FeatureExecQuery,
				gen.FeatureModifier,
				gen.FeatureEntQL,
			},
		},
	}
	for _, opt := range options {
		if err := opt(cfg); err != nil {
			return err
		}
	}
	cfg.genConfig.Hooks = append(cfg.genConfig.Hooks, func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				annotation, ok := parserAnnotation(node)
				if !ok {
					continue
				}
				if cfg.rootTemplate == nil {
					return errors.New("rootTemplate is nil")
				}
				if err := annotation.Generate(cfg.rootTemplate, node); err != nil {
					return fmt.Errorf("generate template: %w", err)
				}
			}
			return next.Generate(g)
		})
	})

	return entc.Generate(schemaPath, &cfg.genConfig, cfg.entcOptions...)
}
