package prompt

import "netraa-agent/internal/discovery"

type PromptItem struct {
	Component discovery.DiscoveredComponent
	Selected  bool
}

