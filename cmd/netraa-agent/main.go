package main

import (
	"fmt"
	"log"

	"netraa-agent/internal/discovery"
	"netraa-agent/internal/orchestrator"
	"netraa-agent/internal/prompt"
)

func main() {
	// 1. Discover components
	components, err := discovery.Discover()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Ask user what to monitor
	selected := prompt.AskUser(components)

	fmt.Println("\nSelected components to monitor:")
	for _, c := range selected {
		fmt.Println("-", c)
	}

	// 3. Run installers based on selection
	if err := orchestrator.Run(selected); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n✔ Installation completed")
}

