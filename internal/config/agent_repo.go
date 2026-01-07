package config

import "os"

func AgentRepoPath() string {
	if path := os.Getenv("NETRAA_AGENT_REPO"); path != "" {
		return path
	}
	// default for dev
	return "/home/rakesh.patel@apmosys.mahape/Documents/agents-apmosys/AGENT_DIR"
}

