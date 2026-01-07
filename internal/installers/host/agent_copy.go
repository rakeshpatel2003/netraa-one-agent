package host

import (
	"fmt"
	"os"
	"path/filepath"

	"netraa-agent/internal/agents"
	"netraa-agent/internal/utils"
)

// installAgent copies an agent artifact from AGENT_DIR to host base dir
func installAgent(spec agents.AgentSpec, targetName string) error {
	// Always resolve platform dynamically
	spec.Platform = utils.Platform()

	src, err := agents.Resolve(spec)
	if err != nil {
		return err
	}

	dest := filepath.Join(HostBaseDir, targetName)

	fmt.Printf(
		"  • Installing %s (%s)\n",
		spec.Artifact,
		spec.Platform,
	)

	// Remove old version if exists
	_ = os.RemoveAll(dest)

	// Copy securely (fs.go handles security)
	return utils.CopyDir(src, dest)
}

