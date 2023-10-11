package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type ConfigExec struct {
	Path    string `yaml:"path"`
	Package string `yaml:"package"`
}

type Config struct {
	Base         string     `yaml:"base"`
	Translations []string   `yaml:"translations"`
	Exec         ConfigExec `yaml:"exec"`
}

func (c *Config) ListTranslationFiles() ([]string, []string, error) {
	files := []string{}
	for _, t := range c.Translations {
		secs := strings.Split(t, "*")
		if len(secs) == 1 {
			walkedFiles, err := WalkMatch(secs[0], "*.json")
			if err != nil {
				return nil, nil, err
			}
			files = append(files, walkedFiles...)
		} else if len(secs) == 2 {
			walkedFiles, err := WalkMatch(secs[0], "*"+secs[1])
			if err != nil {
				return nil, nil, err
			}
			files = append(files, walkedFiles...)
		} else {
			return nil, nil, fmt.Errorf("invalid translation file pattern: %s (valid: translations/*.json)", t)
		}
	}

	langs := []string{}

FILES:
	for _, f := range files {
		lang := strings.Split(filepath.Base(f), ".")[0]

		for _, l := range langs {
			if l == lang {
				continue FILES
			}
		}

		langs = append(langs, lang)
	}

	return files, langs, nil
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func readConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
