package config

import (
	"fmt"
	"os"
	"path/filepath"
	"portal/internal/constants"

	"gopkg.in/yaml.v3"
)

type GitRepository struct {
	Path string "yaml:path"
	Name string "yaml:name"
}

type Config struct {
	Repos  []GitRepository     `yaml:"repos"`
	Groups map[string][]string `yaml:"groups"`
}

func configDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return constants.PORTAL_CONFIG_DIR
	}
	return filepath.Join(home, constants.PORTAL_CONFIG_DIR)
}

func configPath() string {
	return filepath.Join(configDir(), constants.PORTAL_CONFIG_FILE)
}

func Load() (*Config, error) {
	cfg := &Config{
		Repos:  []GitRepository{},
		Groups: make(map[string][]string),
	}

	data, err := os.ReadFile(configPath())
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.Groups == nil {
		cfg.Groups = make(map[string][]string)
	}

	return cfg, nil
}

func (c *Config) Save() error {
	dir := configDir()
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath(), data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

func (c *Config) AddGitRepositoryToConfig(path, name string) error {
	for _, r := range c.Repos {
		if r.Path == path {
			return fmt.Errorf("Git repository already registered: %s", path)
		}
		if r.Name == name {
			return fmt.Errorf("Git repository name already in use: %s", name)
		}
	}

	c.Repos = append(c.Repos, GitRepository{Path: path, Name: name})
	return nil
}

// TODO: needs optimization O(n^2)
func (c *Config) RemoveGitRepositoryFromConfig(name string) error {
	for i, r := range c.Repos {
		if r.Name == name {
			c.Repos = append(c.Repos[:i], c.Repos[i+1:]...)
			for groupName, repos := range c.Groups {
				for j, repoName := range repos {
					if repoName == name {
						c.Groups[groupName] = append(repos[:j], repos[j+1:]...)
						break
					}
				}
			}
			return nil
		}
	}
	return fmt.Errorf("Git repository not found: %s", name)
}

func (c *Config) GetGitRepository(name string) (*GitRepository, error) {
	for _, r := range c.Repos {
		if r.Name == name {
			return &r, nil
		}
	}
	return nil, fmt.Errorf("Git repository not found: %s", name)
}

func (c *Config) GetGitRepositoryByGroup(groupName string) ([]GitRepository, error) {
	repoNames, ok := c.Groups[groupName]
	if !ok {
		return nil, fmt.Errorf("group not found: %s", groupName)
	}

	repos := make([]GitRepository, 0, len(repoNames))
	for _, name := range repoNames {
		repo, err := c.GetGitRepository(name)
		if err != nil {
			continue
		}
		repos = append(repos, *repo)
	}
	return repos, nil
}

func (c *Config) GetAllRepository() []GitRepository {
	return c.Repos
}

func (c *Config) AddGroup(name string, repoNames []string) error {
	for _, repoName := range repoNames {
		if _, err := c.GetGitRepository(repoName); err != nil {
			return fmt.Errorf("Git repository not found: %s", repoName)
		}
	}
	c.Groups[name] = repoNames
	return nil
}

func (c *Config) RemoveGroup(name string) error {
	if _, ok := c.Groups[name]; !ok {
		return fmt.Errorf("group not found: %s", name)
	}
	delete(c.Groups, name)
	return nil
}
