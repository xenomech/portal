package utils

import (
	"fmt"
	"strings"
)

// ParseFromSpecs parses repo=branch pairs into a map
func ParseFromSpecs(specs []string) (map[string]string, error) {
	result := make(map[string]string)
	for _, spec := range specs {
		parts := strings.SplitN(spec, "=", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return nil, fmt.Errorf("invalid --from format '%s': expected repo=branch", spec)
		}
		repoName := parts[0]
		baseBranch := parts[1]

		if _, exists := result[repoName]; exists {
			return nil, fmt.Errorf("duplicate --from specification for repo '%s'", repoName)
		}
		result[repoName] = baseBranch
	}
	return result, nil
}




