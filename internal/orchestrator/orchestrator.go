package orchestrator

import (
	"fmt"

	"netraa-agent/internal/discovery"
	"netraa-agent/internal/installers/host"
)

// Run executes installers based on user selection
func Run(selected []discovery.ComponentType) error {
	for _, component := range selected {
		switch component {

		case discovery.ComponentHost:
			fmt.Println("\n▶ Starting Host installer")
			if err := host.Install(); err != nil {
				return err
			}

		// Future components (placeholders)
		// case discovery.ComponentJava:
		//     java.Install()
		// case discovery.ComponentMySQL:
		//     mysql.Install()

		default:
			fmt.Println("ℹ No installer implemented for:", component)
		}
	}
	return nil
}

