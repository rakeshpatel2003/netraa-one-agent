package discovery

import (
	"runtime"
)

func DetectHost() DiscoveredComponent {
	return DiscoveredComponent{
		Type: ComponentHost,
		Metadata: map[string]string{
			"os":   runtime.GOOS,
			"arch": runtime.GOARCH,
		},
	}
}

