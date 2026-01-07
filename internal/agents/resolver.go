package agents

import (
	"fmt"
	"path/filepath"
	"strings"

	"netraa-agent/internal/config"
)

func Resolve(spec AgentSpec) (string, error) {
	base := config.AgentRepoPath()

	// Normalize
	category := strings.ToUpper(spec.Category)
	platform := strings.ToUpper(spec.Platform)
	product := strings.ToUpper(spec.Product)

	// Step 1: base path
	path := filepath.Join(base, category, platform, product)

	// Step 2: version handling
	if spec.Version != "" {
		path = filepath.Join(path, spec.Version)
	} else {
		// auto-pick latest if versioned
		if exists(path) {
			latest, err := latestVersion(path)
			if err == nil {
				path = filepath.Join(path, latest)
			}
		}
	}

	// Step 3: artifact
	artifactPath := filepath.Join(path, spec.Artifact)

	if !exists(artifactPath) {
		return "", fmt.Errorf("agent artifact not found: %s", artifactPath)
	}

	return artifactPath, nil
}

